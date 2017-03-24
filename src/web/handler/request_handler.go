package requesthandler

import (
	fast "github.com/valyala/fasthttp"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func Index(ctx *fast.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func Hello(ctx *fast.RequestCtx)  {
	Executed(ctx, helloController)
}
// 实际执行的controller中的逻辑
func helloController(param *simplejson.Json) (*simplejson.Json, error) {
	return param, nil
}
