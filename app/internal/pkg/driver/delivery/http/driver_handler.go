package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"app/internal/pkg/models"
)

type driverHandler struct {
	driverUsecase models.DriverUsecaseI
}

func NewDriverHandler(m *mux.Router, driverUsecase models.DriverUsecaseI) {
	handler := &driverHandler{
		driverUsecase: driverUsecase,
	}

	m.HandleFunc("/drivers", handler.Create).Methods("POST")
	m.HandleFunc("/drivers", handler.GetAll).Methods("GET")
	m.HandleFunc("/drivers/{id}", handler.GetDriverById).Methods("GET")
	m.HandleFunc("/drivers/{id}", handler.Update).Methods("PUT")
	m.HandleFunc("/drivers/{id}", handler.Delete).Methods("DELETE")
}

func (handler *driverHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	drivers, err := handler.driverUsecase.GetAll()
	if err != nil {
		http.Error(w, `db err`, http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(drivers)
	if err != nil {
		http.Error(w, `encode err`, http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) GetDriverById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `id err`, http.StatusInternalServerError)
		return
	}
	driver, err := handler.driverUsecase.GetDriverById(id)
	if err != nil {
		http.Error(w, `db err`, http.StatusInternalServerError)
		return
	}
	err = encoder.Encode(driver)
	if err != nil {
		http.Error(w, `encode err`, http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	driver := new(models.Driver)
	err := decoder.Decode(driver)
	if err != nil {
		http.Error(w, `Form err`, http.StatusBadRequest)
		return
	}
	fmt.Printf("%#v\n", driver)
	_, err = handler.driverUsecase.Create(driver)
	if err != nil {
		fmt.Println(err)
		http.Error(w, `db err:`, http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `id err`, http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	driver := new(models.Driver)
	err = decoder.Decode(driver)
	if err != nil {
		http.Error(w, `Form err`, http.StatusBadRequest)
		return
	}
	err = handler.driverUsecase.Update(id, driver)
	if err != nil {
		fmt.Println(err)
		http.Error(w, `db err:`, http.StatusInternalServerError)
		return
	}
}

func (handler *driverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `id err`, http.StatusInternalServerError)
		return
	}
	err = handler.driverUsecase.Delete(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, `db err:`, http.StatusInternalServerError)
		return
	}
}
