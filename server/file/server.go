package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/utils/parse"
	"github.com/qinhan-shu/gp-server/utils/zip"
)

var (
	tag    string
	commit string
	branch string

	version    = flag.Bool("version", false, "show version") // show version
	uploadPath = flag.String("uploadPath", "./tmp", "upload path")
)

var (
	maxSize int64 = 0
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println(formatFullVersion())
		return
	}

	// get config
	c := config.NewConfig()
	dataStorage, err := c.GetDataStorageConfigs()
	if err != nil {
		logger.Sugar.Fatalf("failed to get data storage : %v", err)
	}

	// init logger
	logLevel, isExist := dataStorage.Configs.Load("LOG_LEVEL")
	if !isExist {
		logLevel = "info"
	}
	logger.InitLogger(parse.String(logLevel))

	// get gate env
	port, isExist := dataStorage.Configs.Load("FILE_PORT")
	if !isExist {
		port = ":8081"
	}
	maxUploadSize, isExist := dataStorage.Configs.Load("MAXUPLOADSIZE")
	if !isExist {
		port = 2
	}
	maxSize = parse.Int(maxUploadSize) * 1024 * 1024
	http.HandleFunc("/upload", uploadFileHandler())
	http.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir(*uploadPath))))

	logger.Sugar.Infof("starting file server server on %s", port)
	logger.Sugar.Infof("uploadPath : %s", *uploadPath)
	logger.Sugar.Infof("maxUploadSize : %d", maxSize/(1024*1024))
	logger.Sugar.Info("use /upload for uploading files and /files/{fileName} for downloading")
	logger.Sugar.Fatal(http.ListenAndServe(parse.String(port), nil))
}

func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "x-requested-with") //header的类型
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxSize)
		if err := r.ParseMultipartForm(maxSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		targetFile, header, err := r.FormFile("uploadFile")
		if err != nil {
			logger.Sugar.Errorf("%v", err)
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer targetFile.Close()
		fileBytes, err := ioutil.ReadAll(targetFile)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		switch filetype {
		case "application/zip":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}

		// create a dir to store zip
		results := strings.Split(header.Filename, ".")
		if len(results) != 2 || results[1] != "zip" {
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}

		dir := *uploadPath + results[0]
		newPath := filepath.Join(dir, header.Filename)

		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			fmt.Println(err)
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}

		// unzip
		if err := zip.Unzip(dir+"/"+header.Filename, dir); err != nil {
			logger.Sugar.Error(err)
		}
		w.Write([]byte("ok, path = " + dir))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func formatFullVersion() string {
	var parts = []string{"gp_server"}

	if tag != "" {
		parts = append(parts, tag)
	} else {
		parts = append(parts, "(tag: unknown)")
	}

	if branch != "" || commit != "" {
		if branch == "" {
			branch = "unknown_branch"
		}
		if commit == "" {
			commit = "unknown_commit"
		}
	}
	git := fmt.Sprintf("(git: %s %s)", branch, commit)
	parts = append(parts, git)

	return strings.Join(parts, "  ")
}
