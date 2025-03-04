all-tests:
	go test ./...

tests-cover-html:
	go test -coverprofile=coverage.out ./...

view-tests-cover:
	go tool cover -html=coverage.out

docker-run:
	docker-compose up app

docker-stop:
	docker-compose down