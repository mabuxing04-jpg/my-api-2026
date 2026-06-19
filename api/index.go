package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 将所有请求无缝转接到全球最稳定的官方免编译后端核心上
	target, _ := url.Parse("https://oneapi.html.zone")
	proxy := httputil.NewSingleHostReverseProxy(target)
	
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = target.Host
	proxy.ServeHTTP(w, r)
}
