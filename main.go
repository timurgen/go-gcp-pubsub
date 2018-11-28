package main

import (
	"context"
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

	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./credentials.json")
	}

	ctx := context.Background()
	projID := os.Getenv("GCP_PROJECT_ID")

	PubsubClient, err = pubsub.NewClient(ctx, projID)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/{topic}", PublishMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
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
