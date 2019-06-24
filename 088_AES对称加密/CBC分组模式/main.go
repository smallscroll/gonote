package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

/*

	AES-CBC模式：
		AES（加密算法）: 16字节，分组长度：16字节
		CBC（分组模式）: 对明文分组进行加密，分组长度与算法长度一致，因此需要填充明文


*/

//AESCBCEncrypter AES-CBC模式加密
func AESCBCEncrypter(text, key []byte) ([]byte, error) {

	plaintext := text

	//填充内容保证明文长度是分组长度的倍数：
	//填充内容的长度 = 分组长度 - 分组之后剩余的长度
	padlenght := aes.BlockSize - len(plaintext)%aes.BlockSize
	//填充内容 = 需要填充内容的长度数值转化为字符并按照数值进行重复（例：需要填充12个字符，那么将“12”重复12次追加到明文，以便解密时能准确获取到明文已填充的内容长度并作正确去除）
	padtext := bytes.Repeat([]byte{byte(padlenght)}, padlenght)
	//将填充内容追加到明文数据切片
	plaintext = append(plaintext, padtext...)

	//创建分组接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//密文数据：数据长度 = 向量长度 + 明文长度
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	//向量：向量长度=分组长度
	iv := ciphertext[:aes.BlockSize]                        //截取出向量
	if _, err := io.ReadFull(rand.Reader, iv); err != nil { //按长度读取随机数到向量
		return nil, err
	}
	//创建分组
	mode := cipher.NewCBCEncrypter(block, iv)
	//加密
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

//AESCBCDecrypter AES-CBC模式解密
func AESCBCDecrypter(text, key []byte) ([]byte, error) {

	ciphertext := text

	//创建分组接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//截取密文数据获得向量：向量长度=分组长度
	iv := ciphertext[:aes.BlockSize]

	//创建分组
	mode := cipher.NewCBCDecrypter(block, iv)

	//明文数据：明文长度=密文长度
	plaintext := make([]byte, len(ciphertext[aes.BlockSize:]))

	//解密
	mode.CryptBlocks(plaintext, ciphertext[aes.BlockSize:])

	//去除明文的除填充内容
	padlenght := plaintext[len(plaintext)-1]   //获取填充内容的长度（填充时追加的内容值）
	num := int(padlenght)                      //将填充的长度转换为数值
	return plaintext[:len(plaintext)-num], nil //按填充的长度值截取明文切片后返回

}

func main() {
	//明文
	src := "hello~ 你好。"
	//对称密钥（16字节）
	key := "1234567812345678"

	//加密
	ciphertext, err := AESCBCEncrypter([]byte(src), []byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%x\n", ciphertext) //以16进制打印密文

	//解密
	plaintext, err := AESCBCDecrypter(ciphertext, []byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", plaintext) //以字符串打印明文
}
