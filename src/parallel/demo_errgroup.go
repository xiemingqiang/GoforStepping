package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func myHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, this is a server")
}

func main()  {
	srv := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", myHandle)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	group, _ := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return srv.ListenAndServe()
	})

	group.Go(func() error {
		sig := <-sigs
		fmt.Println("=======")
		fmt.Println(sig)
		return srv.Shutdown(context.Background())
	})

	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}

}
