package http

import (
	"app/internal/pkg/auth"
	"app/internal/pkg/auth/token"
	"app/internal/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type authHandler struct {
	userUsecase models.UserUsecaseI
}

func NewAuthHandler(m *mux.Router, userUsecase models.UserUsecaseI) {
	handler := &authHandler{
		userUsecase: userUsecase,
	}
	m.HandleFunc("/auth/login", handler.LogIn).Methods("POST")
	m.HandleFunc("/auth/register", handler.Register).Methods("POST")
	m.HandleFunc("/auth/logout", handler.LogOut).Methods("DELETE")
}

func (handler *authHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	loginData := new(auth.LogInData)
	err := decoder.Decode(loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok, err := handler.userUsecase.Authenticate(loginData.Login, loginData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}
	user, err := handler.userUsecase.GetUserByLogin(loginData.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tokenString, err := token.GenerateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:   "jwt-token",
		Value:  tokenString,
		MaxAge: 60 * 60 * 24,
		Path:   "/",
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("Logged in successfully!"))
}

func (handler *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := new(models.User)
	err := decoder.Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Role = "user"
	id, err := handler.userUsecase.Create(user)
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

func (handler *authHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "jwt-token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("Logged out successfully!"))
}
