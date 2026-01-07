up:
	docker compose up

certs-install:
	script/install-local-certs.sh

lint-frontend:
	cd src/frontend && npx prettier . --write

lint-api:
	cd src/api && go fmt

lint: lint-frontend lint-api
