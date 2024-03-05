// main.go
package main

import (
	"fxdemo/config"
	"fxdemo/db"
	"fxdemo/http"
	"go.uber.org/fx"
	stdhttp "net/http"
)

func main() {
	app := fx.New(
		config.Module,
		db.Module,
		//因为FX只在需要时调用提供程序，我们没使用刚才构建的配置，所以FX不会加载。
		//我们添加了对fix.Invoke的调用，注册在应用程序一开始就调用的函数，这将是程序的入口，稍后将启动我们的HTTP服务器。
		//fx.Invoke(func(cfg *config.Config) {}),
		http.Module,
		fx.Invoke(func(cfg *config.Config, handler stdhttp.Handler) error {
			go stdhttp.ListenAndServe(cfg.HTTP.ListenAddress, handler)
			return nil
		}),
	)
	app.Run()

}

/*
理解:
在你的 main 包中，你使用了 Uber 的 fx 框架来组织你的应用程序。在 main 包中，你创建了一个 fx.App 实例，并传入了一些模块，其中包括 http.Module。

在 http 包中，你定义了一个模块 Module，它通过 fx.Provide 函数提供了一个名为 NewServer 的函数。在 NewServer 函数中，你创建了一个 Server 类型的实例，并将其返回。此外，你通过 fx.Annotate 函数为 NewServer 函数添加了一个注释，表示该函数返回一个 http.Handler 类型的实例。

当你在 main 包中调用 fx.New 函数创建 fx.App 实例时，你将 http.Module 传入其中。fx 框架会检测到这个模块，并执行其中的 fx.Provide 函数提供的函数。在这种情况下，它会调用 NewServer 函数，并通过 fx.As 注释将返回的实例标记为 http.Handler 类型。

因此，当你在 main 包中使用 http.Module 时，fx 框架会自动将 NewServer 函数返回的 Server 实例转换为 http.Handler 类型的实例，并将其传递给依赖于 http.Handler 类型的地方，比如你在 main 包中的 fx.Invoke 函数中。
*/
