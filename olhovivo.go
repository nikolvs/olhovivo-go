package olhovivo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	API_URL     = "http://api.olhovivo.sptrans.com.br"
	API_VERSION = "v2.1"
)

type OlhoVivo struct {
	URL     string
	Version string
	Token   string
	client  *http.Client
}

func New(token string) *OlhoVivo {
	return &OlhoVivo{
		URL:     API_URL,
		Version: API_VERSION,
		Token:   token,
	}
}

func (ov *OlhoVivo) Authenticate() (ok bool, err error) {
	resp, err := ov.request("POST", "/Login/Autenticar", url.Values{
		"token": []string{ov.Token},
	})

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	return strconv.ParseBool(string(body))
}

func (ov *OlhoVivo) request(method, path string, params url.Values) (resp *http.Response, err error) {
	if ov.client == nil {
		ov.client = &http.Client{}
	}

	parsedUrl, err := ov.mountUrl(path)
	if err != nil {
		return
	}

	req := &http.Request{
		Method:   method,
		URL:      parsedUrl,
		PostForm: params,
	}

	return ov.client.Do(req)
}

func (ov *OlhoVivo) mountUrl(path string) (parsedUrl *url.URL, err error) {
	urlStr := fmt.Sprintf("%s/%s/%s", ov.URL, ov.Version, path)
	return url.Parse(urlStr)
}
