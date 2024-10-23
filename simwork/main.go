package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var (
	multi = flag.Bool("multi", false, "Set to true to use all CPU/cores")
)

func api(w http.ResponseWriter, r *http.Request) {
	numCpu := 1
	if *multi {
		numCpu = runtime.NumCPU()
	}

	var ts int64 = 30
	hts := r.Header.Get("x-timeout-s")
	if hts != "" {
		v, err := strconv.ParseInt(hts, 10, 64)
		if err != nil {
			log.Println(err)
			return
		}

		ts = v
	}

	done := make(chan int)
	for i := 0; i < numCpu; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	go func() {
		for i := 0; i < int(ts); i++ {
			if !*multi {
				log.Println(i, "loading a single CPU...")
			} else {
				log.Println(i, "loading all CPUs...")
			}

			time.Sleep(time.Second * 1)
		}

		close(done)
	}()
}

func main() {
	flag.Parse()
	http.HandleFunc("/api", api)
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
