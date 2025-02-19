// spot bugs, performance, incidents
// log timestamp,debug, error, info
package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("/<os_path_to_store_file>/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	logrus.SetOutput(file)
	logrus.Println("Hello, World!")

	// logrus.SetLevel(logrus.DebugLevel)
	// logrus.Debug("Hello, World!")

	// logrus.SetLevel(logrus.TraceLevel)
	// logrus.Trace("Hello, World!")

}

/*
	--------------- 7 types ------------------
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
	Panic
*/
