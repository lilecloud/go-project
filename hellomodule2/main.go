package main

import (
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

func main() {
	logrus.Println("hello module")
	logrus.Println(uuid.NewString())
}
