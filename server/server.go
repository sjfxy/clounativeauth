package server

import (
	"net/http"

	"github.com/astaxie/beego/session"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/cf-tools"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type authConfig struct {
	ClientID     string
	ClientSecret string
	Domain       string
	CallbackURL  string
}

//NewServer configures and returns a Negroni server
func NewServer(appEnv *cfenv.App) *negroni.Negroni {
	// HACK handle these failures for realzies
	authClientID, _ := cftools.GetVCAPServiceProperty("authzero", "id", appEnv)
	authSecret, _ := cftools.GetVCAPServiceProperty("authzero", "secret", appEnv)
	authDomain, _ := cftools.GetVCAPServiceProperty("authzero", "domain", appEnv)
	authCallback, _ := cftools.GetVCAPServiceProperty("authzero", "callback", appEnv)
	if len(authClientID) == 0 {
		authClientID = "EcFMxUkuZh3Onf55c0xNXndOP51K6x7V"
	}
	if len(authSecret) == 0 {
		authSecret = "UdQPI1zOVoshFLB-qww1ffXZnnHS4eY2psrBb9chZzUwzDI-YGPRVryIB-shpbPx"
	}
	if len(authDomain) == 0 {
		authDomain = "dev-y9rmal81.auth0.com"
	}
	if len(authCallback) == 0 {
		//这里设置一下我们的回调的数据格式
		//	authCallback = "wercker-auth.cfapps.io" //设置回调的是 3006作为代理端口进行提供回调处理
		authCallback = "wercker-authv5.cfapps.io/callback"
		//	authCallback = url.QueryEscape(authCallback)
		//	fmt.Println(authCallback)
		//	log.Println(authCallback)
		//	authCallback = "wercker-auth/callback"
	}
	config := &authConfig{
		ClientID:     authClientID,
		ClientSecret: authSecret,
		Domain:       authDomain,
		CallbackURL:  authCallback,
	}
	//cf := &session.ManagerConfig{
	//	CookieName: "gosessionid",
	//	Gclifetime: 3600,
	//}

	//TODO: real-world app needs out-of-process session management (e.g. backed by Redis)
	sessionManager, _ := session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go sessionManager.GC()

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, sessionManager, config)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, sessionManager *session.Manager, config *authConfig) {
	mx.HandleFunc("/", homeHandler(config))
	mx.HandleFunc("/callback", callbackHandler(sessionManager, config))
	mx.Handle("/user", negroni.New(
		negroni.HandlerFunc(isAuthenticated(sessionManager)),
		negroni.Wrap(http.HandlerFunc(userHandler(sessionManager))),
	))
	mx.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
}
