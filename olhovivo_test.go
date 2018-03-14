package olhovivo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("true"))
	}))

	defer ts.Close()

	ov := &OlhoVivo{
		URL:     ts.URL,
		Version: "v2.1",
		Token:   "l00k4tm3l00k4tm3th3m0nst3r1ns1d30fm3h4sgr0wnth1sl4rg3",
	}

	ok, err := ov.Authenticate()
	if err != nil {
		t.Errorf("error while trying to authenticate: %s", err.Error())
	}

	if !ok {
		t.Errorf("expected authentication to be '%v', but got '%v'", true, ok)
	}
}
