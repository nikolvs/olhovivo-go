package olhovivo

type BusPositionInfo struct {
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

func (ov *OlhoVivo) Positions() (positions BusPositionInfo, err error) {
	err = ov.request(&positions, "GET", "/Posicao", nil)
	return
}
