package peaks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name   string
		input  Spectrum
		output FinderResult
	}{
		{
			name:   "should process first spectrum",
			input:  SpectrumOne,
			output: ResultOne,
		},
		{
			name:   "should process second spectrum",
			input:  SpectrumTwo,
			output: ResultTwo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Find(tt.input.GetPoints(), NewDefaultSettings())

			assert.Equal(t, tt.output.PeaksCount, r.PeaksCount)
			assert.Equal(t, tt.output.Peaks, r.Peaks)
			assert.Equal(t, tt.output.BackgroundData, r.BackgroundData)
			assert.Equal(t, tt.output.SpectrumArea, r.SpectrumArea)
			assert.InDeltaSlice(t, tt.output.Destination, r.Destination, epsilon)
			assert.InDeltaSlice(t, tt.output.OriginalPeaksPositions, r.OriginalPeaksPositions, epsilon)

			// Compare wavelet data
			assert.Len(t, r.WaveletData, len(tt.output.WaveletData))
			for i, w := range r.WaveletData {
				assert.InDelta(t, tt.output.WaveletData[i].X, w.X, epsilon)
				assert.InDelta(t, tt.output.WaveletData[i].Y, w.Y, epsilon)
			}
		})
	}
}
