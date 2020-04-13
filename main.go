package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/sjfxy/clounativeauth/server"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3006"
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Printf("FATAL: Could not retrieve CF environment: %v\n. This app needs auth0 configuration injected via VCAP_SERVICES.\n", err)
		//	os.Exit(1)
	}
	s := server.NewServer(appEnv)
	s.Run(":" + port)
}

//整个的流程是这样的处理方式
//1.服务单入口的服务启动server_run_main.php
//2.然后server/目录下面的server.go进行启动
//3.进行就是控制器路由处理方式 /home是定位到登录的跳转页面 准备相应的参数 cofig对象 type Config struct
// 4.放在 types.go 文件中定义结构体和相应的方法和接口约束 用来做适配器模式 代理模式 接口模式
// 5. const.go 变量的配置文件 配置文件可以读取 在服务启动的使用 读取配置中心分发的配置
// 6. /callback我们设置的是对应的端口的30000端口
