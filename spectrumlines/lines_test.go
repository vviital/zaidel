package spectrumlines

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	peaks := Find(context.Background(), peaksOne, DefaultSettings())
	// fmt.Println("--- peaks ---", peaks)

	bts, err := json.Marshal(peaks)

	fmt.Println("--- err ---", err)
	fmt.Println(string(bts))

	assert.Equal(t, true, false)
}
