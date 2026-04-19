.PHONY: setup-local run-server env-example sqlc-gen mockery-gen install-goose migrate-sql migrate-up migrate-down

ENV_FILE=.env
MIGRATIONS_FOLDER=./data/migrations
GOOSE_VERSION=v3.27.0

setup-local: # setup-local-stripe
	echo "Setting up the local environment"
	docker compose -f ./docker/docker-compose.local.yml up -d

run-server:
	go run ./cmd/server/

env-example:
	echo "Generating .env.example from .env"
	awk -F'=' 'BEGIN {OFS="="} \
    	/^[[:space:]]*#/ {print; next} \
    	/^[[:space:]]*$$/ {print ""; next} \
    	NF>=1 {gsub(/^[[:space:]]+|[[:space:]]+$$/, "", $$1); print $$1"="}' .env > .env.example
	echo ".env.example generated successfully."

sqlc-gen:
	echo "Generating Go code from SQL queries using sqlc"
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc:1.30.0 generate

mockery-gen:
	echo "Generating mock implementations using mockery"
	docker run --rm -v $(PWD):/src -w /src vektra/mockery:v3.7.0

install-goose:
	echo "Installing the Goose database migration tool"
	go install github.com/pressly/goose/v3/cmd/goose@$(GOOSE_VERSION)
	ls "$(shell go env GOPATH)/bin/" | grep goose

migrate-sql:
	echo "Creating a new SQL migration with name: $(NAME)"
	goose -dir=$(MIGRATIONS_FOLDER) create $(NAME) sql

migrate-up:
	echo "Running database migrations up to the latest version"
	goose -env $(ENV_FILE) up

migrate-down:
	echo "Rolling back the last database migration"
	goose -env $(ENV_FILE) down

install-stripe-mock:
	echo "Installing the Stripe mock server"
	go install github.com/stripe/stripe-mock@latest

setup-local-stripe: install-stripe-mock
	echo "Setting up the local Stripe API"
	stripe-mock -http-port 12111

