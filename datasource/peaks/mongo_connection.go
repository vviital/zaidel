package peaks

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func exit(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func initPeaksCollection() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	definitions := []struct {
		name       string
		background bool
		unique     bool
		fields     []string
		types      []interface{}
	}{
		{
			name:       "peaks_id_primary_key_index",
			fields:     []string{"id"},
			types:      []interface{}{"hashed"},
			background: true,
		},
		{
			name:       "peaks_owner_id_search_key_index",
			fields:     []string{"ownerid"},
			types:      []interface{}{1},
			background: true,
		},
	}

	collection = client.Database("zaidel").Collection("peaks")
	indexes := collection.Indexes()

	for _, d := range definitions {
		keys := bson.D{}
		for i := range d.fields {
			keys = append(keys, bson.E{Key: d.fields[i], Value: d.types[i]})
		}
		options := options.Index().SetBackground(d.background).SetUnique(d.unique).SetName(d.name)
		name, err := indexes.CreateOne(ctx, mongo.IndexModel{Keys: keys, Options: options})
		exit(err)
		log.Printf("%s index was added", name)
	}
}

func init() {
	var err error
	url, _ := os.LookupEnv("MONGODB_URL")
	// panic if cannot establish connection to the MongoDB in 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.NewClient(options.Client().ApplyURI(url))
	exit(err)

	err = client.Connect(ctx)
	exit(err)
	err = client.Ping(ctx, nil)
	exit(err)

	log.Println("Connected to the peaks mongoDB")
	initPeaksCollection()
}
