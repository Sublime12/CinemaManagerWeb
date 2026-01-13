COMPOSE := docker compose

up:
	$(COMPOSE) --profile dev up -d --build

watch: 
	$(COMPOSE) --profile dev up --build --watch

down:
	$(COMPOSE) --profile dev down

down-v:
	$(COMPOSE) down -v

test:
	$(COMPOSE) --profile test build api-test
	$(COMPOSE) --profile test run --rm api-test

test-clean:
	$(COMPOSE) --profile test down -v

restart:
	$(COMPOSE) down && $(COMPOSE) --profile dev up -d

logs:
	$(COMPOSE) logs -f

ps:
	$(COMPOSE) ps

db:
	$(COMPOSE) exec -it db psql -h localhost -U cinema_manager -p 5432

migrate:
	$(COMPOSE) exec api ./cinema_manager_seeder

rebuild:
	$(COMPOSE) --profile dev build --no-cache

clean:
	$(COMPOSE) --profile dev down -v

certs-install:
	script/install-local-certs.sh

lint-frontend:
	cd src/frontend && npx prettier . --write

lint-api:
	cd src/api && go fmt

lint: lint-frontend lint-api
