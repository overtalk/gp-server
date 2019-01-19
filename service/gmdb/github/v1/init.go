package gmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"graduationProject/server/logger"
)

// Github describes github repo of gmdata
type Github struct {
	username string
	token    string
	repoURL  string
}

// NewGithub creates a new Github
func NewGithub(repoURL string) *Github {
	if os.Getenv("GITHUB_USERNAME") == "" {
		log.Fatal(`Enviroment "GITHUB_USERNAME" must be set`)
	}
	if os.Getenv("GITHUB_TOKEN") == "" {
		log.Fatal(`Environment "GITHUB_TOKEN" must be set`)
	}
	return &Github{
		username: os.Getenv("GITHUB_USERNAME"),
		token:    os.Getenv("GITHUB_TOKEN"),
		repoURL:  repoURL,
	}
}

func (g *Github) fetch(fileName string) ([]byte, error) {
	fileURL := fmt.Sprintf("%s/%s", g.repoURL, fileName)
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fileURL, nil)
	if err != nil {
		logger.Sugar.Errorf("failed to fetch %s: %v", fileName, err)
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.VERSION.raw")
	req.SetBasicAuth(g.username, g.token)

	resp, err := client.Do(req)
	if err != nil {
		logger.Sugar.Errorf("failed to fetch %s: %v", fileName, err)
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
		return nil, ErrInvalidConfigJSON
	}

	if !json.Valid(body) {
		return nil, ErrInvalidConfigJSON
	}
	return body, nil
}
