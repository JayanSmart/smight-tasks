// This package drives the backend of this API and manages all interfacing with the database.
package engine

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE_CONNECTION_STRING string = "mongodb+srv://cluster-0.sfbug.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func Connect() {

	credentials := options.Credential{
		Username: "smight",
		Password: os.Getenv("MONGODBPWD"),
	}

	clientOptions := options.Client().ApplyURI(DATABASE_CONNECTION_STRING).SetAuth(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	session, err := client.StartSession()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("session: %v\n", session)

	sessions := fmt.Sprintf("Sessions in Progress: %d", client.NumberSessionsInProgress())
	log.Print(sessions)

	_ = client
}
