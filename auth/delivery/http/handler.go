package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/devstackq/go-clean/auth"
	"github.com/devstackq/go-clean/auth/models"
)

//DI - for example mock
type Handler struct {
	authUseCase auth.UseCase
}

//for example unit test; mock service
func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{authUseCase: useCase}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	input := &models.User{}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = json.Unmarshal(bytes, input)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	ctx := context.Background()
	err = h.authUseCase.SignUp(ctx, input)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(200)
	http.Redirect(w, r, "/signin", 302)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	input := &models.User{}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = json.Unmarshal(bytes, input)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	ctx := context.Background()
	res, err := h.authUseCase.SignIn(ctx, input.Username, input.Password)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	log.Print(res)
	//parseTOken
	w.WriteHeader(200)
}

//register handlers
func InitRoutes(useCase auth.UseCase) {
	hr := NewHandler(useCase)
	http.HandleFunc("/signup", hr.SignUp)
	http.HandleFunc("/signin", hr.SignIn)
}
