go test -covermode=count -coverprofile=count.out fmt 

go tool cover -func=count.out
go tool cover -html=count.out
