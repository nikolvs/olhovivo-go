package olhovivo

type BusCorridor struct {
	Cc int    `json:"cc"`
	Nc string `json:"nc"`
}

func (ov *OlhoVivo) Corridors() (corridors []BusCorridor, err error) {
	err = ov.requestJSON(&corridors, "GET", "/Corredor", nil)
	return
}
