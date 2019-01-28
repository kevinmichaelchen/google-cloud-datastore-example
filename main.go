package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

type Task struct {
	Category        string
	Done            bool
	Priority        float64
	Description     string `datastore:",noindex"`
	PercentComplete float64
	Created         time.Time
}

func (t *Task) String() string {
	return fmt.Sprintf("[Task category=%s]", t.Category)
}

func main() {
	ctx := context.Background()

	log.Println("Creating datastore client")

	client := InitDatastore()

	log.Println("Created datastore client")

	//id := uuid.Must(uuid.Parse("dc874a6f-a4a0-42c0-a539-0d4ad450922a"))
	//k := datastore.NameKey("Task", id.String(), nil)
	k := datastore.IDKey("Task", 1, nil)

	task := &Task{
		Category:        "Personal",
		Done:            false,
		Priority:        4,
		Description:     "Learn Cloud Datastore",
		PercentComplete: 10.0,
		Created:         time.Now(),
	}

	log.Println("Inserting...")
	if _, err := client.Put(ctx, k, task); err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	var fetched Task
	log.Println("Fetching...")
	err := client.Get(ctx, k, &fetched)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	log.Println(fetched)
}

func getProjectID() string {
	projectID, ok := os.LookupEnv("DATASTORE_PROJECT_ID")
	if !ok {
		log.Fatal("Failed to specify Google Cloud Datastore Project ID")
	}
	return projectID
}

func InitDatastore() *datastore.Client {
	client, err := datastore.NewClient(context.Background(), getProjectID())
	if err != nil {
		log.Fatalf("Could not create datastore: %s", err.Error())
	}
	return client
}
