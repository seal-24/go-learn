package main

import (
 "net/http"
 "log"
	"os"
	"os/signal"
	"syscall"
	"time"
 "context"
)

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request time:", time.Now().String())
		time.Sleep(time.Second * 10)
		log.Println("response time:", time.Now().String())
		w.Write([]byte("Hello World!"))
	})

	return mux
}

func main() {
	server := &http.Server{
		Addr:    ":8001",
		Handler: newHandler(),
	}
	//non grace stop
	//go handleSig(server)
	//err := server.ListenAndServe()

	//grace stop
	go server.ListenAndServe()
	handleSig(server)
	//log.Printf("[ListenAndServe] done err=%#v", err)
}

func handleSig(s *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	s.Shutdown(context.Background())
	log.Println("[Shutdown] done." + time.Now().Format("2006-01-02 15:04:05.999999"))
}

