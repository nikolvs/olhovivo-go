package olhovivo

import (
	"net/url"
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
	err = ov.requestJSON(&lines, "GET", "/Linha/Buscar", url.Values{
		"termosBusca": []string{search},
	})

	return
}

func (ov *OlhoVivo) QueryLinesByDirection(search string, direction byte) (lines []BusLine, err error) {
	err = ov.requestJSON(&lines, "GET", "/Linha/BuscarLinhaSentido", url.Values{
		"termosBusca": []string{search},
		"sentido":     []string{string(direction)},
	})

	return
}
