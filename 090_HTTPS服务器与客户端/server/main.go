/*
	HTTPS服务器：

	1.生成RSA私钥
	openssl genrsa -out server.key 2048

	2.生成证书签名请求(CSR)
	openssl req -new -key server.key -out server.csr

	3.生成自签名证书(使用自己的私钥给自己颁发)
	openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
*/

package main

import (
	"fmt"
	"net/http"
)

/*
	HTTPS服务器（单向认证：客户端认证服务器）
*/

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, https!"))
	})

	//使用自签名证书作为该服务器的根证书
	if err := http.ListenAndServeTLS("127.0.0.1:8085", "server.crt", "server.key", nil); err != nil {
		fmt.Println(err)
		return
	}
}

/*
	HTTPS双向认证（服务器认证客户端,客户端也得认证服务器）
*/

// type handler struct {
// }

// func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello, https!"))
// }

// func main() {

// 	//注册客户端的自签名证书
// 	caCert, err := ioutil.ReadFile("../client/client.crt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//添加到CA池
// 	caCerPool := x509.NewCertPool()
// 	caCerPool.AppendCertsFromPEM(caCert)

// 	//创建TLS配置
// 	config := tls.Config{
// 		//需要客户端的证书并且验证
// 		ClientAuth: tls.RequireAndVerifyClientCert,
// 		ClientCAs:  caCerPool,
// 	}

// 	//使用TLS配置创建服务器
// 	server := http.Server{
// 		Addr:      ":8085",
// 		Handler:   &handler{},
// 		TLSConfig: &config,
// 	}

// 	//启动服务器
// 	err = server.ListenAndServeTLS("server.crt", "server.key")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
