package steam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

var client *http.Client
var baseURL string

func init() {
	client = &http.Client{
		Timeout: time.Second * 20,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	if os.Getenv("STEAM_API_URL") != "" {
		baseURL = os.Getenv("STEAM_API_URL")
	}
}

func exec(method string, path string, params any) (status int, body []byte, err error) {
	url := fmt.Sprintf("%s%s", baseURL, path)

	var buf io.ReadWriter
	if params != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(params)
		if err != nil {
			return
		}
	}

	req, err := http.NewRequest(method, url, buf)

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return
	}

	defer res.Body.Close()

	status = res.StatusCode

	body, _ = ioutil.ReadAll(res.Body)
	return
}
