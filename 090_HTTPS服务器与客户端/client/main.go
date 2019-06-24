/*
	HTTPS客户端：

	1.生成RSA私钥
	openssl genrsa -out client.key 2048

	2.生成证书签名请求(CSR)
	openssl req -new -key client.key -out client.csr

	3.生成自签名证书(使用自己的私钥给自己颁发)
	openssl x509 -req -days 365 -in client.csr -signkey client.key -out client.crt
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	HTTPS单向认证（客户端认证服务器）
*/

func main() {

	//请求服务器时服务器会返回自己的证书，但客户端无法认证该证书，因此需要添加可接受的证书池

	//读取服务器的根证书
	caCert, err := ioutil.ReadFile("../server/server.crt")
	if err != nil {
		log.Fatal(err)
	}

	//创建CA池
	certPool := x509.NewCertPool()

	//将证书添加到CA池
	ok := certPool.AppendCertsFromPEM(caCert)
	if !ok {
		log.Fatal(err)
	}

	//创建TLS客户端配置(ssl 3.0 = tls 1.0)
	config := tls.Config{
		RootCAs: certPool, //配置CA池
	}

	//使用TLS配置创建客户端
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &config,
		},
	}

	//client = http.Client{}

	//发起GET请求，获取响应包
	response, err := client.Get("https://localhost:8085")
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取响应包的body数据
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	//打印body数据
	fmt.Printf("%s\n", body)

}

/*
	HTTPS双向认证（客户端认证服务器,服务器也得认证客户端）
*/

// func main() {
// 	//注册服务器的自签名证书（根证书）
// 	//读取
// 	caCert, err := ioutil.ReadFile("../server/server.crt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//创建CA池
// 	certPool := x509.NewCertPool()

// 	//将证书添加到CA池
// 	ok := certPool.AppendCertsFromPEM(caCert)
// 	if !ok {
// 		log.Fatal(err)
// 	}

// 	//加载客户端自己的证书和私钥
// 	//证书需要传递给服务器
// 	//私钥是为了解开服务器用客户端的公钥加密的信息
// 	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//创建TLS客户端配置
// 	config := tls.Config{
// 		RootCAs:      certPool, //配置CA池
// 		Certificates: []tls.Certificate{cert},
// 	}

// 	//使用TLS配置创建客户端
// 	client := http.Client{
// 		Transport: &http.Transport{
// 			TLSClientConfig: &config,
// 		},
// 	}

// 	//client = http.Client{}

// 	//发起GET请求，获取响应包
// 	response, err := client.Get("https://localhost:8085")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	//获取响应包的body数据
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	//打印body数据
// 	fmt.Printf("%s\n", body)
// }
