// lib/config/fx.go
package config

import "go.uber.org/fx"

var Module = fx.Module("config", fx.Provide(NewConfig))

//这里通过使用fx.Module发布FX模块，这个函数接受两种类型的参数:
//
//第一个参数是用于日志记录的模块的名称。
//其余参数是希望向应用程序公开的依赖项。
//这里我们只使用fx.Provide导出Config对象，这个函数告诉FX使用NewConfig函数来加载配置。
//
//值得注意的是，如果Viper加载配置失败，NewConfig也会返回错误。如果错误不是nil, FX将显示错误并退出。
//
//第二个要点是，该模块不导出Viper，而只导出配置实例，从而允许我们轻松的用任何其他配置框架替换Viper。
