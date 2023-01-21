.PHONY: help
help: ## Prints out the options available in this makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: profile
profile: ## Run the solver and grab a CPU/memory profile using pprof
	go run main.go -profile -refresh=false -start=7 -finish=10 -numIterations=1
	pprof -web cpu.pprof
	pprof -web mem.pprof


.PHONY: compete
compete: ## Run the solver using the serial solver
	go run main.go

.PHONY: test
test: ## Run unit tests
	go test -timeout=30s -short ./...

.PHONY: bench
bench: ## Runs benchmarks TODO build these
	go test -benchmem -run="^$$" -bench "^BenchmarkAll$$" ./... -short -count=1

.PHONY: lint
lint: ## Runs linters (via golangci-lint) on golang code
	golangci-lint run -v ./...
