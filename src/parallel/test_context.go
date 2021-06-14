package main

import (
	"context"
	"fmt"
	. "time"
)

func node0(ctx context.Context){
	cancelCtx, cancel := context.WithCancel(ctx)
	go node1(cancelCtx)
	go node2(cancelCtx)

	Sleep(2 * Second)
	fmt.Println("node0 exit")
	cancel()
}

func node1(ctx context.Context) {
	cancelCtx, _ := context.WithCancel(ctx)
	go node3(cancelCtx)
	select {
	case <-cancelCtx.Done():
		fmt.Println("node1 exit")
	case <-After(Second * 5):
		fmt.Println("node1 finish")
	}
}

func node2(ctx context.Context) {
	cancelCtx, _ := context.WithCancel(ctx)
	select {
	case <-cancelCtx.Done():
		fmt.Println("node2 exit")
	case <-After(Second * 5):
		fmt.Println("node2 finish")
	}
}

func node3(ctx context.Context) {
	cancelCtx, _ := context.WithCancel(ctx)
	select {
		case <-cancelCtx.Done():
			fmt.Println("node3 exit")
		case <-After(Second * 5):
			fmt.Println("node3 finish")
	}
}

func main()  {
	firstCtx, _ := context.WithCancel(context.Background())
	go node0(firstCtx)
	Sleep(10* Second)
}
