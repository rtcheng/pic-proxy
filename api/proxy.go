package api

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Proxy(c *gin.Context) {
	s := c.Param("url")
	if s == "" {
		c.String(400, "Bad Request")
		return
	}
	u, err := url.Parse(s[1:])
	if err != nil {
		c.String(400, "Bad Request")
		return
	}
	c.Request.URL.Path = u.Path
	c.Request.Host = u.Host
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: u.Scheme, Host: u.Host})
	proxy.ServeHTTP(c.Writer, c.Request)
	c.Abort()
}
