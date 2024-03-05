// lib/db/fx.go
package db

import "go.uber.org/fx"

var Module = fx.Module("db",
	fx.Provide(
		fx.Annotate(
			NewDatabase,
			fx.As(new(Database)),
		),
	),
)

//与配置模块一样，我们提供了NewDatabase函数。但这一次需要添加一个annotation。
//这个annotation告诉FX不应该将NewDatabase函数的结果公开为*GormDatabase，而应该公开为Database接口。这再次允许我们将使用与实现解耦，因此可以稍后替换Gorm，而不必更改其他地方的代码。
//不要忘记在main.go中注册db.Module。
