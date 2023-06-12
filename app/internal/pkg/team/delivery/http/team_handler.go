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

type teamHandler struct {
	teamUsecase models.TeamUsecaseI
}

func NewTeamHandler(m *mux.Router, teamUsecase models.TeamUsecaseI) {
	handler := &teamHandler{
		teamUsecase: teamUsecase,
	}

	m.Handle("/api/teams", middleware.AuthMiddleware(http.HandlerFunc(handler.GetAll), "admin", "user")).Methods("GET")
	m.Handle("/api/teams/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.GetDriverById), "admin", "user")).Methods("GET")
	m.Handle("/api/teams_of_season", middleware.AuthMiddleware(http.HandlerFunc(handler.GetTeamsOfSeason), "admin", "user")).Methods("GET")
	m.Handle("/api/teams_of_season/{season}", middleware.AuthMiddleware(http.HandlerFunc(handler.GetTeamsOfSeason), "admin", "user")).Methods("GET")
	m.Handle("/api/teams", middleware.AuthMiddleware(http.HandlerFunc(handler.Create), "admin")).Methods("POST")
	m.Handle("/api/teams/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Update), "admin")).Methods("PUT")
	m.Handle("/api/teams/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Delete), "admin")).Methods("DELETE")
}

func (handler *teamHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	teams, err := handler.teamUsecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(teams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *teamHandler) GetDriverById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	team, err := handler.teamUsecase.GetTeamById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *teamHandler) GetTeamsOfSeason(w http.ResponseWriter, r *http.Request) {
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
	teams, err := handler.teamUsecase.GetTeamsOfSeason(season)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(teams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *teamHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	team := new(models.Team)
	err := decoder.Decode(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := handler.teamUsecase.Create(team)
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

func (handler *teamHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	team := new(models.Team)
	err = decoder.Decode(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.teamUsecase.Update(id, team)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *teamHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.teamUsecase.Delete(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
