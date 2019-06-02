package peaks

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/google/uuid"
	"github.com/vviital/zaidel/peaks"
	"go.mongodb.org/mongo-driver/mongo"
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
	now := time.Now().UTC()

	model.FinderResult = obj
	model.Settings = settings
	model.OwnerID = ownerID
	model.ID = id.String()
	model.Created = now
	model.Updated = now
	_, err = c.collection.InsertOne(ctx, model)
	if err != nil {
		model = PeaksModel{}
		return
	}
	return
}

func (c *mongoPeaks) Update(ctx context.Context, obj PeaksModel, id string) (model PeaksModel, err error) {
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"updated":                 time.Now(),
				"finderresult.peaks":      obj.Peaks,
				"finderresult.peakscount": len(obj.Peaks),
			},
		},
	}

	r, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return model, err
	}
	if r.ModifiedCount == 0 {
		return model, errors.New("Peaks record not found")
	}

	return c.FindByID(ctx, id)
}
