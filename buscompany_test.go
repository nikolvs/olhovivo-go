package olhovivo_test

import (
	"testing"

	olhovivo "github.com/nikolvs/olhovivo-go"
	"github.com/nikolvs/olhovivo-go/olhovivotest"
)

func TestCompanies(t *testing.T) {
	ts := olhovivotest.NewServer(API_TEST_VERSION)
	defer ts.Close()

	ov := &olhovivo.OlhoVivo{
		URL:     ts.URL,
		Version: API_TEST_VERSION,
		Token:   API_TEST_TOKEN,
	}

	companies, err := ov.Companies()
	if err != nil {
		t.Errorf("error while fetching companies: %s", err.Error())
	}

	if len(companies) == 0 {
		t.Errorf("expected to found some bus companies, but got nothing")
	}
}
