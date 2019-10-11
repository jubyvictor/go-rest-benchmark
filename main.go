package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jubyvictor/go-benchmark/gobench"

	"github.com/gorilla/mux"
)

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user gobench.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Incorrect data format")
	}
	json.Unmarshal(reqBody, &user)
	user.LastUpdatedAt = time.Now().UnixNano() / 1e6
	json.NewEncoder(w).Encode(user)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/up", updateUser)
	log.Fatal(http.ListenAndServe(":8080", router))
}
