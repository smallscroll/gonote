package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

/*
	RSA加解密与数字签名
*/

//RSAEncrypt 使用公钥进行加密(传入公钥和明文)
func RSAEncrypt(filename string, src []byte) ([]byte, error) {

	//读取公钥
	publickKey, err := ReadRSAPublicKey(filename)
	if err != nil {
		return nil, err
	}
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publickKey, src)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

//RSADecrypt 使用私钥进行解密(传入私钥和密文)
func RSADecrypt(filename string, src []byte) ([]byte, error) {

	//读取私钥
	privateKey, err := ReadRSAPrivateKey(filename)
	if err != nil {
		return nil, err
	}
	//解密
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {

	//原文
	src := []byte("hello~ 你好。")

	/*
		RSA加密与解密
	*/

	//加密
	ciphertext, err := RSAEncrypt("public_key.pem", src)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("密文：%x\n", ciphertext) //以16进制打印密文

	//解密
	plaintext, err := RSADecrypt("private_key.key", ciphertext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("明文：%s\n", plaintext) //以字符串打印明文

	/*
		数字签名与校验
	*/

	//对原文进行数字签名
	signature, err := RSASignature("private_key.key", src)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("签名：%x\n", signature) //以16进制打印签名

	//校验原文的数字签名
	src = []byte("hello~") //此处篡改了原文

	if err := RSAVerify("public_key.pem", src, signature); err == nil {
		fmt.Println("签名有效")
	} else {
		fmt.Println(err)
		fmt.Println("签名无效")
		return
	}

}
