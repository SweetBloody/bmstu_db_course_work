package http

import (
	"app/internal/app/middleware"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	"app/internal/pkg/models"
)

type driverHandler struct {
	driverUsecase models.DriverUsecaseI
}

func NewDriverHandler(m *mux.Router, driverUsecase models.DriverUsecaseI) {
	handler := &driverHandler{
		driverUsecase: driverUsecase,
	}

	m.Handle("/api/drivers", middleware.AuthMiddleware(http.HandlerFunc(handler.GetAll), "admin", "user")).Methods("GET")
	m.Handle("/api/drivers", middleware.AuthMiddleware(http.HandlerFunc(handler.Create), "admin")).Methods("POST")
	m.Handle("/api/drivers/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.GetDriverById), "admin", "user")).Methods("GET")
	m.HandleFunc("/api/drivers_of_season", handler.GetDriversOfSeason).Methods("GET")
	m.Handle("/api/drivers_of_season/{season}", middleware.AuthMiddleware(http.HandlerFunc(handler.GetDriversOfSeason), "admin", "user")).Methods("GET")
	m.Handle("/api/drivers_standing", middleware.AuthMiddleware(http.HandlerFunc(handler.GetDriversStanding), "admin", "user")).Methods("GET")
	m.Handle("/api/drivers/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Update), "admin")).Methods("PUT")
	m.Handle("/api/drivers/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Delete), "admin")).Methods("DELETE")
	m.Handle("/api/drivers_teams", middleware.AuthMiddleware(http.HandlerFunc(handler.LinkDriverTeam), "admin")).Methods("POST")
}

func (handler *driverHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	drivers, err := handler.driverUsecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(drivers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) GetDriverById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	driver, err := handler.driverUsecase.GetDriverById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) GetDriversOfSeason(w http.ResponseWriter, r *http.Request) {
	var season int
	var err error
	vars := mux.Vars(r)
	if x, ok := vars["season"]; !ok {
		season = time.Now().Year() - 1
	} else {
		season, err = strconv.Atoi(x)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	encoder := json.NewEncoder(w)
	drivers, err := handler.driverUsecase.GetDriversOfSeason(season)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(drivers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) GetDriversStanding(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	standing, err := handler.driverUsecase.GetDriversStanding()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(standing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	driver := new(models.Driver)
	err := decoder.Decode(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := handler.driverUsecase.Create(driver)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	driver := new(models.Driver)
	err = decoder.Decode(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.driverUsecase.Update(id, driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.driverUsecase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) LinkDriverTeam(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	link := new(models.DriversTeams)
	err := decoder.Decode(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.driverUsecase.LinkDriverTeam(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
