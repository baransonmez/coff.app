package web

import (
	"encoding/json"
	"github.com/baransonmez/coff.app/business/common"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Param returns the web call parameters from the request.
func Param(r *http.Request, key string) string {
	m := httptreemux.ContextParams(r.Context())
	return m[key]
}

// Decode reads the body of an HTTP request looking for a JSON document. The
// body is decoded into the provided value.
//
// If the provided value is a struct then it is checked for validation tags.
func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}

func ReadIDParam(r *http.Request) (uuid.UUID, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := common.StringToID(params.ByName("id"))
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil

}
