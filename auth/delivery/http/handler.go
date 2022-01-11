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

//DI - fro mock
type Handler struct {
	authUseCase auth.UseCase
}

//for unit test; mock service
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

func InitRoutes(useCase auth.UseCase) {
	hr := NewHandler(useCase)
	http.HandleFunc("/signup", hr.SignUp) //register handler
	http.HandleFunc("/signin", hr.SignIn) //register handler
}

// [Interface]
// Address = 10.0.0.1/24
// SaveConfig = true
// PostUp = iptables -A FORWARD -i wg0 -j ACCEPT; iptables -t nat -A POSTROUTING ->
// PostDown = iptables -D FORWARD -i wg0 -j ACCEPT; iptables -t nat -D POSTROUTING>
// ListenPort = 3785
// PrivateKey = mAyHVBewjYa8zeEGL+Y5xkMyplVaLaev4FMuKshQx1A=

// [PEER]
// PublicKey = 5j8q9wnQjnDx31E9KarACMbriviBeI1mbBGCrWy+h2Q=
// AllowedIPs = 110.0.0.2/32
// [PEER]
// PublicKey = dDQ/C/Xd0xIpZ40dYjVWJ4m53ddVc/Z3jV/yUmGoF3s=
// AllowedIPs = 110.0.0.3/32
// [PEER]
// PublicKey = vBtxVLa6CeRki5+I7AHIbe4CJv2oKsKsyyqToKYHiGc=
// AllowedIPs = 110.0.0.4/32
