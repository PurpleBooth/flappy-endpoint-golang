package main

import (
	"net/http"
	"fmt"
	"math/rand"
	"time"
)

var threeSeconds time.Duration

func main() {
	threeSeconds, _ = time.ParseDuration("3s")
	http.HandleFunc("/", makeHandler(defaultPage))
	http.HandleFunc("/flappy", makeHandler(flappyPage))
	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}

func defaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Everything is okay! see /flappy for something that's less reliable")
}

func flappyPage(w http.ResponseWriter, r *http.Request) {
	switch rand.Intn(3) {
	case 0:
		http.Error(w, "NOT FOUND", 404)
		break
	case 1:
		http.Error(w, "Something terrible has happened", 500)
		break
	case 2:
		time.Sleep(threeSeconds)
		fmt.Fprint(w, "Slow response")
		break
	case 3:
		fmt.Fprint(w, "Normal response")
		break
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v %v\n", r.Method, r.RequestURI)
		w.Header().Set("Content-Type", "text/plain")

		fn(w, r)
	}
}
