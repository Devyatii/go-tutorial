package verify

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/response"
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"net/http"
	"net/smtp"
)

type VerifyHandlerDeps struct {
	*configs.Config
}

type VerifyHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var payload VerifyRequest
		err := json.NewDecoder(request.Body).Decode(payload)
		if err != nil {
			response.Json(writer, err.Error(), 402)
			return
		}
		emailAddr := handler.Config.Email
		address := handler.Config.Address
		password := handler.Config.Password
		if emailAddr != payload.Email || password != payload.Password {
			response.Json(writer, err.Error(), 403)
			return
		}
		e := email.NewEmail()
		e.From = "Jordan Wright <test@gmail.com>"
		e.To = []string{"test@example.com"}
		e.Bcc = []string{"test_bcc@example.com"}
		e.Cc = []string{"test_cc@example.com"}
		e.Subject = "Awesome Subject"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		e.Send(fmt.Sprintf("%s:587", address), smtp.PlainAuth("", emailAddr, password, address))

	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
	}
}
