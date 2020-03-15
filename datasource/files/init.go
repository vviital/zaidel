package files

import "os"

var url string

func init() {
	url, _ = os.LookupEnv("FILES_SERVICE_URL")
}
