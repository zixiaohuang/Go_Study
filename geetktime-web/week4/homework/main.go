package homework

import "flag"

var (
	flagconf string
)

func init() {
	// 初始化配置
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	//主程序

	app, cleanup, err := initApp(bc.Server, &rc, bc.Data, logger, tp)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err:=app.Run(); err != nil {
		panic(err)
	}
}