package olhovivo

import (
	"net/url"
	"strconv"
)

type BusPrevisions struct {
	Hr string           `json:"hr"`
	P  BusStopPrevision `json:"p"`
}

type BusStopPrevisions struct {
	Hr string           `json:"hr"`
	P  BusStopPrevision `json:"p"`
}

type BusStopPrevision struct {
	Cp int                   `json:"cp"`
	Np string                `json:"np"`
	Py float64               `json:"py"`
	Px float64               `json:"px"`
	L  []BusLineLocalization `json:"l"`
}

type BusLinePrevision struct {
	Cp int                        `json:"cp"`
	Np string                     `json:"np"`
	Py float64                    `json:"py"`
	Px float64                    `json:"px"`
	Vs []BusPrevisionLocalization `json:"vs"`
}

type BusPrevisionLocalization struct {
	P  int     `json:"p"`
	T  string  `json:"t"`
	A  bool    `json:"a"`
	Ta string  `json:"ta"`
	Py float64 `json:"py"`
	Px float64 `json:"px"`
}

type BusLinePrevisions struct {
	Hr string             `json:"hr"`
	Ps []BusLinePrevision `json:"ps"`
}

func (ov *OlhoVivo) Previsions(stopCode, lineCode int) (previsions BusPrevisions, err error) {
	err = ov.request(&previsions, "GET", "/Previsao", url.Values{
		"codigoParada": []string{strconv.Itoa(stopCode)},
		"codigoLinha":  []string{strconv.Itoa(lineCode)},
	})

	return
}

func (ov *OlhoVivo) LinePrevisions(lineCode int) (previsions BusLinePrevisions, err error) {
	err = ov.request(&previsions, "GET", "/Previsao/Linha", url.Values{
		"codigoLinha": []string{strconv.Itoa(lineCode)},
	})

	return
}

func (ov *OlhoVivo) StopPrevisions(stopCode int) (previsions BusStopPrevisions, err error) {
	err = ov.request(&previsions, "GET", "/Previsao/Parada", url.Values{
		"codigoParada": []string{strconv.Itoa(stopCode)},
	})

	return
}
