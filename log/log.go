package log

import (
	"log"
	"os"
)

var Log = log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
