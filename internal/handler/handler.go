package handler

import (
	"encoding/json"
	"net/http"
	model "nliSessionToken/internal/models"
	"time"

	"github.com/gorilla/mux"
)

func (c *Controller) GetTokenHandler(w http.ResponseWriter, r *http.Request) {

	var getTokenRequest model.GetTokenRequest
	// Vars returns the route variables for the current request, if any.
	vars := mux.Vars(r)
	deviceID := vars["device_id"]
	if err := json.NewDecoder(r.Body).Decode(&getTokenRequest); err != nil {
		c.Fail(w, http.StatusBadRequest, "failed to decode json", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//get token
	nliToken, err := c.service.GenerateToken(SetName, getTokenRequest.DeviceID)
	if err != nil {
		c.Fail(w, http.StatusBadRequest, "failed to get token", err)
	}
	getTokenResponse := model.GetTokenResponse{
		CreatedAt: time.Now().String(),
		NLIToken:  nliToken,
		DeviceID:  deviceID,
	}

	json.NewEncoder(w).Encode(&getTokenResponse)
	c.Success(w, http.StatusOK, nil)
}

/*var err error
getTokenRequest.DeviceID, err = r.FormValue("device_id")

if err != nil {
	customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
	return
}
w.Header().Set("Content-Type", "application/json")
//get token

getTokenResponse := GetTokenResponse{
	CreatedAt: time.Now().String(),
	NLIToken:  " ",
	DeviceID:  getTokenRequest.DeviceID,
}

json.NewEncoder(w).Encode(&getTokenResponse)*/
