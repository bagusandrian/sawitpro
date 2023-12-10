package impl

import (
	"net/http"
	"strconv"

	"github.com/bagusandrian/sawitpro/helper"
	m "github.com/bagusandrian/sawitpro/model"
)

func (h *handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		// request  m.RequestLogin
		response m.ResponseGetProfile
	)
	tokenString := helper.ExtractTokenFromHeader(r)
	claims, err := helper.VerifyToken(tokenString, []byte(h.config.Server.JWTSecretKey))
	if err != nil {
		helper.SendValidationErrorResponse(w, http.StatusForbidden, []string{err.Error()})
		return
	}
	userIDstr, err := claims.GetIssuer()
	if err != nil {
		helper.SendValidationErrorResponse(w, http.StatusForbidden, []string{err.Error()})
		return
	}
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		helper.SendValidationErrorResponse(w, http.StatusForbidden, []string{err.Error()})
		return
	}
	response, err = h.dbRepository.GetUserDataByID(r.Context(), userID)
	if err != nil {
		helper.SendValidationErrorResponse(w, http.StatusInternalServerError, []string{err.Error()})
		return
	}
	helper.SendJSONResponse(w, http.StatusCreated, map[string]interface{}{"data": response})
}
