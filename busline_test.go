package olhovivo_test

import (
	"testing"

	olhovivo "github.com/nikolvs/olhovivo-go"
	"github.com/nikolvs/olhovivo-go/olhovivotest"
)

func TestQueryLines(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	lines, err := ov.QueryLines("8000")
	if err != nil {
		t.Errorf("error while querying lines: %s", err.Error())
	}

	if len(lines) == 0 {
		t.Errorf("expected to found some bus lines, but got nothing")
	}
}

func TestQueryLinesByDirection(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	lines, err := ov.QueryLinesByDirection("8000", 1)
	if err != nil {
		t.Errorf("error while querying lines: %s", err.Error())
	}

	if len(lines) == 0 {
		t.Errorf("expected to found some bus lines, but got nothing")
	}
}
