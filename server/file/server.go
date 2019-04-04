package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/qinhan-shu/gp-server/logger"
)

var (
	tag    string
	commit string
	branch string

	version       = flag.Bool("version", false, "show version") // show version
	maxUploadSize = flag.Int64("maxUploadSize", 2, "max upload size")
	port          = flag.String("addr", ":8888", "listen address")
	uploadPath    = flag.String("uploadPath", "./tmp", "upload path")
	certFile      = flag.String("certFile", "", "ssl certficate filename")
	keyFile       = flag.String("keyFile", "", "ssl private key filename")
	logLevel      = flag.String("log-level", "error", "log level, optional( debug | info | warn | error | dpanic | panic | fatal), default is error")
)

func main() {
	flag.Parse()
	*maxUploadSize = *maxUploadSize * 1024 * 1024

	if *version {
		fmt.Println(formatFullVersion())
		return
	}

	// init logger
	logger.InitLogger(*logLevel)

	http.HandleFunc("/upload", uploadFileHandler())
	http.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir(*uploadPath))))

	logger.Sugar.Infof("starting file server server on %s", *port)
	logger.Sugar.Infof("uploadPath : %s", *uploadPath)
	logger.Sugar.Infof("maxUploadSize : %d", *maxUploadSize)
	logger.Sugar.Info("use /upload for uploading files and /files/{fileName} for downloading")
	logger.Sugar.Fatal(http.ListenAndServe(*port, nil))
}

func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, *maxUploadSize)
		if err := r.ParseMultipartForm(*maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		file, header, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
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
		fileName := randToken(12)

		newPath := filepath.Join(*uploadPath, fileName+header.Filename)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("SUCCESS"))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
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
