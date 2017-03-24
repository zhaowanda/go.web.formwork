package requesthandler

import (
	"errors"
	"github.com/bitly/go-simplejson"
	fast "github.com/valyala/fasthttp"
	"web/convert"
	"web/interceptor"
)

// 请求之前执行拦截
func before(ctx *fast.RequestCtx) (bool, error) {
	entity, ok := interceptor.GetInterceptor("authInterceptor")
	if !ok {
		return false, errors.New("请检查系统配置")
	}
	entity = entity.AddExcludePattern("authInterceptor", "/hello")
	flag, err := entity.ExecutedInterceptor(ctx, entity)
	return flag, err
}

// 请求之后执行转换
func after(ctx *fast.RequestCtx, args interface{}, errMsg error) {
	convertResult := convert.HttpMessageConvertFunc(convert.JsonHttpMessageConvert)
	convertResult.MessageConvert(ctx, args, errMsg)
}

// base controller interface
type BaseControllerInterface interface {
	BaseController(ctx *fast.RequestCtx, execute ExecuteFunc)
}

// 注册执行实际业务逻辑的controller
type BaseControllerFunc func(ctx *fast.RequestCtx, execute ExecuteFunc)

// base controller interface impl
func (bcf BaseControllerFunc) BaseController(ctx *fast.RequestCtx, execute ExecuteFunc) {
	bcf(ctx, execute)
}

// 执行controller中的业务逻辑
func baseController(ctx *fast.RequestCtx, execute ExecuteFunc) {
	flag, err := before(ctx)
	if flag {
		// todo 执行业务逻辑
		//execute()
		param, err := convert.ArgumentResolver(ctx)
		if err != nil {
			after(ctx, nil, err)
			return
		}
		result, error := execute(param)
		after(ctx, result, error)
		return
	}
	after(ctx, nil, err)
}

type ExecuteFunc func(param *simplejson.Json) (*simplejson.Json, error)

// 所有的controller层都调用此方法
func Executed(ctx *fast.RequestCtx, execute ExecuteFunc) {
	controller := BaseControllerFunc(baseController)
	controller.BaseController(ctx, execute)
}
