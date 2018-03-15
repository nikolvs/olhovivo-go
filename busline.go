package olhovivo

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/pkg/errors"
)

type BusLine struct {
	Cl int    `json:"cl"`
	Lc bool   `json:"lc"`
	Lt string `json:"lt"`
	Tl int    `json:"tl"`
	Sl int    `json:"sl"`
	Tp string `json:"tp"`
	Ts string `json:"ts"`
}

func (ov *OlhoVivo) QueryLines(search string) (lines []BusLine, err error) {
	resp, err := ov.request("GET", "/Linha/Buscar", url.Values{
		"termosBusca": []string{search},
	})

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error while reading response body")
	}

	defer resp.Body.Close()
	if err := json.Unmarshal(body, &lines); err != nil {
		return nil, errors.Wrap(err, "error while unmarshaling response body")
	}

	return
}
