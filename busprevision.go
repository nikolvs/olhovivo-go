package olhovivo

import (
	"net/url"
	"strconv"
)

type BusPrevisions struct {
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

type BusLinePrevisions struct {
	Hr string             `json:"hr"`
	Ps []BusStopPrevision `json:"ps"`
}

type BusStopPrevisions struct {
	Hr string           `json:"hr"`
	P  BusStopPrevision `json:"p"`
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
