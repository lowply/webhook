package logger

import (
	"log"
	"os"
	"time"
)

func Log(filename, t string) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	datetime := time.Now().Format(layout)
	output := datetime + ": " + t + "\n"

	_, err := os.Stat(filename)
	if err != nil {
		log.Println("Created " + filename)
		os.Create(filename)
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.WriteString(output); err != nil {
		log.Fatal(err)
	}
}
