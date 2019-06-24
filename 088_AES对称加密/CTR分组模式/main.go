package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

/*
	AES-CTR模式：
		AES（加密算法）: 密钥长度：16/24/32字节，分组长度：16字节
		CTR（分组模式）: 分组模式： 对初始向量进行加密，向量长度与算法长度一致，因此不需要填充明文

*/

//AESCTREncrypt AES-CTR模式加密
func AESCTREncrypt(text, key []byte) ([]byte, error) {

	//创建一个cipher.Block接口（参数：密钥），返回一个分组接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//密文数据：数据长度=向量长度+明文长度（向量长度须与算法长度一致）
	ciphertext := make([]byte, aes.BlockSize+len(text))

	//向量：向量长度=分组长度
	iv := ciphertext[:aes.BlockSize]                        //截取出向量
	if _, err := io.ReadFull(rand.Reader, iv); err != nil { //按长度读取随机数到向量
		return nil, err
	}

	//创建分组（参数1：分组接口，参数2：初始向量），返回一个Stream接口
	stream := cipher.NewCTR(block, iv)

	//加密（参数1：密文空间，参数2：明文），密文长度须与明文一致
	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)

	//返回密文数据（向量+密文）
	return ciphertext, nil
}

//AESCTRDecrypt AES-CTR模式解密
func AESCTRDecrypt(text, key []byte) ([]byte, error) {

	//创建一个cipher.Block接口（参数：密钥），返回一个分组接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//截取密文数据获得向量：向量长度=分组长度
	iv := text[:aes.BlockSize]

	//创建分组（参数1：分组接口，参数2：初始向量），返回一个Stream接口
	stream := cipher.NewCTR(block, iv)

	//明文数据：明文长度=密文长度
	plaintext := text[aes.BlockSize:]

	//解密（参数1：明文空间，参数2：密文），明文长度须与密文一致
	stream.XORKeyStream(plaintext, text[aes.BlockSize:])

	//返回明文数据
	return plaintext, nil
}

func main() {
	//明文
	src := "hello~ 你好。"
	//对称密钥（16字节）
	key := "123456781234567812345678"

	//加密
	ciphertext, err := AESCTREncrypt([]byte(src), []byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%x\n", ciphertext) //以16进制打印密文

	//解密
	plaintext, err := AESCTRDecrypt(ciphertext, []byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", plaintext) //以字符串打印明文
}
