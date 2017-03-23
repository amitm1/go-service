package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/amitm1/go-service/cmd"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf(cmd.Reverse("Hello, world Amit Mishra.\n"))

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

}
