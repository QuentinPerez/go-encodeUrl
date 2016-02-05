package main

import (
	"os"

	"github.com/QuentinPerez/encodeUrl/encoding"
)

type Toto struct {
	UID string `url:"uid,ifStringIsNotEmpty"`
}

func main() {
	encurl.PrintAllFunctions(os.Stdout)
}
