package logger

import (
	"log"
	"os"
	"time"

	"github.com/lowply/webhook/config"
)

func Log(t string) {
	c := config.GetConfig()
	logfile := c.Logfile
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	datetime := time.Now().Format(layout)
	output := datetime + ": " + t + "\n"

	_, err := os.Stat(logfile)
	if err != nil {
		log.Println("Created " + logfile)
		os.Create(logfile)
	}

	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(output)
	if err != nil {
		log.Fatal(err)
	}
}
