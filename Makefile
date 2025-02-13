.PHONY: generate
generate:
	go generate ./...


MIGRATE=migrate
ENV_FILE=.env

include $(ENV_FILE)
DB_DSN=mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?charset=utf8mb4&parseTime=True&multiStatements=true

DB_DSN_TEST=mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)_test?charset=utf8mb4&parseTime=True&multiStatements=true

.PHONY: migrate-up migrate-down migrate-create

# 新しいマイグレーションファイルを作成
migrate-create:
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir migrations -seq $$name

# マイグレーションを適用
migrate-up:
	$(MIGRATE) -path migrations -database "$(DB_DSN)" up

migrate-up-test:
	$(MIGRATE) -path migrations -database "$(DB_DSN_TEST)" up

# マイグレーションをロールバック
migrate-down:
	$(MIGRATE) -path migrations -database "$(DB_DSN)" down

migrate-down-test:
	$(MIGRATE) -path migrations -database "$(DB_DSN_TEST)" down
