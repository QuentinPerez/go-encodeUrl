package main

import (
	"fmt"

	"github.com/QuentinPerez/encodeUrl/encoding"
	"github.com/Sirupsen/logrus"
)

type Toto struct {
	UID string `url:"uid,ifStringIsNotEmpty"`
}

func main() {
	values, errs := encurl.Translate(&Toto{"qperez"})
	if errs != nil {
		logrus.Fatal(errs)
	}
	fmt.Println(values)
}
