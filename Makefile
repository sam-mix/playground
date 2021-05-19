.PHONY:up
up:
	@docker-compose up -d

.PHONY:down
down:
	@docker-compose down

.PHONY:reboot
reboot:
	@docker-compose down
	@docker-compose up -d

.PHONY:one
one:
	@cd demos/one && go run main.go

.PHONY:two
two:
	@cd demos/two && go run main.go

.PHONY:run
run:
	@cd demos/run && go run main.go


.PHONY:err
err:
	@cd demos/err && go run main.go


.PHONY:info
info:
	@cd demos/info && go run main.go

.PHONY:dao-log
dao-log:
	@cd demos/dao && go run main.go

.PHONY:where-true
where-true:
	@cd demos/err1 && go run main.go


.PHONY:hacking
hacking:
	@cd playground/hacking && go run main.go

.PHONY:new-insert
new-insert:
	@cd demos/new-insert && go run main.go

.PHONY:godebug
godebug:
	@cd playground/godebug && GODEBUG=gctrace=1 go run main.go

.PHONY:error-stack
error-stack:
	@cd playground/error-stack/caller && go run main.go




