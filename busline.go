package olhovivo

import (
	"net/url"
	"strconv"
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
	err = ov.request(&lines, "GET", "/Linha/Buscar", url.Values{
		"termosBusca": []string{search},
	})

	return
}

func (ov *OlhoVivo) QueryLinesByDirection(search string, direction int) (lines []BusLine, err error) {
	err = ov.request(&lines, "GET", "/Linha/BuscarLinhaSentido", url.Values{
		"termosBusca": []string{search},
		"sentido":     []string{strconv.Itoa(direction)},
	})

	return
}
