package olhovivo_test

import (
	"testing"

	olhovivo "github.com/nikolvs/olhovivo-go"
	"github.com/nikolvs/olhovivo-go/olhovivotest"
)

func TestPositions(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	positions, err := ov.Positions()
	if err != nil {
		t.Errorf("error while fetching position: %s", err.Error())
	}

	if len(positions.L) == 0 {
		t.Errorf("expected to found some bus lines, but got nothing")
	}
}
