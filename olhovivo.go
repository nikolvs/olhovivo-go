package olhovivo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

const (
	API_URL     = "http://api.olhovivo.sptrans.com.br"
	API_VERSION = "v2.1"
)

type OlhoVivo struct {
	URL     string
	Version string
	Token   string

	client *http.Client
}

type BusLine struct {
	Cl int    `json:"cl"`
	Lc bool   `json:"lc"`
	Lt string `json:"lt"`
	Tl int    `json:"tl"`
	Sl int    `json:"sl"`
	Tp string `json:"tp"`
	Ts string `json:"ts"`
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
		return false, errors.Wrap(err, "request error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, errors.Wrap(err, "error while reading response body")
	}

	defer resp.Body.Close()

	ok, err = strconv.ParseBool(string(body))
	if err != nil {
		return false, errors.Wrap(err, "error while parsing response body")
	}

	return
}

func (ov *OlhoVivo) QueryLines(search string) (lines []BusLine, err error) {
	// TODO: move authentication to another place
	ok, err := ov.Authenticate()
	if err != nil {
		return nil, errors.Wrap(err, "error on OlhoVivo authentication")
	}
	if !ok {
		return nil, errors.New("failed to authenticate")
	}

	resp, err := ov.request("GET", "/Linha/Buscar", url.Values{
		"termosBusca": []string{search},
	})

	if err != nil {
		return nil, errors.Wrap(err, "request error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error while reading response body")
	}

	defer resp.Body.Close()
	if err := json.Unmarshal(body, &lines); err != nil {
		return nil, errors.Wrap(err, "error while unmarshaling response body")
	}

	return
}

func (ov *OlhoVivo) request(method, path string, params url.Values) (resp *http.Response, err error) {
	if ov.client == nil {
		ov.client = &http.Client{}
	}

	parsedUrl, err := ov.mountUrl(path, params)
	if err != nil {
		return nil, errors.Wrap(err, "error while mounting request url")
	}

	req := &http.Request{
		Method: method,
		URL:    parsedUrl,
	}

	return ov.client.Do(req)
}

func (ov *OlhoVivo) mountUrl(path string, params url.Values) (parsedUrl *url.URL, err error) {
	parsedUrl, err = url.Parse(ov.URL)
	if err != nil {
		return nil, errors.Wrap(err, "error while parsing OlhoVivo url")
	}

	parsedUrl.Path = "/" + ov.Version + path
	parsedUrl.RawQuery = params.Encode()

	return
}
