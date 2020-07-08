package main

import (
	"fmt"
	"log"
	"net/http"

	
)


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/charging", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Title: Meijer  - Tesla Super charger ,Position: 41.75203 -87.590263},{Title: Jns / Sears 79th st 1 ,Position:41.752062 -87.590274}")
	}
	
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)

}