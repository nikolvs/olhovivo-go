package olhovivo

type BusCompanyInfo struct {
	Hr string           `json:"hr"`
	E  []BusCompanyArea `json:"e"`
}

type BusCompanyArea struct {
	A int          `json:"a"`
	E []BusCompany `json:"e"`
}

type BusCompany struct {
	A int    `json:"a"`
	C int    `json:"c"`
	N string `json:"n"`
}

func (ov *OlhoVivo) Companies() (companies []BusCompanyInfo, err error) {
	err = ov.request(&companies, "GET", "/Empresa", nil)
	return
}
