// lib/http/fx.go
package http

import (
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Module("http", fx.Provide(
	fx.Annotate(
		NewServer,
		fx.As(new(http.Handler)),
	),
))

//将服务器公开为http.Handler，这样就可以用更高级的工具(如Gin或Gorilla Mux)替换刚才构建的简单HTTP服务器。
//现在，我们可以将模块导入到main函数中，并编写一个Invoke调用来启动服务器。
