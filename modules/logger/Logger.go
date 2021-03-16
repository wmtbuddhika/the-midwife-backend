package logger

import "log"

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

func AddLogger(prefix string, message string) {
	log.SetPrefix(prefix + ": ")
	if prefix == INFO {
		log.Println(message)
	} else if prefix == WARN {
		log.Println(message)
	} else if prefix == ERROR {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Fatal(message)
	}
}
