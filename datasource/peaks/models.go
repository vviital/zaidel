package peaks

import (
	"context"

	"github.com/vviital/zaidel/peaks"
)

// PeaksModel struct
type PeaksModel struct {
	peaks.FinderResult
	Settings peaks.PeakSearchSettings `json:"settings"`
	ID       string                   `json:"id"`
	OwnerID  string                   `json:"ownerId"`
}

// Peaks interface
type Peaks interface {
	FindByID(context.Context, string) (PeaksModel, error)
	InsertOne(context.Context, peaks.FinderResult, peaks.PeakSearchSettings, string) (PeaksModel, error)
}
