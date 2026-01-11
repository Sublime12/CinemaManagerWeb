COMPOSE := docker compose

up:
	$(COMPOSE) --profile dev up -d --build

watch: 
	$(COMPOSE) --profile dev up --build --watch

down:
	$(COMPOSE) down

down-v:
	$(COMPOSE) down -v

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
	$(COMPOSE) build --no-cache

clean:
	$(COMPOSE) down -v

certs-install:
	script/install-local-certs.sh

lint-frontend:
	cd src/frontend && npx prettier . --write

lint-api:
	cd src/api && go fmt

lint: lint-frontend lint-api
