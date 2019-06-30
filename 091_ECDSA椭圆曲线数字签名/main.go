package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

/*
	ECC椭圆曲线进行数字签名
		ECC：Elliptic Cure Cryptography（椭圆曲线密码学）
		ECDSA：Elliptic Curve Digital Signature Algorithm（椭圆曲线数字签名）
*/

func main() {
	data := "hello, world."             //待签名数据
	hash := sha256.Sum256([]byte(data)) //计算签名

	//创建私钥
	curve := elliptic.P256()                                 //创建曲线
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader) //生成私钥
	if err != nil {
		fmt.Println(err)
		return
	}
	//通过私钥获得公钥
	publicKey := privateKey.PublicKey

	//私钥签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		fmt.Println(err)
		return
	}

	//将r,s字节流拼接为签名进行传输，对端再将r和s从字节流取出
	signature := append(r.Bytes(), s.Bytes()...)

	//...

	//此处篡改了哈希值
	//hash = sha256.Sum256([]byte("hello."))

	fmt.Printf("%x\n", signature)

	//从字节流中获取r和s
	var r1, s1 big.Int
	r1.SetBytes(signature[:len(signature)/2]) //截取前32字节作为r1
	s1.SetBytes(signature[len(signature)/2:]) //截取后32字节作为s1

	//公钥验证
	res := ecdsa.Verify(&publicKey, hash[:], &r1, &s1)
	if res {
		fmt.Println("签名有效")
	} else {
		fmt.Println("签名无效")
	}
}
