package main

import (
	"github.com/kataras/iris"
	"fmt"
	"os"
)

func Sum(a int, b int) int {
	return a+b
}

func Sub(a int, b int) int {
	return a-b;
}

func main() {
	app := iris.Default()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	app.Handle("GET", "/sum/{a:int}/{b:int}", func(ctx iris.Context) {
		a, errA := ctx.Params().GetInt("a") 
		b, errB := ctx.Params().GetInt("b")
		if errA != nil {
			panic(errA)
		} else if errA != nil {
			panic(errB)
		}
		html := fmt.Sprintf("<h1>%d + %d is %d</h1>", a, b, Sum(a, b))
		ctx.HTML(html)		
	})

	app.Handle("GET", "/sub/{a:int}/{b:int}", func(ctx iris.Context) {
		a, errA := ctx.Params().GetInt("a") 
		b, errB := ctx.Params().GetInt("b")
		if errA != nil {
			panic(errA)
		} else if errA != nil {
			panic(errB)
			
		}
		html := fmt.Sprintf("<h1>%d - %d is %d</h1>", a, b, Sub(a, b))
		ctx.HTML(html)		
	})

	app.Handle("GET", "/exp/{a:int}/{b:int}", func(ctx iris.Context) {
		a, errA := ctx.Params().GetInt("a") 
		b, errB := ctx.Params().GetInt("b")
		if errA != nil {
			panic(errA)
		} else if errA != nil {
			panic(errB)
		}
		ans := a
		for i := 0; i<b; i++ {
			ans *= a
		}
		html := fmt.Sprintf("<h1>%d + %d is %d</h1>", a, b, ans)
		ctx.HTML(html)		
	})

	app.Handle("GET", "/version/", func(ctx iris.Context) {
		vers := os.Getenv("GIT_TAG") 
		html := fmt.Sprintf("<h1>Website version: %s</h1>", vers)
		ctx.HTML(html)		
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Run(iris.Addr(":8080"))
}