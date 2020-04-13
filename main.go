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
//API endpoint:   https://api.run.pivotal.io (API version: 2.146.0)
//User:           986244259@qq.com
//Org:            e.g.9862442592
//Space:          sin
//只要业务逻辑通了 则什么都可以进行处理 只需要这样的方式去处理
// 我们只关心对应的代码的质量 代码的规范 代码的路口地址 代码的微服务的架构
// 可以扩展可以高可用 接口模式 适配器模式 代理模式 初始化模式 注入模式 控制翻转模式 反射模式
// 网络缓存模式 命令处理器 客户端client 接受处理器 处理 校验 认证 Token Jwt 等等 然后 重新
// 接受 json-body json-cmd-body 处理 cmdStuct
// EventComon->推送到对应的 事件仓库
//Rabimq
//处理器事件处理器
//分发管理处理
