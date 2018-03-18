package olhovivo_test

import (
	"testing"

	olhovivo "github.com/nikolvs/olhovivo-go"
	"github.com/nikolvs/olhovivo-go/olhovivotest"
)

func TestCorridors(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	corridors, err := ov.Corridors()
	if err != nil {
		t.Errorf("error while fetching corridors: %s", err.Error())
	}

	if len(corridors) == 0 {
		t.Errorf("expected to found some bus corridors, but got nothing")
	}
}
