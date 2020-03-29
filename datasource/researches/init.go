package researches

import (
	"os"
)

var url string

func init() {
	url, _ = os.LookupEnv("RESEARCH_SERVICE_URL")
}
