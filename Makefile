run_local:
	go run main.go

run_migrate:
	go run ./db/migrate/migrate.go

run_migrate_down:
	go run ./db/migrate/migrate_down.go

run_docker:
	docker stop productapiserver || true && docker rm productapiserver || true
	docker build --tag productapiserver-api:local .
	docker run --name productapiserver -d -p 3000:3000 productapiserver-api:local
