package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func dbQuery(ctx context.Context, sql string)  {
	reqId := ctx.Value("req_id")
	fmt.Printf("req_id:%s, sql:%s", reqId, sql )
}

func slowDBQuery(ctx context.Context, cancel context.CancelFunc)  {
	go func() {
		select {
		case <-ctx.Done():
			println("slowDBQuery():timeout")
		case <-time.After(1 * time.Second):
			reqId := ctx.Value("req_id")
			fmt.Printf("req_id:%s\n", reqId)
			println("slowDBQuery():normal exit")
			cancel()
		}
	}()
}
func main()  {
	g := gin.New()

	g.GET("/app", func(context *gin.Context) {
		reqId := context.Request.URL.Query().Get("req_id")
		sql := "select * from apps where appName=SDK"
		context.Set("req_id", reqId)
		dbQuery(context, sql)
		context.JSON(200, gin.H{"reqId":reqId, "appName":"SDK"})
	})

	g.GET("/timeout", func(ctx *gin.Context) {
		ctx.Set("req_id", "test")
		timeoutCtx, timeoutCancel := context.WithTimeout(ctx,  2 * time.Second)
		slowDBQuery(timeoutCtx, timeoutCancel)
		select {
		case <-timeoutCtx.Done():
			timeoutCancel()
			println("slowDBQuery timeout")
		}

	})
	g.Run(":8080")
}
