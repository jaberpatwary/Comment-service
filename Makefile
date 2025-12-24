include .env
export

start:
	@go run src/main.go

migration-%:
	@migrate create -ext sql -dir src/database/migrations create-table-$(subst :,_,$*)

migrate-up:
	@migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path src/database/migrations up

migrate-down:
	@migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path src/database/migrations down
