package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/gorilla/mux"
)

var (
	//PubsubClient pub sub cleint
	PubsubClient *pubsub.Client
)

// service entry point
func main() {
	var err error

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Starting service on port %s", port)

	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./credentials.json")
	}

	credentialsFileContent := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_CONTENT")
	if credentialsFileContent != "" {
		err = ioutil.WriteFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"), []byte(credentialsFileContent), 0750)
	}

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	projID := os.Getenv("GCP_PROJECT_ID")
	if projID == "" {
		log.Fatal("GCP project ID must be assigned to GCP_PROJECT_ID environment var.")
	}

	PubsubClient, err = pubsub.NewClient(ctx, projID)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/{topic}", PublishMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

//PublishMessage - publisher endpoint handler
func PublishMessage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s\n", r.Host)
	ctx := context.Background()
	params := mux.Vars(r)
	topicName := params["topic"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	t := PubsubClient.Topic(topicName)
	result := t.Publish(ctx, &pubsub.Message{
		Data: body,
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published a message; msg ID: %v\n", id)
}
