package main

import (
	"fmt"
	"log"
	"net/url"
)

var authCallback string

func main() {
	authCallback = "wercker-auth.cfapps.io/callback"
	//	authCallback = url.QueryEscape(authCallback)
	l, err := url.Parse(authCallback)
	if err != nil {
		log.Println("url encoding exception : ", err)
	}
	encodeurl := l.Scheme + "://" + l.Host + "?" + l.Query().Encode()
	//打印一下
	fmt.Println(encodeurl)
}
