package controller

import (
	"github.com/RangelReale/osin"
	ex "github.com/RangelReale/osin/example"
	"net/http"
)

// ex.NewTestStorage implements the "osin.Storage" interface
var server = osin.NewServer(osin.NewServerConfig(), ex.NewTestStorage())

// Access token endpoint
func CreateToken(w http.ResponseWriter, r *http.Request) {
	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		ar.Authorized = true
		server.FinishAccessRequest(resp, r, ar)
	}
	osin.OutputJSON(resp, w, r)
}

// Authorize via OSIN
func Authorize(w http.ResponseWriter, r *http.Request) {
	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {

		// HANDLE LOGIN PAGE HERE

		ar.Authorized = true
		server.FinishAuthorizeRequest(resp, r, ar)
	}
	osin.OutputJSON(resp, w, r)
}
