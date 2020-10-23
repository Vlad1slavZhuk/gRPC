package server

import (
	"encoding/json"
	"fmt"
	"gRPC/internal/pkg/auth"
	e "gRPC/internal/pkg/constErr"
	"gRPC/internal/pkg/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Login - Ready!TODO Add Validation JSON
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var acc data.Account
	// получение данных JSON account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		http.Error(w, e.DecodeError.Error(), http.StatusBadRequest)
		return
	}
	token, err := s.service.Login(&acc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Token", token)
	fmt.Fprint(w, "See in tab Headers.")
}

// SignUp - in Process...
func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	var acc data.Account
	// получение данных JSON account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		http.Error(w, "Error JSON!", 400)
		return
	}

	token, err := s.service.SignUp(&acc)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Token", token) // в Header вставить токен для примера
	fmt.Fprint(w, "See in headers refresh Token.")
}

//TODO
func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	//token, err := auth.GetTokenFromHeader(r)
	//if err != nil {
	//	http.Error(w, "Empty token. Login again and get token.", 400)
	//	return
	//}
	//if err = auth.VerifyToken(token); err != nil {
	//	http.Error(w, "No valid token.", 400)
	//	return
	//}
	//baseAcc, err := s.service.GetStorage().GetAccounts()
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusNotFound)
	//	return
	//}
	//acc, err := auth.ContainsToken(token, baseAcc)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusUnauthorized) // TODO
	//	return
	//}
	baseAcc, err := s.service.GetStorage().GetAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	acc, err := auth.AccountIdentification(r, baseAcc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acc.SetToken("")
	fmt.Fprint(w, "You Logout! Bye-bye...")
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------
	// Зона проверки токена
	baseAcc, err := s.service.GetStorage().GetAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = auth.AccountIdentification(r, baseAcc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	//-----------------------------------------------------------------
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil && id < 0 {
		http.Error(w, "Ошибка конвертирования в int.", http.StatusBadRequest) // TODO
		return
	}

	ad, err := s.service.Get(uint(id))
	if err != nil {
		http.Error(w, "Не найден.", http.StatusNotFound) // TODO
		return
	}

	js, err := json.Marshal(ad)
	if err != nil {
		http.Error(w, "Ошибка в маршале", 400) // TODO
		return
	}
	fmt.Fprint(w, string(js))
}

func (s *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------
	// Зона проверки токена
	baseAcc, err := s.service.GetStorage().GetAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_, err = auth.AccountIdentification(r, baseAcc)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//-----------------------------------------------------------------
	w.Header().Set("Content-Type", "application/json")
	base, err := s.service.GetAll()
	if base == nil && err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // TODO
		return
	}

	js, err := json.Marshal(base)
	if err != nil {
		http.Error(w, "Ошибка в маршале", 400) // TODO
		return
	}
	fmt.Fprint(w, string(js))
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------
	// Зона проверки токена
	baseAcc, err := s.service.GetStorage().GetAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_, err = auth.AccountIdentification(r, baseAcc)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//-----------------------------------------------------------------
	var ad data.Ad
	if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
		http.Error(w, "Error JSON!", 400)
		return
	}
	//TODO
	err = s.service.Add(&ad)
	if err != nil {
		http.Error(w, "Error to add.", 400)
		return
	}
	js, _ := json.Marshal(ad)
	fmt.Fprintf(w, "Create new Ad\n%v", string(js))
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------
	// Зона проверки токена и аккаунта
	baseAcc, err := s.service.GetStorage().GetAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_, err = auth.AccountIdentification(r, baseAcc)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//-----------------------------------------------------------------
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil && id < 0 {
		http.Error(w, "Ошибка конвертирования в int.", http.StatusBadRequest) // TODO
		return
	}
	if err := s.service.Delete(uint(id)); err != nil {
		http.Error(w, "Ошибка удаления.", http.StatusBadRequest) // TODO
		return
	}
	fmt.Fprint(w, "Delete")
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	//-----------------------------------------------------------------
	// Зона проверки токена
	baseAcc, err := s.service.GetStorage().GetAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_, err = auth.AccountIdentification(r, baseAcc)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//-----------------------------------------------------------------
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil && id < 0 {
		http.Error(w, "Ошибка конвертирования в int.", http.StatusBadRequest) // TODO
		return
	}

	var ad data.Ad
	if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
		http.Error(w, "Error JSON!", 400)
		return
	}

	if err = s.service.Update(uint(id), &ad); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Fprint(w, "Update")
}
