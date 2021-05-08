.PHONY:up
up:
	@docker compose up -d

.PHONY:down
down:
	@docker compose down

.PHONY:reboot
reboot:
	@docker compose down
	@docker compose up -d

.PHONY:run
run:
	@cd demos/one && go run main.go

