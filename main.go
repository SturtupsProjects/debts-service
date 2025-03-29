package main

import (
	"log"
)

func main() {
	a := map[string]string{"a": "b", "c": "d"}

	data, ok := a["d"]

	log.Println(ok)
	log.Println(data)
}
