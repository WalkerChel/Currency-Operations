app-build:
	docker compose build currency-operations-app

app-run:
	docker compose up currency-operations-app

pstg_docker_up:
	docker run --name=currency-operations-db -e POSTGRES_PASSWORD=qwerty -p 5432:5432 -d postgres

pstg_docker_down:
	docker stop currency-operations-db && docker rm currency-operations-db

migrations_up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

migrations_down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down 

migrations_force:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' force 000001