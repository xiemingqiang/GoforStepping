package main

import (
	"context"
	"fmt"
	. "time"
)

func nodeT0(ctx context.Context){
	cancelCtx, _ := context.WithTimeout(ctx, 5 * Second)
	go nodeT1(cancelCtx)
	go nodeT2(cancelCtx)

	select {
	case <-cancelCtx.Done():
		fmt.Println("nodeT0 exit")
	}
}

func nodeT1(ctx context.Context) {
	go nodeT3(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("nodeT1 exit")
	case <-After(Second * 10):
		fmt.Println("nodeT1 finish")
	}
}

func nodeT2(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("nodeT2 exit")
	case <-After(Second * 10):
		fmt.Println("nodeT2 finish")
	}
}

func nodeT3(ctx context.Context) {
	select {
		case <-ctx.Done():
			fmt.Println("nodeT3 exit")
		case <-After(Second * 10):
			fmt.Println("nodeT3 finish")
	}
}

func main()  {
	firstCtx, _ := context.WithCancel(context.Background())
	go nodeT0(firstCtx)
	Sleep(10* Second)
}
