package main

import (
	"fmt"
	"log"
	"net/http"

	
)


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/eatdrink", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Title: Going out ,Position:52.51018 13.3754},{Title: The Curtain Club,Position: 52.5042 13.383}")
	}
	
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)

}