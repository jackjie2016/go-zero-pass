package main

import "fmt"

type UserIsEmpty struct {
}

type UserHasField struct {
	Age uint64 `json:"age"`
}

func FPrint(uIsEmpty UserIsEmpty, uHasField UserHasField) {
	fmt.Printf("FPrint uIsEmpty:%p uHasField:%p\n", &uIsEmpty, &uHasField)
}

func main() {
	uIsEmpty := UserIsEmpty{}
	uHasField := UserHasField{
		Age: 10,
	}
	FPrint(uIsEmpty, uHasField)
	fmt.Printf("main: uIsEmpty:%p uHasField:%p\n", &uIsEmpty, &uHasField)
}
