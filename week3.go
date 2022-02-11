package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func serveApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello world")
	})
	return errors.New("handleFunc error")
}

func serveDebug(stop <-chan struct{}) error {
	return serve("127.0.0.1:8081", http.DefaultServeMux, stop)
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func Signal(sin os.Signal, stop <-chan struct{}) error {
	done := make(chan os.Signal)
	defer close(done)

	signal.Notify(done, sin)
	go func() {
		<-stop
		signal.Stop(done)
	}()

	return errors.New("signal err")
}

func main() {
	done := make(chan error, 4)
	stop := make(chan struct{})

	var sign os.Signal
	group := new(errgroup.Group)
	group.Go(func() error {
		done <- serveDebug(stop)
		return nil
	})
	group.Go(func() error {
		done <- serveApp(stop)
		return nil
	})
	group.Go(func() error {
		done <- Signal(sign, stop)
		return nil
	})

	err := group.Wait()
	if err != nil {
		done <- err
	}
	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("err:", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
