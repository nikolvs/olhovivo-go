package olhovivo

import "net/url"

type BusStop struct {
	Cp int     `json:"cp"`
	Np string  `json:"np"`
	Ed string  `json:"ed"`
	Py float64 `json:"py"`
	Px float64 `json:"px"`
}

func (ov *OlhoVivo) QueryStops(search string) (stops []BusStop, err error) {
	err = ov.requestJSON(&stops, "GET", "/Parada/Buscar", url.Values{
		"termosBusca": []string{search},
	})

	return
}
