package olhovivo_test

import (
	"testing"

	olhovivo "github.com/nikolvs/olhovivo-go"
	"github.com/nikolvs/olhovivo-go/olhovivotest"
)

func TestQueryStops(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	stops, err := ov.QueryStops("8000")
	if err != nil {
		t.Errorf("error while querying stops: %s", err.Error())
	}

	if len(stops) == 0 {
		t.Errorf("expected to found some bus stops, but got nothing")
	}
}
