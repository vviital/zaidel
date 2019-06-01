package spectrumlines

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Run("first peaks json", func(t *testing.T) {
		spectrumLines := Find(context.Background(), peaksOne, DefaultSettings())
		assert.Equal(t, spectrumLinesOne, spectrumLines)
	})

	t.Run("second peaks json", func(t *testing.T) {
		spectrumLines := Find(context.Background(), peaksTwo, DefaultSettings())
		assert.Equal(t, spectrumLinesTwo, spectrumLines)
	})
}
