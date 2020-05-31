package naughty

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var (
	naughtListUrl = "https://raw.githubusercontent.com/minimaxir/big-list-of-naughty-strings/master/blns.base64.json"
	naughtyList   []string
)

func FetchNaugtyList(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &naughtyList); err != nil {
		return nil, err
	}

	return naughtyList, nil
}

func IsNaugthy(word string) bool {
	if len(naughtyList) == 0 {
		naughtyList, _ = FetchNaugtyList(naughtListUrl)
	}
	for _, item := range naughtyList {
		if item == word {
			return true
		}
	}
	return false
}
