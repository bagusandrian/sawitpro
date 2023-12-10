package impl

import (
	"errors"
	"net/http"
	"strings"

	"github.com/bagusandrian/sawitpro/helper"
	m "github.com/bagusandrian/sawitpro/model"
)

func (h *handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		request  m.RequestUpdateProfile
		response m.ResponseUpdateProfile
	)

	if err := request.ParseFromHTTPRequest(r, []byte(h.config.Server.JWTSecretKey)); err != nil {
		helper.SendValidationErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	response, err = h.dbRepository.UpdateProfile(r.Context(), request)
	if err != nil {
		var statusCode int
		if strings.Contains(err.Error(), "duplicate") {
			statusCode = http.StatusConflict
			err = errors.New("phone number already used")
		} else {
			statusCode = http.StatusInternalServerError
		}
		helper.SendValidationErrorResponse(w, statusCode, []string{err.Error()})
		return
	}
	helper.SendJSONResponse(w, http.StatusCreated, map[string]interface{}{"data": response})
}
