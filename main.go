package main

import (

	"github.com/Sirupsen/logrus",

	"github.com/amitm1/go-service/cmd"
)

func main() {
	//fmt.Printf("Hello, world.\n")
	fmt.Printf(cmd.Reverse("Hello, world Amit Mishra.\n"))

	logrus.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

}
