package proxy

import (
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)


func RevProxy(conf *struct {
	Destination string `json:"destination"`
	SrcHost     string `json:"src_host"`
	SrcPort     string `json:"src_port"`
	SrcProto    string `json:"src_protocol"`
	StripPrefix string `json:"strip_prefix"`
}) http.Handler {
	proxy := &httputil.ReverseProxy{Director: func(request *http.Request) {
		origin := conf.SrcHost
		request.Header.Add("X-Forwarded-Host", request.Host)
		request.Header.Add("X-Origin-Host", origin)
		request.Host = origin
		request.URL.Host = origin
		request.URL.Scheme = conf.SrcProto
	},
	Transport: &http.Transport{
		Dial: (&net.Dialer{Timeout: 10 * time.Second}).Dial,
	}}
	return proxy
}