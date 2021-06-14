package main

import (
	"context"
	"fmt"
	. "time"
)

func nodeV0(ctx context.Context){
	valueCtx := context.WithValue(ctx, "req", "jjmatch")
	go nodeV1(valueCtx)

	Sleep(2 * Second)
	fmt.Println("node0 exit")
}

func nodeV1(ctx context.Context) {
	valueCtx := context.WithValue(ctx, "req1", "jjmatch01")
	go nodeV2(valueCtx)

}

func nodeV2(ctx context.Context) {
	fmt.Println(ctx.Value("req"))
	fmt.Println(ctx.Value("req1"))
}

func main()  {
	firstCtx, _ := context.WithCancel(context.Background())
	go nodeV0(firstCtx)
	Sleep(5* Second)
}
