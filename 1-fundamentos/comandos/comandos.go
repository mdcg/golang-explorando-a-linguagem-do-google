package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!\n", "Go")
}

/*
$ go
$ go help <comando>
$ go version
$ godoc -http=:6060 // Documentação offline
$ go env
$ go doc cmd/vet // Verificar codigo morto.
$ go vet <arquivo>.go
$ go build <arquivo>.go // Cria um executável
$ go run <arquivo>.go // Compila e executa
*/
