go test -covermode=count -coverprofile=count.out fmt 

go tool cover -func=count.out
go tool cover -html=count.out


Environment:

MONGODB_URL - the URL to the MongoDB instance
FILES_SERVICE_URL - the URL to the file service
RESEARCH_SERVICE_URL - the URL to the research service
