package error

import "log"

func Error(err error)  {
	if err != nil {
		log.Fatalf("we got an error", err)
		return
	}
}
