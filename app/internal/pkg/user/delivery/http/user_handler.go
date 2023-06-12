package http

import (
	middleware "app/internal/app/middleware"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"app/internal/pkg/models"
)

type userHandler struct {
	userUsecase models.UserUsecaseI
}

func NewUserHandler(m *mux.Router, userUsecase models.UserUsecaseI) {
	handler := &userHandler{
		userUsecase: userUsecase,
	}

	m.Handle("/api/users", middleware.AuthMiddleware(http.HandlerFunc(handler.Create), "admin")).Methods("POST")
	m.Handle("/api/users/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Update), "admin")).Methods("PUT")
	m.Handle("/api/users/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.Delete), "admin")).Methods("DELETE")
}

func (handler *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := new(models.User)
	err := decoder.Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := handler.userUsecase.Create(user)
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

func (handler *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	user := new(models.User)
	err = decoder.Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.userUsecase.Update(id, user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.userUsecase.Delete(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
