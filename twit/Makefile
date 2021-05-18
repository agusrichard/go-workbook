run-dev:
	docker-compose up --remove-orphans

down-dev:
	docker-compose down

run-test:
	docker-compose -f docker-compose-test.yaml --env-file test.env up --remove-orphans

down-test:
	docker-compose -f docker-compose-test.yaml --env-file test.env down

check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml