package olhovivo_test

import (
	"testing"

	olhovivo "github.com/nikolvs/olhovivo-go"
	"github.com/nikolvs/olhovivo-go/olhovivotest"
)

func TestPrevisions(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	previsions, err := ov.Previsions(4200953, 4634)
	if err != nil {
		t.Errorf("error while fetching previsions: %s", err.Error())
	}

	if len(previsions.P.L) == 0 {
		t.Errorf("expected to found some bus lines, but got nothing")
	}
}

func TestLinePrevisions(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	previsions, err := ov.LinePrevisions(4634)
	if err != nil {
		t.Errorf("error while fetching previsions: %s", err.Error())
	}

	if len(previsions.Ps) == 0 {
		t.Errorf("expected to found some bus stops, but got nothing")
	}
}

func TestStopPrevisions(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	previsions, err := ov.StopPrevisions(666)
	if err != nil {
		t.Errorf("error while fetching previsions: %s", err.Error())
	}

	if len(previsions.P.L) == 0 {
		t.Errorf("expected to found some bus lines, but got nothing")
	}
}
