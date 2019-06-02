package peaks

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/google/uuid"
	"github.com/vviital/zaidel/peaks"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

type mongoPeaks struct {
	collection *mongo.Collection
}

// NewMongoPeaks returns mongodb related implementation of Peaks interface
func NewMongoPeaks() Peaks {
	return &mongoPeaks{collection}
}

func (c *mongoPeaks) FindByID(ctx context.Context, id string) (model PeaksModel, err error) {
	filter := bson.D{{"id", id}}
	err = collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		model = PeaksModel{}
	}
	return
}

func (c *mongoPeaks) InsertOne(ctx context.Context, obj peaks.FinderResult, settings peaks.PeakSearchSettings, ownerID string) (model PeaksModel, err error) {
	id, _ := uuid.NewRandom()
	model.FinderResult = obj
	model.Settings = settings
	model.OwnerID = ownerID
	model.ID = id.String()
	_, err = c.collection.InsertOne(ctx, model)
	if err != nil {
		model = PeaksModel{}
		return
	}
	return
}

func exit(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func initPeaksCollection() {
	collection = client.Database("zaidel").Collection("peaks")

}

func init() {
	var err error
	url, _ := os.LookupEnv("MONGODB_URL")
	// panic if cannot istablished connection to the MongoDB in 10 seconds
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
