package impl

import (
	"net/http"

	"github.com/bagusandrian/sawitpro/helper"
	m "github.com/bagusandrian/sawitpro/model"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		request  m.RequestLogin
		response m.ResponseLogin
	)
	errString := request.ParseFromHTTPRequest(r)
	if err != nil {
		helper.SendValidationErrorResponse(w, http.StatusBadRequest, errString)
		return
	}
	response, err = h.dbRepository.Login(r.Context(), request)
	if err != nil {
		helper.SendValidationErrorResponse(w, http.StatusInternalServerError, []string{err.Error()})
		return
	}
	helper.SendJSONResponse(w, http.StatusCreated, map[string]interface{}{"data": response})
}
