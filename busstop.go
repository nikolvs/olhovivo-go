package olhovivo

import (
	"net/url"
	"strconv"
)

type BusStop struct {
	Cp int     `json:"cp"`
	Np string  `json:"np"`
	Ed string  `json:"ed"`
	Py float64 `json:"py"`
	Px float64 `json:"px"`
}

func (ov *OlhoVivo) QueryStops(search string) (stops []BusStop, err error) {
	err = ov.request(&stops, "GET", "/Parada/Buscar", url.Values{
		"termosBusca": []string{search},
	})

	return
}

func (ov *OlhoVivo) QueryStopsByLine(lineCode int) (stops []BusStop, err error) {
	err = ov.request(&stops, "GET", "/Parada/BuscarParadasPorLinha", url.Values{
		"codigoLinha": []string{strconv.Itoa(lineCode)},
	})

	return
}

func (ov *OlhoVivo) QueryStopsByCorridor(corridorCode int) (stops []BusStop, err error) {
	err = ov.request(&stops, "GET", "/Parada/BuscarParadasPorCorredor", url.Values{
		"codigoCorredor": []string{strconv.Itoa(corridorCode)},
	})

	return
}
