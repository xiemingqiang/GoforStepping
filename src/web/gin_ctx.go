package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type NewContext struct {
	Gin *gin.Context
	context.Context
	context.CancelFunc
}

type NewContextHandle func(ctx *NewContext)

func newSlowDBQuery(ctx *NewContext)  {
	go func() {
		select {
		case <-ctx.Context.Done():
			println("slowDBQuery():timeout")
		case <-time.After(1 * time.Second):
			reqId := ctx.Gin.Value("req_id")
			fmt.Printf("req_id:%s\n", reqId)
			println("slowDBQuery():normal exit")
			//可以使用Gin.context的方法
			fmt.Printf("clientIp:%s\n", ctx.Gin.ClientIP())
			ctx.CancelFunc()
		}
	}()
}

func WithNewContext(handle NewContextHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeoutCtx, timeoutCancel := context.WithTimeout(c, 2 * time.Second)
		newCtx := NewContext{Gin:c, Context:timeoutCtx, CancelFunc:timeoutCancel}
		handle(&newCtx)
	}
}

func main()  {
	g := gin.New()

	g.GET("/timeout", WithNewContext(func(ctx *NewContext) {
		ctx.Gin.Set("req_id", "test")
		newSlowDBQuery(ctx)
		select {
		case <-ctx.Context.Done():
			ctx.CancelFunc()
			println("slowDBQuery timeout")
		}
	}))

	g.Run(":8080")
}
