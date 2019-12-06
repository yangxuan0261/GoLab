package test_http

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// 参考官方示例: https://github.com/buaazp/fasthttprouter/tree/master/examples/auth

// basicAuth returns the username and password provided in the request's
// Authorization header, if the request uses HTTP Basic Authentication.
// See RFC 2617, Section 2.
func basicAuth(ctx *fasthttp.RequestCtx) (username, password string, ok bool) {
	auth := ctx.Request.Header.Peek("Authorization")
	if auth == nil {
		return
	}
	return parseBasicAuth(string(auth))
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}

	enStr := auth[len(prefix):]
	fmt.Printf("--- enStr:%v\n", enStr)
	c, err := base64.StdEncoding.DecodeString(enStr)
	if err != nil {
		return
	}

	// 解码后根据自己的规则解出对应的信息
	cs := string(c)
	fmt.Printf("--- cs:%v\n", cs)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// BasicAuth is the basic auth handler
func BasicAuth(h fasthttp.RequestHandler, requiredUser, requiredPassword string) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := basicAuth(ctx)

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(ctx)
			return
		}
		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
		ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
	})
}

// Index is the index handler
func Index222(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Not protected!\n")
}

// Protected is the Protected handler
func Protected(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Protected!\n")
}

func Test_SrvFasthttpAuth(t *testing.T) {
	user := "gordon"
	pass := "secret!"

	router := fasthttprouter.New()
	router.GET("/", Index222)
	router.POST("/protected", BasicAuth(Protected, user, pass))

	log.Fatal(fasthttp.ListenAndServe(":8001", router.Handler))
}

func Test_parseBasicAuth(t *testing.T) {
	authStr := "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ=="
	user, password, hasAuth := parseBasicAuth(authStr)
	fmt.Println("--- result:", user, password, hasAuth) // Aladdin open sesame true
}
