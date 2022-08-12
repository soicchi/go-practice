package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuidObj, _ := uuid.NewUUID()
	fmt.Println(uuidObj.String())
}