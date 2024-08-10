example:
	go run examples/$(name)/main.go

run-examples:
	go run examples/gopher/main.go
	go run examples/negative/main.go
	go run examples/blur/main.go
	go run examples/flip/main.go
	go run examples/opacity/main.go
	go run examples/reflect/main.go
