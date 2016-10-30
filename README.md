File Share
===============

Run using docker-compose 
------------------------
https://docs.docker.com/compose/
To run docker container you must have two free ports: 8090, 8091

- docker-compose build
- docker-compose up

Now you can send api requests to localhost:8090  
To see REST api documentation please open localhost:8091  
API documentation written using swagger.io framework. You can send test requests from it.

To run project manually:
-----------------------
- install mongodb
- install golang
- create $GOPATH folder
- copy project to src folder
- cd to project folder
- execute `go get`
- execute `go test ./...`
- execute `go build`
- execute `./file_share`

Tests
-------
Project has 95.27% of coverage, you can see report in `coverage.html` file  
To calculate coverage:
- execute `go get github.com/axw/gocov/gocov`
- execute `go get -u gopkg.in/matm/v1/gocov-html`
- from project folder execute `gocov test ./...  | gocov-html > coverage.html `


