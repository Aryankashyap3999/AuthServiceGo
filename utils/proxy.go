package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseUrl string, pathPrefix string) http.HandlerFunc {
	target, err := url.Parse(targetBaseUrl)

	if err != nil {
		fmt.Println("Error parsing target URL:", err)
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)

		fmt.Println("Proxying request to:", targetBaseUrl+r.URL.Path)
		fmt.Println("Original request path:", r.URL.Path)
		fmt.Println("Path prefix to trim:", pathPrefix)

		r.URL.Path = strings.TrimPrefix(r.URL.Path, pathPrefix)

		fmt.Println("Modified request path:", r.URL.Path)

		r.Host = target.Host

		if userId, ok := r.Context().Value("userId").(string); ok {
			r.Header.Set("X-User-ID", userId)
		}
	}

	return proxy.ServeHTTP
}