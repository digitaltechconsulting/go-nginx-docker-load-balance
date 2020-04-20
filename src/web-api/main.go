package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	port := ":4040"
	http.HandleFunc("/", ProcessRequest)
	err := http.ListenAndServe(port, nil)
	fmt.Println("Starting server...")
	fmt.Println(err)
}

func ProcessRequest(response http.ResponseWriter, request *http.Request) {
	dockerInstanceId := os.Getenv("DOCKER_INSTANCE_ID")
	message := "Request is served by container"

	if len(strings.TrimSpace(dockerInstanceId)) <= 0 {
		dockerInstanceId = "Unknown"
	}
	message = fmt.Sprintf("%s : %s", message, dockerInstanceId)
	response.Write([]byte(message))
}
