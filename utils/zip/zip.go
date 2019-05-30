package zip

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/qinhan-shu/gp-server/utils/file"
)

// Info : details of test data
type Info struct {
	TestCaseNumber int                 `json:"test_case_number"`
	Spj            bool                `json:"spj"`
	TestCases      map[string]TestCase `json:"test_cases"`
}

// TestCase : info of test case
type TestCase struct {
	OutputMd5         string `json:"output_md5"`
	StrippedOutputMd5 string `json:"stripped_output_md5"`
	OutputSize        int64  `json:"output_size"`
	InputName         string `json:"input_name"`
	InputSize         int64  `json:"input_size"`
	OutputName        string `json:"output_name"`
}

func isZip(zipPath string) bool {
	f, err := os.Open(zipPath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// Unzip : decode zip
func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			continue
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			continue
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			continue
		}
	}

	testCases, err := getAllTestCases(target)
	if err != nil {
		return err
	}
	info := Info{
		Spj:            false,
		TestCaseNumber: len(testCases),
		TestCases:      testCases,
	}

	// create info
	bytes, err := json.Marshal(info)
	if err != nil {
		return err
	}
	path := filepath.Join(target, "info")
	file.Write(path, bytes)

	return nil
}

// getAllTestCases : get all test cases
func getAllTestCases(dirPth string) (map[string]TestCase, error) {
	testCases := make(map[string]TestCase)
	pthSep := string(os.PathSeparator)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if !fi.IsDir() {
			ok := strings.HasSuffix(fi.Name(), ".in")
			if ok {
				// get the path of .in and .out file
				in := fi.Name()
				results := strings.Split(in, ".")
				if len(results) != 2 {
					return nil, fmt.Errorf("invaild in file name : %s", in)
				}
				out := results[0] + ".out"
				files := []string{
					dirPth + pthSep + in,
					dirPth + pthSep + out,
				}

				// add new test case
				testCase := TestCase{
					InputName:  in,
					OutputName: out,
				}

				// get file size
				for _, file := range files {
					fileInfo, err := os.Stat(file)
					if err != nil {
						return nil, err
					}
					if file == files[0] {
						testCase.InputSize = fileInfo.Size()
					} else {
						testCase.OutputSize = fileInfo.Size()
					}
				}

				bytes, err := ioutil.ReadFile(files[1])
				if err != nil {
					return nil, err
				}
				testCase.OutputMd5 = fmt.Sprintf("%x", md5.Sum(bytes))
				testCase.StrippedOutputMd5 = fmt.Sprintf("%x", md5.Sum([]byte(strings.TrimSpace(string(bytes)))))
				testCases[results[0]] = testCase
			}
		}
	}

	return testCases, nil
}
