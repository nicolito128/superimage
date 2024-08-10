example:
	go run examples/$(name)/main.go

run-examples:
	go run examples/gopher/main.go
	go run examples/negative-gopher/main.go
	go run examples/blur-gopher/main.go
	go run examples/flip-gopher/main.go
	go run examples/opacity-gopher/main.go
	go run examples/reflect-gopher/main.go
