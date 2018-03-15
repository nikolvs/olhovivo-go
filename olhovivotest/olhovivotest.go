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
		if path == "/Login/Autenticar" {
			handleAuth(w, req)
			return
		}

		if !isAuthenticated(req) {
			w.Write([]byte(`{"Message": "Authorization has been denied for this request."}`))
			return
		}

		switch path {
		case "/Linha/Buscar":
			handleQueryLines(w, req)
		case "/Linha/BuscarLinhaSentido":
			handleQueryLinesByDirecction(w, req)
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
