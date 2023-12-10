package impl

import (
	"net/http"

	"github.com/bagusandrian/sawitpro/helper"
	m "github.com/bagusandrian/sawitpro/model"
)

func (h *handler) Registration(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		request  m.RequestRegistration
		response m.ResponseRegristration
	)
	errString := request.ParseFromHTTPRequest(r)
	if len(errString) > 0 {
		helper.SendValidationErrorResponse(w, http.StatusBadRequest, errString)
		return
	}
	response, err = h.dbRepository.Registration(r.Context(), request)
	if err != nil {
		helper.SendValidationErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
		return
	}
	helper.SendJSONResponse(w, http.StatusCreated, map[string]interface{}{"data": response})
}
