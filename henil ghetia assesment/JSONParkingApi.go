package main

import (
	"fmt"
	"log"
	"net/http"

	
)


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/parking", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Title: Sp+ Parking,Position:41.710332,-87.633473},{Title: State & Harrison Parking,Position:41.712735,-87.626251}")
	}
	
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)

}