package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"f-license/config"
	"f-license/lcs"
	"f-license/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func GenerateLicense(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)

	var l lcs.License
	_ = json.Unmarshal(bytes, &l)

	err := l.Generate()
	if err != nil {
		logrus.WithError(err).Error("License couldn't be generated")
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = storage.LicenseHandler.AddIfNotExisting(&l)
	if err != nil {
		logrus.WithError(err).Error("License couldn't be stored")
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ReturnResponse(w, http.StatusOK, map[string]interface{}{
		"id":    l.ID.Hex(),
		"token": l.Token,
	})
}

func GetLicense(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var l lcs.License
	err := storage.LicenseHandler.GetByID(id, &l)
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ReturnResponse(w, 200, l)
}

func GetAllLicenses(w http.ResponseWriter, r *http.Request) {
	var licenses []*lcs.License
	err := storage.LicenseHandler.GetAll(&licenses)
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ReturnResponse(w, 200, licenses)
}

func ChangeLicenseActiveness(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	inactivate := strings.Contains(r.URL.Path, "/inactivate")

	err := storage.LicenseHandler.Activate(id, inactivate)
	if err != nil {
		logrus.WithError(err).Error("Error while activeness change")
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var message string

	if inactivate {
		message = "Inactivated"
	} else {
		message = "Activated"
	}

	ReturnResponse(w, 200, map[string]interface{}{
		"message": message,
	})
}

func VerifyLicense(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	var l lcs.License
	err := storage.LicenseHandler.GetByToken(token, &l)
	if err != nil {
		logrus.WithError(err).Error("Error while getting license")
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ok, err := l.IsLicenseValid(token)
	if err != nil {
		ReturnResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"valid":   false,
			"message": err.Error(),
		})

		return
	}

	ReturnResponse(w, 200, map[string]interface{}{
		"valid": ok,
	})
}

func DeleteLicense(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := storage.LicenseHandler.DeleteByID(id)
	if err != nil {
		logrus.WithError(err).Error("Error while deleting license")
		ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ReturnResponse(w, 200, map[string]interface{}{
		"message": "License successfully deleted",
	})
}

func Ping(w http.ResponseWriter, r *http.Request) {

}

func ReturnResponse(w http.ResponseWriter, statusCode int, resp interface{}) {
	bytes, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = fmt.Fprintf(w, string(bytes))
}

func ReturnError(w http.ResponseWriter, statusCode int, errMsg string) {
	resp := map[string]interface{}{
		"error": errMsg,
	}
	bytes, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = fmt.Fprintf(w, string(bytes))
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != config.Global.AdminSecret {
			ReturnResponse(w, http.StatusUnauthorized, map[string]interface{}{
				"message": "Authorization failed",
			})
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
