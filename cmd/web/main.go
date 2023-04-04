package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"fordycephotos.com/ui"
	"github.com/julienschmidt/httprouter"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

type PhotoData struct {
	FilmPaths   []string
	IphonePaths []string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Print("Request to /home")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	filmFilePaths, err := filepath.Glob("ui/static/img/film/*")
	strippedFilmPaths := make([]string, len(filmFilePaths))

	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	for i, path := range filmFilePaths {
		basepath := (filepath.Base(path))
		strippedFilmPaths[i] = basepath
	}

	sort.Sort(sort.Reverse(sort.StringSlice(strippedFilmPaths)))

	photoData := &PhotoData{
		FilmPaths: strippedFilmPaths,
	}

	ts, err := template.ParseFS(ui.Files, "html/pages/index.tmpl")
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "index.tmpl", photoData)
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	router.HandlerFunc(http.MethodGet, "/", app.home)

	infoLog.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", router)
	errorLog.Fatal(err)
}
