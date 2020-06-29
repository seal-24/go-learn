package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}

func main() {
	go func() {
		for {
			log.Println(Add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}

//go tool pprof http://127.0.0.1:6060/debug/pprof/profile\?seconds\=60
//go tool pprof -http=:8080 /Users/sealwang/pprof/pprof.samples.cpu.001.pb.gz

//go tool pprof http://127.0.0.1:6060/debug/pprof/goroutine
//go tool pprof -http=:8080 /Users/sealwang/pprof/pprof.goroutine.001.pb.gz
