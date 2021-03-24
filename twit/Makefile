run-dev:
	docker-compose up --remove-orphans

down-dev:
	docker-compose down

run-test:
	docker-compose -f docker-compose-test.yaml --env-file test.env up --remove-orphans

down-test:
	docker-compose -f docker-compose-test.yaml --env-file test.env down