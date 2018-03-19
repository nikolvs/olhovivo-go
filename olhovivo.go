package olhovivo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

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
	err = ov.request(&ok, "POST", "/Login/Autenticar", url.Values{
		"token": []string{ov.Token},
	})

	return
}

func (ov *OlhoVivo) request(v interface{}, method, path string, params url.Values) (err error) {
	shouldAuth := (path != apiAuthPath)
	if err = ov.setupHttpClient(shouldAuth); err != nil {
		return
	}

	parsedUrl, err := ov.mountUrl(path, params)
	if err != nil {
		return errors.Wrap(err, "error while mounting request url")
	}

	req, err := http.NewRequest(method, parsedUrl.String(), nil)
	if err != nil {
		return errors.Wrap(err, "error while mounting http request")
	}

	resp, err := ov.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error while sending http request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "error while reading response body")
	}

	defer resp.Body.Close()
	if err := json.Unmarshal(body, &v); err != nil {
		return errors.Wrap(err, "error while unmarshaling response body")
	}

	return
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
