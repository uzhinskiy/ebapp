package main

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	var i int
	var sleep int
	var service string

	tmp1, e1 := os.LookupEnv("SERVICE")
	if e1 {
		service = tmp1
	} else {
		service = "Unknown"
	}

	tmp2, e2 := os.LookupEnv("SLEEP")
	if e2 {
		sleep, _ = strconv.Atoi(tmp2)
	} else {
		sleep = 10
	}

	for {
		i++
		log.Info().Str("service", service).Int("iterator", i).Msg("Hello")
		time.Sleep(time.Millisecond * time.Duration(500*sleep))
	}

}
