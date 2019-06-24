/*
	RSA非对称密钥对生成
*/

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateRSAKey(bit int) error {

	/*
		私钥创建
	*/

	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥。
	privateKey, err := rsa.GenerateKey(rand.Reader, bit) //rand.Reader为随机数生成器，生成一个随机数
	if err != nil {
		return err
	}

	//对私钥进行编码，生成DER数据
	derText, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return err
	}

	//将DER数据拼接到pem格式的数据块中
	block := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   derText,
	}

	//创建文件
	file, err := os.OpenFile("private_key.key", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//将pem进行base64编码
	err = pem.Encode(file, &block) //将编码后的字符串输出到文件
	if err != nil {
		return err
	}
	// err = pem.Encode(os.Stdout, &block) //将编码后的字符串直接输出到终端
	// if err != nil {
	// 	return err
	// }

	/*
		公钥创建
	*/

	//通过私钥获取公钥
	publicKey := privateKey.PublicKey

	//对公钥进行编码，生成DER数据
	derPkix, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	//将DER数据拼接到pem格式的数据块中
	block = pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   derPkix,
	}
	//创建文件
	file, err = os.OpenFile("public_key.pem", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//将pem进行base64编码
	err = pem.Encode(file, &block) //将编码后的字符串输出到文件
	if err != nil {
		return err
	}
	// err = pem.Encode(os.Stdout, &block) //将编码后的字符串直接输出到终端
	// if err != nil {
	// 	return err
	// }

	return nil
}

func main() {

	err := generateRSAKey(2048)
	if err != nil {
		fmt.Println(err)
		return
	}

}
