package test_http

import (
	syserr "GoLab/common/error"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"testing"
)

/*
http 测试最后在命令行工具中执行 $ go test -v -run ^(Test_SrvFasthttp01)$
在 goland 中可能 fmt.Printf 不能输出日志
*/

// 参考官方示例: https://github.com/buaazp/fasthttprouter/tree/master/examples/basic

// index 页
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Printf("--- Index")
	fmt.Fprint(ctx, "Welcome")
}

// 简单路由页
func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Printf("--- Hello")
	fmt.Fprintf(ctx, "hello")

	defer syserr.Recover()
	panic("wolegequ") // 再请求中 defer 才有效, 每一个请求都是一个 gor, 只能在当前 gor 中 recover()
}

// 获取GET请求json数据
// 使用 ctx.QueryArgs() 方法
func GetTest(ctx *fasthttp.RequestCtx) {
	fmt.Printf("--- GetTest")
	values := ctx.QueryArgs()
	fmt.Fprint(ctx, "--- get ret abc:"+string(values.Peek("abc")))
	// http://localhost:8001/get?abc=123
	//--- get ret abc:123
}

// 获取 url 中 占位符 的值
func MultiParams(ctx *fasthttp.RequestCtx) {
	fmt.Printf("--- MultiParams")
	fmt.Fprintf(ctx, "hi, %s, %s!\n", ctx.UserValue("name"), ctx.UserValue("word"))
	// http://localhost:8001/multi/aaa/bbb
	// hi, aaa, bbb!

	// http://localhost:8001/multi/aaa 则获取不到
	// Not Found
}

// 获取post的请求json数据
func PostTest(ctx *fasthttp.RequestCtx) {
	fmt.Printf("--- PostTest\n")

	postValues := ctx.PostArgs() // 貌似木有卵用
	fmt.Printf("--- postValues:%+v\n", string(postValues.Peek("bbb")))
	/*
		// PostArgs returns POST arguments.
		// It doesn't return query arguments from RequestURI - use QueryArgs for this.
		// Returned arguments are valid until returning from RequestHandler.
	*/

	formValues := ctx.FormValue("aaa") // ?aaa=111&bbb=222 请求参数 (表单数据)
	fmt.Printf("--- formValues aaa:%+v\n", string(formValues))

	ck := ctx.Request.Header.Peek("ccc") // 获取 token 之类的数据, 等价于官方 http 的 req.Header.Get("ccc")
	fmt.Printf("--- Cookie ccc:%+v\n", string(ck))

	// 这两行可以获取PostBody数据，在上传数据文件的时候有用
	postBody := ctx.PostBody()
	fmt.Fprint(ctx, "--- post ret abc:"+string(postBody))
}

// 测试 设置返回码跟返回信息
func Post403(ctx *fasthttp.RequestCtx) {
	fmt.Printf("--- Post403\n")

	if false { // 复杂 接口
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.SetBodyString(fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed))
		return
	}

	if true { // 简洁 接口
		ctx.Error("--- Forbidden", fasthttp.StatusForbidden)
		return
	}
}

func Test_SrvFasthttp01(t *testing.T) {

	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/hello", Hello)
	router.GET("/get", GetTest)
	router.GET("/multi/:name/:word", MultiParams)
	router.POST("/post", PostTest)
	router.POST("/test403", Post403)

	log.Fatal(fasthttp.ListenAndServe(":8001", router.Handler))
}

// -----------------
// https://juejin.im/post/5c3dc85f51882524f2302ce6

func Test_GetFasthttp01(t *testing.T) {
	url := `http://httpbin.org/get`

	status, rsp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	fmt.Println("--- rsp", string(rsp))
}

func Test_PostFasthttp01(t *testing.T) {
	url := `http://httpbin.org/post?key=123`

	// 填充表单，类似于net/url
	args := &fasthttp.Args{}
	args.Add("name", "test")
	args.Add("age", "18")

	status, rsp, err := fasthttp.Post(nil, url, args)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	fmt.Println("--- rsp", string(rsp))
}

func Test_PostFasthttp02(t *testing.T) {
	url := `http://httpbin.org/post?key=123`

	req := fasthttp.AcquireRequest()
	rsp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(rsp)
		fasthttp.ReleaseRequest(req)
	}()

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)

	requestBody := []byte(`{"request":"test"}`)
	req.SetBody(requestBody)

	if err := fasthttp.Do(req, rsp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	b := rsp.Body()
	fmt.Println("--- rsp", string(b))
}