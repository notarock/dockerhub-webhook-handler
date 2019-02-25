package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
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

type Callback struct {
	State       string
	Description string
	Context     string
	Target_url  string
}

const owner = "notarock"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", LoadService).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func LoadService(w http.ResponseWriter, req *http.Request) {
	var hook Webhook

	var responseData Callback
	responseData.Context = "Gopdater répond"
	responseData.Description = "Yup "
	responseData.State = "error"
	responseData.Target_url = "https://testingtaskdjasd.domainename.com"

	content, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(content, &hook)

	if IsValidRequest(hook) {
		err := UpdateContainer(hook.Repository)
		if err != nil {
			fmt.Print(err.Error())
		} else {
			responseData.State = "success"
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func UpdateContainer(repository Repository) error {
	exec.Command("docker", "pull", repository.Repo_name).Run()
	exec.Command("docker", "stop", repository.Name).Run()
	exec.Command("docker", "rm", repository.Name).Run()
	err := exec.Command("docker", "run", "-d", "--name="+repository.Name, repository.Repo_name).Run()
	if err != nil {
		return err
	}

	return nil
}

func IsValidRequest(hook Webhook) bool {
	return hook.Repository.Owner == owner && strings.Contains(hook.Repository.Repo_name, owner)
}
