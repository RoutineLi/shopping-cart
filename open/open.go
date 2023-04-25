package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"graduate_design/user/rpc/userclient"
	"io"
	"net/http"

	"graduate_design/open/internal/config"
	"graduate_design/open/internal/handler"
	"graduate_design/open/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/open-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			data, _ := io.ReadAll(r.Body)
			_, err := ctx.RpcUser.OpenAuth(context.Background(), &userclient.OpenAuthRequest{Body: data})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}
			r.Body = io.NopCloser(bytes.NewReader(data))
			next(w, r)
		}
	})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
