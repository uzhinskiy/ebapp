package main

import (
	"os"
	"strconv"
	"time"
	"math/rand"
	"strings"
	"github.com/rs/zerolog/log"
)

// define the given charset, char only
var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var number = []byte("0123456789")
var alphaNumeric = append(charset, number...)

// n is the length of random string we want to generate
func randStr(n int) string {
    b := make([]byte, n)
    for i := range b {
	// randomly select 1 character from given charset
	b[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
    }
    return string(b)
}

func main() {
	var i int
	var sleep int
	var service string

	rand.Seed(time.Now().UnixNano())

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
        msg := strings.Repeat(randStr(1024), 10)
	for {
		i++
		log.Info().Str("service", service).Int("iterator", i).Msg(msg)
		time.Sleep(time.Millisecond * time.Duration(500*sleep))
	}

}
