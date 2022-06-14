check_defined = \
	$(strip $(foreach 1,$1, \
		$(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
	$(if $(value $1),, \
	$(error Undefined $1$(if $2, ($2))))

all: deps swagger build run

swagger:
	@echo "Generating Swagger documentation..."
	# General Swagger documentation
	swag init \
		--dir ./,./docs/response \
		--exclude ./endpoints/admin \
		-g ./main.go

build:
	$(call check_defined, env, Environment of API)
	go build -ldflags "-X main.environment=$(env)"

run:
	./golang-api

deps:
	@echo "[Install Deps] running phase"
	go get -u github.com/swaggo/swag/cmd/swag@v1.8.0