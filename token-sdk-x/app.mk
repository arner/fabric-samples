
# Setup application
.PHONY: setup-app
setup-app:
	./scripts/gen_crypto.sh
	docker-compose build

# Start application
.PHONY: start-app
start-app:
	PLATFORM=$(PLATFORM) docker-compose up -d

# Restart application
.PHONY: restart-app
restart-app:
	PLATFORM=$(PLATFORM) docker-compose restart

# Stop application
.PHONY: stop-app
stop-app:
	docker-compose stop

# Teardown application
.PHONY: teardown-app
teardown-app:
	docker-compose down
	rm -rf "$(CONF_ROOT)"/*/data

# Clean just the databases.
.PHONY: clean-data
clean-data:
	rm -rf "$(CONF_ROOT)"/*/data

# Clean everything and remove all the keys
.PHONY: clean-app
clean-app:
	rm -rf "$(CONF_ROOT)"/*/keys "$(CONF_ROOT)"/*/data "$(CONF_ROOT)"/namespace/*.json
