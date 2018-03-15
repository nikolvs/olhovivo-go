package olhovivo_test

import (
	"testing"

	olhovivo "github.com/nikolvs/olhovivo-go"
	"github.com/nikolvs/olhovivo-go/olhovivotest"
)

func TestQueryLine(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	_, err := ov.QueryLines("8000")
	if err != nil {
		t.Errorf("error while querying lines: %s", err.Error())
	}
}
