package errorx

import "log"

func Must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
