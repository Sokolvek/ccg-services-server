package storage

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var DB *firestore.Client

// a
func InitDB() {
	config := &firebase.Config{
		ProjectID: "ccg-services",
	}
	opt := option.WithCredentialsFile("ccg-services-firebase-adminsdk-i2vau-dc2d957f8b.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error init app: %v", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error init database client: %v", err)
	}
	DB = client
}
