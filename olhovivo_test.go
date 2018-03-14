package olhovivo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	API_TEST_VERSION = "v2.1"
	API_TEST_TOKEN   = "l00k4tm3l00k4tm3th3m0nst3r1ns1d30fm3h4sgr0wnth1sl4rg3"
)

func TestAuthenticate(t *testing.T) {
	ts := createTestServer(t)
	defer ts.Close()

	ov := &OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	ok, err := ov.Authenticate()
	if err != nil {
		t.Errorf("error while trying to authenticate: %s", err.Error())
	}

	if !ok {
		t.Errorf("expected authentication to be '%v', but got '%v'", true, ok)
	}
}

func TestQueryLine(t *testing.T) {
	ts := createTestServer(t)
	defer ts.Close()

	ov := &OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	_, err := ov.QueryLines("8000")
	if err != nil {
		t.Errorf("error while querying lines: %s", err.Error())
	}
}

func createTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	var authenticated bool

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/"+API_TEST_VERSION)

		switch path {
		case "/Login/Autenticar":
			if r.FormValue("token") == "" {
				t.Errorf("expected 'token' form value to be set")
			}

			authenticated = true
			w.Write([]byte("true"))
		case "/Linha/Buscar":
			if !authenticated {
				t.Errorf("http client is not authenticated")
			}

			if r.FormValue("termosBusca") == "" {
				t.Errorf("expected 'termosBusca' form value to be set")
			}

			jsonString := `
				[
					{
						"cl": 1273,
						"lc": false,
						"lt": "8000",
						"sl": 1,
						"tl": 10,
						"tp": "PCA.RAMOS DE AZEVEDO",
						"ts": "TERMINAL LAPA"
					},
					{
						"cl": 34041,
						"lc": false,
						"lt": "8000",
						"sl": 2,
						"tl": 10,
						"tp": "PCA.RAMOS DE AZEVEDO",
						"ts": "TERMINAL LAPA"
					}
				]
			`

			w.Write([]byte(jsonString))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}
