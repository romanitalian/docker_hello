package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const AppPortDefault = "80"
const TimeFormat = "2006-01-02T15:04:05Z"

var AppPort = os.Getenv("APP_PORT")



func main() {
	if AppPort == "" {
		AppPort = AppPortDefault
	}
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// w.Write([]byte("Current time: " + time.Now().Format(TimeFormat)))
		content, err := ioutil.ReadFile("app/static/style.css")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(content)
		log.Printf("User-Agent: %s", r.Header.Get("User-Agent"))
	})
	log.Printf("app port: %s\n", AppPort)

	http.ListenAndServe(":" + AppPort, nil)
}
