package main

import (
	//"fmt"
	//"encoding/json"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	debug bool = false
)

var header = template.Must(template.ParseFiles("./pages/header.html"))
var trailer = template.Must(template.ParseFiles("./pages/trailer.html"))

/*
APIError ...
*/
type appError struct {
	Status    string
	Code      int
	ErrorText string
}

/*
errorRespomse
errorResponse{Code: "1010", Info: "Index", Msg: "Error loading Index page",
			HTTPStatusCode: http.StatusInternalServerError, HTTPStatus: http.StatusText(http.StatusInternalServerError)}
		//	errOut, _ := json.Marshal(errResp
*/
type errorResponse struct {
	Code           string
	Info           string
	Msg            string
	HTTPStatusCode int
	HTTPStatus     string
}

type appHandler struct {
	*appError
	H func(*appError, http.ResponseWriter, *http.Request, int) (int, error)
}

// ServeHTTP requests
func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if debug {
		log.Print("DEBUG: ServerHTTP")
	}

	status, err := ah.H(ah.appError, w, r, http.StatusOK)
	apierr := appError{
		Status:    fmt.Sprint(err),
		Code:      status,
		ErrorText: http.StatusText(status),
	}

	if err != nil {
		switch status {

		case http.StatusUnauthorized:
			log.Print("StatusUnauthorised")
			//http.Redirect(w, r, "/neo4jlogin", http.StatusMovedPermanently)
			//		redirect(&apierr, w, r, status)
			break

		case http.StatusConflict:
			log.Print("StatusConfict")
			//		redirect(&apierr, w, r, status)
			break

		default:
			log.Print("Default")
			errorHandler(&apierr, w, r, status)
			break
		}
	}
	//	fmt.Println("Status: ", strconv.Itoa(status))
} // End of ServeHTTP

/*
httpNotFoundHandler ...
Generic handler to show 404 errors
*/
func httpNotFoundHandler() http.Handler {

	if debug {
		log.Print("DEBUG: httpNotFoundHandler")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Handler not found ...")
		w.WriteHeader(http.StatusNotFound)
		//tmpl, _ := template.New("").ParseFiles("./pages/neo4jheader.html", "./pages/neo4jtrailer.html", "./pages/index/neo4jindex.html")
		tmpl, _ := template.New("").ParseFiles("./pages/header.html", "./pages/trailer.html", "./pages/errors/notfound404.html")
		err := tmpl.ExecuteTemplate(w, "err404", nil)
		if err != nil {
			return
		}
		return
	})
}

/*
errorHandler
Invoked from within the Go web server code
*/
func errorHandler(a *appError, w http.ResponseWriter, r *http.Request, e int) (int, error) {

	if debug {
		log.Print("DEBUG: errorHandler")
	}

	errResp := errorResponse{}
	json.Unmarshal([]byte(a.Status), &errResp)

	tmpl, _ := template.New("").ParseFiles("./pages/header.html", "./pages/trailer.html", "./pages/errors/error.html")
	err := tmpl.ExecuteTemplate(w, "error", errResp) // was er
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNotFound, nil
}

func main() {

	log.Print("Football")
	port := os.Getenv("WEB_PORT")
	if port == "" {
		port = "8999"
	}

	debugval := os.Getenv("FOOTBALL_DEBUG")
	if debugval == "" {
		debug = false
	}
	if debugval != "true" {
		debug = false
	} else {
		debug = true
	}
	debug = true
	log.Print("Debug is set as: ", debug)

	log.Print("Football webserver listener - " + port)

	s := newSubRouter(port)
	log.Fatal(http.ListenAndServe(":"+port, s))

}

// neo4jIndexHandler
func footballIndexHandler(a *appError, w http.ResponseWriter, r *http.Request, s int) (int, error) {

	if debug {
		log.Print("DEBUG: footballIndexHandler")
	}

	tmpl, err := template.New("").ParseFiles("./pages/header.html", "./pages/trailer.html", "./pages/index/index.html")
	if err != nil {
		errResp := errorResponse{Code: "1010", Info: "Index", Msg: "Index page not found",
			HTTPStatusCode: http.StatusNotFound, HTTPStatus: http.StatusText(http.StatusNotFound)}
		errOut, _ := json.Marshal(errResp)
		err := errors.New(string(errOut))
		return http.StatusNotFound, err
	}

	err = tmpl.ExecuteTemplate(w, "home", nil)
	if err != nil {
		errResp := errorResponse{Code: "1010", Info: "Index", Msg: "Error loading Index page",
			HTTPStatusCode: http.StatusInternalServerError, HTTPStatus: http.StatusText(http.StatusInternalServerError)}
		errOut, _ := json.Marshal(errResp)
		err := errors.New(string(errOut))
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil

}

/*
Craete a new router
*/
func newSubRouter(port string) *mux.Router {

	if debug {
		log.Print("DEBUG: newSubRouter")
	}

	r := mux.NewRouter() // works with /page/{key}, but not css
	s := r.PathPrefix("/").Subrouter()
	a := &appError{}

	fs := http.FileServer(http.Dir("./style"))
	//http.Handle("/assets/", fs)

	pp := s.PathPrefix("/style/")
	pp.Handler(http.StripPrefix("/style/", fs))
	//s.Handle("/assets", appHandler{a, neo4jAssets})

	s.Handle("/", appHandler{a, footballIndexHandler})

	// Register the mux router with net/http
	//http.Handle("/", r)

	/*
		Login / Register
	*/
	//s.Handle("/neo4jlogin", appHandler{a, neo4jLoginHandler})
	//s.Handle("/neo4jloginuser", appHandler{a, neo4jLoginUserHandler})

	//s.Handle("/neo4jlogoff", appHandler{a, neo4jLogoffHandler})

	//s.Handle("/neo4jregister", appHandler{a, neo4jRegisterHandler})
	//s.Handle("/neo4jregisteruser", appHandler{a, neo4jRegisterUserHandler})

	/*
		Neo4j end-points
	*/
	s.Handle("/index", appHandler{a, footballIndexHandler})
	//s.Handle("/neo4jabout", appHandler{a, neo4jAbout})
	//s.Handle("/neo4japi", appHandler{a, neo4jAPIHandler})
	//s.Handle("/neo4jsearchresults", appHandler{a, neo4jSearchHandler})
	//s.Handle("/neo4jtableresults", appHandler{a, neo4jTableHandler})

	//s.Handle("/neo4jerror", appHandler{a, neo4jAPIErrorHandler})

	s.NotFoundHandler = httpNotFoundHandler()
	s.MethodNotAllowedHandler = httpNotFoundHandler()

	return s
}
