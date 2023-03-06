package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Print("Request to /home")
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	app.infoLog.Printf("%d resp from /home", http.StatusOK)
}

func main() {
	router := httprouter.New()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	router.HandlerFunc(http.MethodGet, "/", app.home)

	infoLog.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", router)
	errorLog.Fatal(err)
}
