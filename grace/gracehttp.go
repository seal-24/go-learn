package main

import (
"context"
"errors"
"flag"
"fmt"
"log"
"net"
"net/http"
"os"
"os/exec"
"os/signal"
"syscall"
"time"
)

var (
	server   *http.Server
	listener net.Listener
	graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
)

func handler(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()
	time.Sleep(20 * time.Second)
	end := time.Now()

	retMsg := fmt.Sprintf("request from %s ==> %s",
		begin.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	w.Write([]byte(retMsg))
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	server = &http.Server{Addr: ":8001"}

	var err error
	if *graceful {
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		listener, err = net.Listen("tcp", server.Addr)
	}

	if err != nil {
		log.Fatalf("listener error: %v", err)
	}

	go func() {
		err = server.Serve(listener)
		log.Printf("server.Serve err: %v\n", err)
	}()
	signalHandler()
	log.Printf("signal end")
}
/*
启动web
$ ps aux|grep web_server
root    24115  0.0  0.0 378092  3272 pts/13   Sl+  20:50   0:00 ./web_server


会话1
$ curl localhost:8001
与此同时， 会话2
$ killall -USR2 web_server; curl localhost:8001

因为第一个curl还没有处理完，父进程没有退出。
$ ps aux|grep web_server
root    24115  0.0  0.0 386544  3736 pts/13   Sl+  20:50   0:00 ./web_server
root    24441  0.0  0.0 443628  3476 pts/13   Sl+  20:51   0:00 ./web_server -graceful



20s之后，两个curl请求都可以正常收到响应。

第一个curl处理完成后，父进程退出
$ ps aux|grep web_server
root    24441  0.0  0.0 443884  3612 pts/13   Sl   20:51   0:00 ./web_server -graceful
*/
func reload() error {
	tl, ok := listener.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}

	args := []string{"-graceful"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{f}
	return cmd.Start()
}

func signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		log.Printf("signal: %v", sig)

		// timeout context for shutdown
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			// stop
			log.Printf("stop")
			signal.Stop(ch)
			server.Shutdown(ctx)
			log.Printf("graceful shutdown")
			return
		case syscall.SIGUSR2:
			// reload
			log.Printf("reload")
			err := reload()
			if err != nil {
				log.Fatalf("graceful restart error: %v", err)
			}
			server.Shutdown(ctx)
			log.Printf("graceful reload")
			return
		}
	}
}