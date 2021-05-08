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

.PHONY:one
one:
	@cd demos/one && go run main.go

.PHONY:two
two:
	@cd demos/two && go run main.go

.PHONY:run
run:
	@cd demos/run && go run main.go

