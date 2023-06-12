package http

import (
	"app/internal/app/middleware"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"app/internal/pkg/models"
)

type raceResultHandler struct {
	raceResultUsecase models.RaceResultUsecaseI
}

func NewRaceResultHandler(m *mux.Router, raceResultUsecase models.RaceResultUsecaseI) {
	handler := &raceResultHandler{
		raceResultUsecase: raceResultUsecase,
	}

	m.Handle("/api/race_results", middleware.AuthMiddleware(http.HandlerFunc(handler.Create), "admin")).Methods("POST")
	m.Handle("/api/race_results", middleware.AuthMiddleware(http.HandlerFunc(handler.GetAll), "admin", "user")).Methods("GET")
	m.Handle("/api/race_results/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.GetRaceResultById), "admin", "user")).Methods("GET")
	m.Handle("/api/race_results/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Update), "admin")).Methods("PUT")
	m.Handle("/api/race_results/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Delete), "admin")).Methods("DELETE")
}

func (handler *raceResultHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	results, err := handler.raceResultUsecase.GetAll()
	if err != nil {
		//encoder.Encode(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *raceResultHandler) GetRaceResultById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	result, err := handler.raceResultUsecase.GetRaceResultById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *raceResultHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	result := new(models.RaceResult)
	err := decoder.Decode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := handler.raceResultUsecase.Create(result)
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

func (handler *raceResultHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	result := new(models.RaceResult)
	err = decoder.Decode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.raceResultUsecase.Update(id, result)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *raceResultHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.raceResultUsecase.Delete(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
