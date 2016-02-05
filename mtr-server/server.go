package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

var mux *http.ServeMux
var db *sql.DB
var userW, keyW string

func init() {
	userW = os.Getenv("MTR_USER")
	keyW = os.Getenv("MTR_KEY")

	mux = http.NewServeMux()
	mux.HandleFunc("/field/locality", auth(localityHandler))
	mux.HandleFunc("/field/model", auth(modelHandler))
	mux.HandleFunc("/field/metric", auth(fieldMetricHandler))
	mux.HandleFunc("/field/site", auth(siteHandler))
}

func main() {
	var err error
	db, err = sql.Open("postgres",
		os.ExpandEnv("host=${DB_HOST} connect_timeout=30 user=${DB_USER} password=${DB_PASSWORD} dbname=mtr sslmode=disable"))
	if err != nil {
		log.Println("Problem with DB config.")
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(30)

	if err = db.Ping(); err != nil {
		log.Println("ERROR: problem pinging DB - is it up and contactable? 500s will be served")
	}

	go deleteOldMetrics()

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func auth(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT", "DELETE":
			if user, password, ok := r.BasicAuth(); ok && userW == user && keyW == password {
				f(w, r)
			} else {
				http.Error(w, "Access denied", http.StatusUnauthorized)
				return
			}
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	}
}

func deleteOldMetrics() {
	ticker := time.NewTicker(time.Minute * 5).C
	var err error

	for {
		select {
		case <-ticker:
			if _, err = db.Exec(`DELETE FROM field.metric WHERE time < now() - interval '28 days';`); err != nil {
				log.Print(err.Error())
			}
		}
	}
}