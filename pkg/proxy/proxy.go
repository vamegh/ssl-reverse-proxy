package proxy

import (
	"github.com/vamegh/ssl-reverse-proxy/pkg/configHandler"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

func RevProxy(conf configHandler.ProxyDetails) http.Handler {
	proxy := &httputil.ReverseProxy{Director: func(request *http.Request) {
		var origin = conf.SrcHost
		request.Header.Add("X-Forwarded-Host", request.Host)
		request.Header.Add("X-Origin-Host", origin)
		request.Host = origin
		request.URL.Host = origin
		request.URL.Scheme = conf.SrcProto
	},
	Transport: &http.Transport{
		Dial: (&net.Dialer{Timeout: time.Duration(conf.IdleTimeout) * time.Second}).Dial,
	}}
	return proxy
}