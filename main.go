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

type Webhook struct {
	CallbackURL string `json:"callback_url"`
	PushData    struct {
		Images   []string `json:"images"`
		PushedAt int      `json:"pushed_at"`
		Pusher   string   `json:"pusher"`
		Tag      string   `json:"tag"`
	} `json:"push_data"`
	Repository struct {
		CommentCount    int    `json:"comment_count"`
		DateCreated     int    `json:"date_created"`
		Description     string `json:"description"`
		Dockerfile      string `json:"dockerfile"`
		FullDescription string `json:"full_description"`
		IsOfficial      bool   `json:"is_official"`
		IsPrivate       bool   `json:"is_private"`
		IsTrusted       bool   `json:"is_trusted"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		Owner           string `json:"owner"`
		RepoName        string `json:"repo_name"`
		RepoURL         string `json:"repo_url"`
		StarCount       int    `json:"star_count"`
		Status          string `json:"status"`
	} `json:"repository"`
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

	responseData.State = string(content)

	json.Unmarshal(content, &hook)
	fmt.Print(hook.Repository)

	if IsValidRequest(hook) {
		err := UpdateContainer(hook)
		if err != nil {
			responseData.State = err.Error() + ""
		} else {
			responseData.State = "success"
		}
	} else {

	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func UpdateContainer(hook Webhook) error {
	err := exec.Command("docker", "stop", hook.Repository.Name).Run()
	if err != nil {
		return err
	}
	err = exec.Command("docker", "rm", hook.Repository.Name).Run()
	if err != nil {
		return err
	}

	err = exec.Command("docker-compose", "pull", hook.Repository.Name).Run()
	if err != nil {
		return err
	}
	err = exec.Command("docker-compose", "up", "-d", "--no-deps", "--build", hook.Repository.Name).Run()
	if err != nil {
		return err
	}

	return nil
}

func IsValidRequest(hook Webhook) bool {
	return hook.Repository.Owner == owner && strings.Contains(hook.Repository.RepoName, owner)
}
