package webrouter

import (
	router "github.com/buaazp/fasthttprouter"
	"strings"
	fast "github.com/valyala/fasthttp"
)

var (
	POST = "post"
	GET = "get"
	ROUTER = router.New()
	UrlCache = make(map[string]fast.RequestHandler)
)
// 路由
func RegisterRouterUrl(method string, path string, controller fast.RequestHandler) *router.Router {
	method = strings.ToLower(method)
	if strings.EqualFold(POST, method) {
		ROUTER.POST(path, controller)
	}
	if strings.EqualFold(GET, method) {
		ROUTER.GET(path, controller)
	}
	return ROUTER
}

func InitRouter() *router.Router {
	for key, _ := range UrlCache {
		ROUTER = RegisterRouterUrl(POST, key, UrlCache[key])
	}
	return ROUTER
}
