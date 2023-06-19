package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"graduate_design/user/rpc/userclient"
	"graduate_design/websocket/consume"
	"graduate_design/websocket/internal/config"
	"graduate_design/websocket/internal/handler"
	"graduate_design/websocket/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/websocket-api.yaml", "the config file")

func main() {
	flag.Parse()
	logx.DisableStat()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			webToken := r.URL.Query().Get("token")
			if webToken != "" {
				auth, err := ctx.RpcUser.Auth(context.Background(), &userclient.UserAuthRequest{Token: webToken})
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
					return
				}
				ctx.AuthResp = auth
				next(w, r)
				return
			}

			if r.Header.Get("token") == "" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}
			auth, err := ctx.RpcUser.Auth(context.Background(), &userclient.UserAuthRequest{Token: r.Header.Get("token")})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}
			ctx.AuthResp = auth
			next(w, r)
		}
	})

	handler.RegisterHandlers(server, ctx)

	threading.GoSafe(func() {
		consume_mqtt.NewConsumer(ctx)
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
