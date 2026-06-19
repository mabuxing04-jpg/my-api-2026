package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	target, _ := url.Parse("https://oneapi.html.zone")
	proxy := httputil.NewSingleHostReverseProxy(target)
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = target.Host
	proxy.ServeHTTP(w, r)
}
