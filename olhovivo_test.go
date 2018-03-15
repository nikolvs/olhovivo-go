package olhovivo_test

import (
	"testing"

	olhovivo "github.com/nikolvs/olhovivo-go"
	"github.com/nikolvs/olhovivo-go/olhovivotest"
)

const (
	API_TEST_VERSION = "v2.1"
	API_TEST_TOKEN   = "l00k4tm3l00k4tm3th3m0nst3r1ns1d30fm3h4sgr0wnth1sl4rg3"
)

func TestAuthenticate(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
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
