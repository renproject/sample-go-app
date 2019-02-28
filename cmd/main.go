package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello world!")
	id := uuid.New()
	fmt.Printf("id: %s\n", id.String())
}
