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

			assert.Equal(t, tt.output.Peaks, r.Peaks)
		})
	}
}
