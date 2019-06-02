package peaks

import (
	"context"
	"time"

	"github.com/vviital/zaidel/peaks"
)

// PeaksModel struct
type PeaksModel struct {
	peaks.FinderResult
	Settings peaks.PeakSearchSettings `json:"settings"`
	ID       string                   `json:"id"`
	OwnerID  string                   `json:"ownerId"`
	Created  time.Time                `json:"created"`
	Updated  time.Time                `json:"updated"`
}

// Peaks interface
type Peaks interface {
	FindByID(context.Context, string) (PeaksModel, error)
	InsertOne(context.Context, peaks.FinderResult, peaks.PeakSearchSettings, string) (PeaksModel, error)
}
