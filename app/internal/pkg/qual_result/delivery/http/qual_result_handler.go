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

type qualResultHandler struct {
	qualResultUsecase models.QualResultUsecaseI
}

func NewQualResultHandler(m *mux.Router, qualResultUsecase models.QualResultUsecaseI) {
	handler := &qualResultHandler{
		qualResultUsecase: qualResultUsecase,
	}

	m.Handle("/api/qual_results", middleware.AuthMiddleware(http.HandlerFunc(handler.Create), "admin")).Methods("POST")
	m.Handle("/api/qual_results", middleware.AuthMiddleware(http.HandlerFunc(handler.GetAll), "admin", "user")).Methods("GET")
	m.Handle("/api/qual_results/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.GetQualResultById), "admin", "user")).Methods("GET")
	m.Handle("/api/qual_results/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Update), "admin")).Methods("PUT")
	m.Handle("/api/qual_results/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Delete), "admin")).Methods("DELETE")
}

func (handler *qualResultHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	results, err := handler.qualResultUsecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *qualResultHandler) GetQualResultById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	result, err := handler.qualResultUsecase.GetQualResultById(id)
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

func (handler *qualResultHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	result := new(models.QualResult)
	err := decoder.Decode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := handler.qualResultUsecase.Create(result)
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

func (handler *qualResultHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	result := new(models.QualResult)
	err = decoder.Decode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.qualResultUsecase.Update(id, result)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *qualResultHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.qualResultUsecase.Delete(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
