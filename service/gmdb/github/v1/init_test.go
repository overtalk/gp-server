package gmdb_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFetch(t *testing.T) {

}

func test(t *testing.T) {
	repoURL := "https://api.github.com/repos/QHasaki/config/contents"
	fileName := "server.json"
	username := ""
	token := "81bef83a8b57d3a1d49272b7e3e05b4bd32fb890"

	fileURL := fmt.Sprintf("%s/%s", repoURL, fileName)
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fileURL, nil)
	if err != nil {
		t.Errorf("failed to fetch %s: %v", fileName, err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github.VERSION.raw")
	req.SetBasicAuth(username, token)

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to fetch %s: %v", fileName, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read %s: %v", fileName, err)
		return
	}
	defer resp.Body.Close()

	encodeString := base64.StdEncoding.EncodeToString(body)

	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(decodeBytes))
}
