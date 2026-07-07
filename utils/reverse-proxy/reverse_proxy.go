package reverseproxy

import (
	"fmt"
	"log"

	"net/http"
	"net/http/httputil"
	"net/url"
)

func CreateProxyMiddleware(targetURL string, pathPrefix string) http.Handler{
	target,err:=url.Parse(targetURL)

	if err != nil{
		fmt.Println("Error while parsing the url: ",err)
		return nil
	}

	 proxy := &httputil.ReverseProxy{
        Rewrite: func(req *httputil.ProxyRequest) {
            // Set the target host and scheme
            req.SetURL(target)

			req.Out.Host = target.Host

			log.Printf("%s -> %s%s",req.In.RemoteAddr,target.Host,req.Out.URL.Path)
            
        },
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w,err.Error(),http.StatusBadGateway)
		},
    }
	return proxy


}