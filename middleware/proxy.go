package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetURL string, pathPrefix string) http.HandlerFunc{
	target,err:=url.Parse(targetURL)

	if err != nil{
		fmt.Println("Error while parsing the url: ",err)
		return nil
	}

	 proxy := &httputil.ReverseProxy{
        Rewrite: func(req *httputil.ProxyRequest) {
            // Set the target host and scheme
            req.SetURL(target)

            // Strip the /api/v2 prefix from the path
            // "/api/v2/users" becomes "/users"
            if strings.HasPrefix(req.Out.URL.Path, "/api/v2") {
                req.Out.URL.Path = strings.TrimPrefix(req.Out.URL.Path, "/api/v2")
                if req.Out.URL.Path == "" {
                    req.Out.URL.Path = "/"
                }
            }

            log.Printf("Forwarding: %s -> %s%s",
                req.In.RemoteAddr, target.Host, req.Out.URL.Path)
        },
    }


}