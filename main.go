package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

type service_struct struct {
	Service string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", LoadService).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
	fmt.Println("started")

}

func LoadService(w http.ResponseWriter, req *http.Request) {

	service := new(service_struct)
	json.NewDecoder(req.Body).Decode(service)

	w.WriteHeader(http.StatusOK)

	fmt.Println("service : " + service.Service)
	out, err := exec.Command("sh", "-c", "scripts/"+service.Service).Output()
	if err != nil {
		w.Write([]byte("Ca plante!\n"))
	}
	fmt.Println(out)
	w.Write(out)
}
