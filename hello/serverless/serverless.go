package serverless

import (
	"path/filepath"

	"github.com/jzero-io/jzero-contrib/dynamic_conf"
	"github.com/zeromicro/go-zero/core/conf"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin-plugins/hello/internal/config"
	"github.com/jzero-io/jzero-admin-plugins/hello/internal/handler"
	"github.com/jzero-io/jzero-admin-plugins/hello/internal/svc"
)

type Serverless struct {
	SvcCtx      *svc.ServiceContext                                   // 服务上下文
	HandlerFunc func(server *rest.Server, svcCtx *svc.ServiceContext) // 服务路由
}

// Serverless please replace coreSvcCtx any type to real core svcCtx
func New(coreSvcCtx any) *Serverless {
	ss, err := dynamic_conf.NewFsNotify(filepath.Join("plugins", "hello", "etc", "etc.yaml"), dynamic_conf.WithUseEnv(true))
	logx.Must(err)
	cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
		Type: "yaml",
	}, ss)
	c, err := cc.GetConfig()
	logx.Must(err)

	if err := conf.Load(filepath.Join("plugins", "hello", "etc", "etc.yaml"), &c); err != nil {
		panic(err)
	}

	svcCtx := svc.NewServiceContext(cc)
	return &Serverless{
		SvcCtx:      svcCtx,
		HandlerFunc: handler.RegisterHandlers,
	}
}