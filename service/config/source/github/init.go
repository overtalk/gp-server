package source

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

// Github describes github repo of gmdata
type Github struct {
	user    string
	token   string
	repoURL string
}

// NewGithub creates a new Github
func NewGithub() *Github {
	require := []string{
		"GITHUB_URL",
		"GITHUB_USERNAME",
		"GITHUB_TOKEN",
	}

	conf := make(map[string]string)
	for _, key := range require {
		value, isExist := os.LookupEnv(key)
		if !isExist {
			logger.Sugar.Fatalf(`Environment "%s" must be set`, key)
		}
		conf[key] = value
	}

	return &Github{
		user:    conf["GITHUB_USERNAME"],
		token:   conf["GITHUB_TOKEN"],
		repoURL: conf["GITHUB_URL"],
	}
}

// Fetch is to get details from github
func (g *Github) fetch(fileName string) ([]byte, error) {
	fileURL := fmt.Sprintf("%s/%s", g.repoURL, fileName)
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, err := http.NewRequest(http.MethodGet, fileURL, nil)
	if err != nil {
		logger.Sugar.Errorf("failed to fetch %s: %v", fileName, err)
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.VERSION.raw")
	req.SetBasicAuth(g.user, g.token)

	resp, err := client.Do(req)
	if err != nil {
		logger.Sugar.Errorf("failed to fetch file[%s]: %v", fileName, err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Sugar.Errorf("failed to read %s: %v", fileName, err)
		return nil, err
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		logger.Sugar.Errorf("file %s is empty", fileName)
		return nil, module.ErrInvalidConfigJSON
	}

	if !json.Valid(body) {
		return nil, module.ErrInvalidConfigJSON
	}
	return body, nil
}
