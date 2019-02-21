package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type Push_data struct {
	Images    []string
	Pushed_at string
	Pusher    string
	Tag       string
}

type Webhook struct {
	Callback_url string
	Push_data    Push_data
	Repository   Repository
}

type Repository struct {
	Comment_count    string
	Date_created     string
	description      string
	Dockerfile       string
	Full_description string
	Is_official      string
	Is_private       string
	Is_trusted       string
	Name             string
	Namespace        string
	Owner            string
	Repo_name        string
	Repo_url         string
	Star_count       string
	Status           string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", LoadService).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))

	fmt.Println("started")
}

func LoadService(w http.ResponseWriter, req *http.Request) {
	var hook Webhook
	content, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(content, &hook)
	dockerCmd := exec.Command("docker", "run", "-d", "--name="+hook.Repository.Name, hook.Repository.Repo_name)

	err := dockerCmd.Run()

	if err != nil {
		fmt.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("200 - Container started!"))
	}
}
