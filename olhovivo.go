package olhovivo

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/net/publicsuffix"
)

const (
	API_URL     = "http://api.olhovivo.sptrans.com.br"
	API_VERSION = "v2.1"

	apiAuthPath = "/Login/Autenticar"
)

var (
	ErrAuthenticationFailed = errors.New("authentication failed")
)

type OlhoVivo struct {
	URL     string
	Version string
	Token   string

	httpClient *http.Client
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
		return false, err
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

func (ov *OlhoVivo) request(method, path string, params url.Values) (resp *http.Response, err error) {
	shouldAuth := (path != apiAuthPath)
	if err = ov.setupHttpClient(shouldAuth); err != nil {
		return
	}

	parsedUrl, err := ov.mountUrl(path, params)
	if err != nil {
		return nil, errors.Wrap(err, "error while mounting request url")
	}

	req := &http.Request{
		Method: method,
		URL:    parsedUrl,
		Header: http.Header{},
	}

	return ov.httpClient.Do(req)
}

func (ov *OlhoVivo) setupHttpClient(shouldAuth bool) (err error) {
	if ov.httpClient != nil {
		return
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return errors.Wrap(err, "error while creating http cookiejar")
	}

	ov.httpClient = &http.Client{Jar: jar}

	// prevents requesting auth route two times when calling Authenticate() directly
	if !shouldAuth {
		return
	}

	ok, err := ov.Authenticate()
	if err != nil {
		return errors.Wrap(err, "error while trying to authenticate on OlhoVivo")
	}
	if !ok {
		return ErrAuthenticationFailed
	}

	return
}

func (ov *OlhoVivo) mountUrl(path string, params url.Values) (parsedUrl *url.URL, err error) {
	parsedUrl, err = url.Parse(ov.URL)
	if err != nil {
		return nil, errors.Wrap(err, "error while parsing url")
	}

	parsedUrl.Path = "/" + ov.Version + path
	parsedUrl.RawQuery = params.Encode()

	return
}
