package main

import (
	"fmt"

	"github.com/QuentinPerez/go-encodeUrl"
	"github.com/Sirupsen/logrus"
)

type ID struct {
	Name        string `url:"name,ifStringIsNotEmpty"`
	DisplayName string `url:"display-name,ifStringIsNotEmpty"`
}

func main() {
	values, errs := encurl.Translate(&ID{"qperez", "Quentin Perez"})
	if errs != nil {
		logrus.Fatal(errs)
	}
	fmt.Printf("https://example.com/?%v\n", values.Encode())
}
