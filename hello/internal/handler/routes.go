// Code generated by jzero. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"github.com/jzero-io/jzero-admin-plugins/hello/internal/svc"

	version "github.com/jzero-io/jzero-admin-plugins/hello/internal/handler/version"
)

var (
	_ = http.StatusOK
	_ = time.Now()
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/version",
					Handler: version.Get(serverCtx),
				},
			},
			rest.WithPrefix("/api/hello/v1"),
		)
	}

}
