package olhovivotest

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

func NewServer(version string) *httptest.Server {
	return httptest.NewServer(ServerHandler(version))
}

func ServerHandler(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		path := strings.TrimPrefix(req.URL.EscapedPath(), "/"+version)
		if req.Method == "POST" && path == "/Login/Autenticar" {
			handleAuth(w, req)
			return
		}

		if !isAuthenticated(req) {
			w.Write([]byte(`{"Message": "Authorization has been denied for this request."}`))
			return
		}

		if req.Method != "GET" {
			return
		}

		switch path {
		case "/Linha/Buscar":
			handleQueryLines(w, req)
		case "/Linha/BuscarLinhaSentido":
			handleQueryLinesByDirecction(w, req)
		case "/Parada/Buscar":
			handleQueryStops(w, req)
		case "/Parada/BuscarParadasPorLinha":
			handleQueryStopsByLine(w, req)
		case "/Parada/BuscarParadasPorCorredor":
			handleQueryStopsByCorridor(w, req)
		case "/Corredor":
			handleCorridors(w, req)
		case "/Empresa":
			handleCompanies(w, req)
		case "/Posicao":
			handlePositions(w, req)
		case "/Posicao/Linha":
			handleLinePositions(w, req)
		case "/Posicao/Garagem":
			handleGaragePositions(w, req)
		case "/Previsao":
			handlePrevisions(w, req)
		case "/Previsao/Linha":
			handleLinePrevisions(w, req)
		case "/Previsao/Parada":
			handleStopPrevisions(w, req)
		}
	}
}

func isAuthenticated(req *http.Request) bool {
	if _, err := req.Cookie("apiCredentials"); err != nil {
		return false
	}

	return true
}

func handleAuth(w http.ResponseWriter, req *http.Request) {
	if req.FormValue("token") == "" {
		w.Write([]byte("false"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "apiCredentials",
		Value: "wh4t4w0nd3rfuln4m31tw4sj0h4n",
		Path:  "/",
	})

	w.Write([]byte("true"))
}

func handleQueryLines(w http.ResponseWriter, req *http.Request) {
	if req.FormValue("termosBusca") == "" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		[
			{
				"cl": 1273,
				"lc": false,
				"lt": "8000",
				"sl": 1,
				"tl": 10,
				"tp": "PCA.RAMOS DE AZEVEDO",
				"ts": "TERMINAL LAPA"
			},
			{
				"cl": 34041,
				"lc": false,
				"lt": "8000",
				"sl": 2,
				"tl": 10,
				"tp": "PCA.RAMOS DE AZEVEDO",
				"ts": "TERMINAL LAPA"
			}
		]
	`

	w.Write([]byte(jsonString))
}

func handleQueryLinesByDirecction(w http.ResponseWriter, req *http.Request) {
	if req.FormValue("termosBusca") == "" {
		w.Write([]byte(`[]`))
		return
	}

	direction := req.FormValue("sentido")
	if direction != "1" && direction != "2" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		[
			{
				"cl": 1273,
				"lc": false,
				"lt": "8000",
				"sl": 1,
				"tl": 10,
				"tp": "PCA.RAMOS DE AZEVEDO",
				"ts": "TERMINAL LAPA"
			}
		]
	`

	w.Write([]byte(jsonString))
}

func handleQueryStops(w http.ResponseWriter, req *http.Request) {
	if req.FormValue("termosBusca") == "" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		[
		  {
			"cp": 340015329,
			"np": "AFONSO BRAZ B/C1",
			"ed": "R ARMINDA/ R BALTHAZAR DA VEIGA",
			"py": -23.592938,
			"px": -46.672727
		  }
		]
	`

	w.Write([]byte(jsonString))
}

func handleQueryStopsByLine(w http.ResponseWriter, req *http.Request) {
	lineCode := req.FormValue("codigoLinha")
	if lineCode == "" || lineCode == "0" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		[
		  {
			"cp": 340015329,
			"np": "AFONSO BRAZ B/C1",
			"ed": "R ARMINDA/ R BALTHAZAR DA VEIGA",
			"py": -23.592938,
			"px": -46.672727
		  }
		]
	`

	w.Write([]byte(jsonString))
}

func handleQueryStopsByCorridor(w http.ResponseWriter, req *http.Request) {
	corridorCode := req.FormValue("codigoCorredor")
	if corridorCode == "" || corridorCode == "0" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		[
		  {
			"cp": 340015329,
			"np": "AFONSO BRAZ B/C1",
			"ed": "R ARMINDA/ R BALTHAZAR DA VEIGA",
			"py": -23.592938,
			"px": -46.672727
		  }
		]
	`

	w.Write([]byte(jsonString))
}

func handleCorridors(w http.ResponseWriter, req *http.Request) {
	jsonString := `
		[
		  {
			"cc":8,
			"nc":"Campo Limpo"
		  }
		]
	`

	w.Write([]byte(jsonString))
}

func handleCompanies(w http.ResponseWriter, req *http.Request) {
	jsonString := `
		[
		  {
			"hr":"11:20",
			"e": [
			  {
				"a": 1,
				"e": [
				  {
					"a": 1,
					"c": 999,
					"n": "NOME"
				  }
				]
			  }
			]
		  }
		]
	`

	w.Write([]byte(jsonString))
}

func handlePositions(w http.ResponseWriter, req *http.Request) {
	jsonString := `
		{
		  "hr": "11:30",
		  "l": [
			{
			  "c": "5015-10",
			  "cl": 33887,
			  "sl": 2,
			  "lt0": "METRÔ JABAQUARA",
			  "lt1": "JD. SÃO JORGE",
			  "qv": 1,
			  "vs": [
				{
				  "p":68021,
				  "a":true,
				  "ta":"2017-05-12T14:30:37Z",
				  "py":-23.678712500000003,
				  "px":-46.65674
				}
			  ]
			}
		  ]
		}
	`

	w.Write([]byte(jsonString))
}

func handleLinePositions(w http.ResponseWriter, req *http.Request) {
	lineCode := req.FormValue("codigoLinha")
	if lineCode == "" || lineCode == "0" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		{
		  "hr": "19:57",
		  "vs": [
			{
			  "p": 11433,
			  "a": false,
			  "ta": "2017-05-07T22:57:02Z",
			  "py": -23.540150375000003,
			  "px": -46.64414075
			}
		  ]
		}
	`

	w.Write([]byte(jsonString))
}

func handleGaragePositions(w http.ResponseWriter, req *http.Request) {
	companyCode := req.FormValue("codigoEmpresa")
	if companyCode == "" || companyCode == "0" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		{
		  "hr": "11:30",
		  "l": [
			{
			  "c": "5015-10",
			  "cl": 33887,
			  "sl": 2,
			  "lt0": "METRÔ JABAQUARA",
			  "lt1": "JD. SÃO JORGE",
			  "qv": 1,
			  "vs": [
				{
				  "p":68021,
				  "a":true,
				  "ta":"2017-05-12T14:30:37Z",
				  "py":-23.678712500000003,
				  "px":-46.65674
				}
			  ]
			}
		  ]
		}
	`

	w.Write([]byte(jsonString))
}

func handlePrevisions(w http.ResponseWriter, req *http.Request) {
	stopCode := req.FormValue("codigoParada")
	if stopCode == "" || stopCode == "0" {
		w.Write([]byte(`[]`))
		return
	}

	lineCode := req.FormValue("codigoLinha")
	if lineCode == "" || lineCode == "0" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		{
		  "hr": "20:09",
		  "p": {
			"cp": 4200953,
			"np": "PARADA ROBERTO SELMI DEI B/C",
			"py": -23.675901,
			"px": -46.752812,
			"l": [
			  {
				"c": "7021-10",
				"cl": 1989,
				"sl": 1,
				"lt0": "TERM. JOÃO DIAS",
				"lt1": "JD. MARACÁ",
				"qv": 1,
				"vs": [
				  {
					"p": 74558,
					"t": "23:11",
					"a": true,
					"ta": "2017-05-07T23:09:05Z",
					"py": -23.67603,
					"px": -46.75891166666667
				  }
				]
			  }
			]
		  }
		}
	`

	w.Write([]byte(jsonString))
}

func handleLinePrevisions(w http.ResponseWriter, req *http.Request) {
	lineCode := req.FormValue("codigoLinha")
	if lineCode == "" || lineCode == "0" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		{
		  "hr": "20:18",
		  "ps": [
			{
			  "cp": 700016623,
			  "np": "ANA CINTRA B/C",
			  "py": -23.538763,
			  "px": -46.646925,
			  "vs": [
				{
				  "p": "11436",
				  "t": "23:26",
				  "a": false,
				  "ta": "2017-05-07T23:18:02Z",
				  "py": -23.528119999999998,
				  "px": -46.670674999999996
				}
			  ]
			}
		  ]
		}
	`

	w.Write([]byte(jsonString))
}

func handleStopPrevisions(w http.ResponseWriter, req *http.Request) {
	stopCode := req.FormValue("codigoParada")
	if stopCode == "" || stopCode == "0" {
		w.Write([]byte(`[]`))
		return
	}

	jsonString := `
		{
		  "hr": "20:20",
		  "p": {
			"cp": 4200953,
			"np": "PARADA ROBERTO SELMI DEI B/C",
			"py": -23.675901,
			"px": -46.752812,
			"l": [
			  {
				"c": "675K-10",
				"cl": 198,
				"sl": 1,
				"lt0": "METRO STA CRUZ",
				"lt1": "TERM. JD. ANGELA",
				"qv": 1,
				"vs": [
				  {
					"p": 73651,
					"t": "23:22",
					"a": true,
					"ta": "2017-05-07T23:20:06Z",
					"py": -23.676623333333335,
					"px": -46.757641666666665
				  }
				]
			  }
			]
		  }
		}
	`

	w.Write([]byte(jsonString))
}
