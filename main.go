package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

func main() {
	godotenv.Load()
	godotenv.Load(".env.local")
	port, addr := os.Getenv("PORT"), os.Getenv("LISTEN_ADDR")
	if port == "" {
		port = "8080"
	}

	name := "Unnamed"
	name = os.Getenv("NAME")

	email, password := os.Getenv("EMAIL"), os.Getenv("PASSWORD")
	if email == "" || password == "" {
		panic("missing email or password")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// TODO
		decoder := json.NewDecoder(req.Body)
		var form ContactForm
		err := decoder.Decode(&form)
		fmt.Println(form)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		err = SendEmail(name, email, password, form)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("internal server error: %s", err.Error()),
				http.StatusInternalServerError,
			)
			return
		}
		http.Error(w, "ok", http.StatusOK)

	})

	listenAddr := net.JoinHostPort(addr, port)
	log.Printf("listening on %s", listenAddr)
	err := http.ListenAndServe(listenAddr, nil)
	log.Fatal(err)
}

type ContactForm struct {
	From    string
	Title   string
	Message string
}

func SendEmail(name, eMail, password string, form ContactForm) error {
	var err error
	e := email.NewEmail()
	e.From = eMail
	e.To = []string{eMail}
	e.Cc = []string{}
	e.Subject = fmt.Sprintf("Contact API %s - %s", name, form.Title)
	e.Text = []byte(fmt.Sprintf("- From: %s \n- Title: %s\n- Message:\n\t\t%s", form.From, form.Title, form.Message))
	err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", eMail, password, "smtp.gmail.com"))
	return err
}
