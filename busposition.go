package olhovivo

import (
	"net/url"
	"strconv"
)

type BusPositions struct {
	Hr string                `json:"hr"`
	L  []BusLineLocalization `json:"l"`
}

type BusLineLocalization struct {
	C   string            `json:"c"`
	Cl  int               `json:"cl"`
	Sl  int               `json:"sl"`
	Lt0 string            `json:"lt0"`
	Lt1 string            `json:"lt1"`
	Qv  int               `json:"qv"`
	Vs  []BusLocalization `json:"vs"`
}

type BusLocalization struct {
	P  int     `json:"p"`
	A  bool    `json:"a"`
	Ta string  `json:"ta"`
	Py float64 `json:"py"`
	Px float64 `json:"px"`
}

type BusLinePositions struct {
	Hr string            `json:"hr"`
	Vs []BusLocalization `json:"vs"`
}

func (ov *OlhoVivo) Positions() (positions BusPositions, err error) {
	err = ov.request(&positions, "GET", "/Posicao", nil)
	return
}

func (ov *OlhoVivo) LinePositions(lineCode int) (positions BusLinePositions, err error) {
	err = ov.request(&positions, "GET", "/Posicao/Linha", url.Values{
		"codigoLinha": []string{strconv.Itoa(lineCode)},
	})

	return
}

func (ov *OlhoVivo) GaragePositions(companyCode int, lineCode ...int) (positions BusPositions, err error) {
	params := url.Values{
		"codigoEmpresa": []string{strconv.Itoa(companyCode)},
	}

	if len(lineCode) != 0 {
		params["codigoLinha"] = []string{strconv.Itoa(lineCode[0])}
	}

	err = ov.request(&positions, "GET", "/Posicao/Garagem", params)
	return
}
