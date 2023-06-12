package http

import (
	"app/internal/app/middleware"
	"app/internal/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type grandPrixHandler struct {
	grandPrixUsecase  models.GrandPrixUsecaseI
	raceResultUsecase models.RaceResultUsecaseI
	qualResultUsecase models.QualResultUsecaseI
}

func NewDriverHandler(m *mux.Router,
	grandPrixUsecase models.GrandPrixUsecaseI,
	raceResultUsecase models.RaceResultUsecaseI,
	qualResultUsecase models.QualResultUsecaseI) {
	handler := &grandPrixHandler{
		grandPrixUsecase:  grandPrixUsecase,
		raceResultUsecase: raceResultUsecase,
		qualResultUsecase: qualResultUsecase,
	}

	m.HandleFunc("/api/grandprix", handler.GetAll).Methods("GET")
	m.HandleFunc("/api/grandprix/id/{id}", handler.GetGPById).Methods("GET")
	m.HandleFunc("/api/grandprix/{id}", handler.GetGPById).Methods("GET")
	m.HandleFunc("/api/grandprix/season/{season}", handler.GetAllBySeason).Methods("GET")
	m.HandleFunc("/api/grandprix/place/{place}", handler.GetAllByPlace).Methods("GET")
	m.Handle("/api/grandprix", middleware.AuthMiddleware(http.HandlerFunc(handler.Create), "admin")).Methods("POST")
	m.Handle("/api/grandprix/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Update), "admin")).Methods("PUT")
	m.Handle("/api/grandprix/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Delete), "admin")).Methods("DELETE")
	m.Handle("/api/grandprix/{id}/race_results", middleware.AuthMiddleware(http.HandlerFunc(handler.GetRaceResultsOfGP), "admin", "user")).Methods("GET")
	m.Handle("/api/grandprix/{id}/qual_results", middleware.AuthMiddleware(http.HandlerFunc(handler.GetQualResultsOfGP), "admin", "user")).Methods("GET")
	m.HandleFunc("/api/grandprix/{id}/race_winner", handler.GetRaceWinnerOfGP).Methods("GET")
}

func (handler *grandPrixHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	gp, err := handler.grandPrixUsecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(gp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hander *grandPrixHandler) GetGPById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	gp, err := hander.grandPrixUsecase.GetGPById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(gp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hander *grandPrixHandler) GetAllBySeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	season, err := strconv.Atoi(vars["season"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	gp, err := hander.grandPrixUsecase.GetAllBySeason(season)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(gp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hander *grandPrixHandler) GetAllByPlace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	place := vars["place"]
	encoder := json.NewEncoder(w)
	gp, err := hander.grandPrixUsecase.GetAllByPlace(place)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(gp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *grandPrixHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	gp := new(models.GrandPrix)
	err := decoder.Decode(gp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := handler.grandPrixUsecase.Create(gp)
	if err != nil {
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

func (handler *grandPrixHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	gp := new(models.GrandPrix)
	err = decoder.Decode(gp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.grandPrixUsecase.Update(id, gp)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *grandPrixHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.grandPrixUsecase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *grandPrixHandler) GetRaceResultsOfGP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	raceResults, err := handler.raceResultUsecase.GetRaceResultsOfGP(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(raceResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *grandPrixHandler) GetQualResultsOfGP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	qualResults, err := handler.qualResultUsecase.GetQualResultsOfGP(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(qualResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *grandPrixHandler) GetRaceWinnerOfGP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	qualResults, err := handler.raceResultUsecase.GetRaceWinnerOfGP(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(qualResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
