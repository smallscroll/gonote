package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

/*
	RSA数字签名与校验
*/

//RSASignature 数字签名（私钥，原文内容）
func RSASignature(filename string, src []byte) ([]byte, error) {

	//读取私钥
	privateKey, err := ReadRSAPrivateKey(filename)
	if err != nil {
		return nil, err
	}
	//计算原文的哈希值
	hash := sha256.Sum256(src)

	//对哈希值进行签名（随机数，私钥，计算哈希的方法，原文的哈希值）
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

//RSAVerify 签名校验（公钥，待校验的内容，签名）
func RSAVerify(filename string, src []byte, signnature []byte) error {

	//读取公钥
	publickKey, err := ReadRSAPublicKey(filename)
	if err != nil {
		return err
	}

	//计算待校验内容的哈希值
	hash := sha256.Sum256(src)

	//签名验证（公钥，计算哈希的算法，待校验内容的哈希值，待验证的签名）
	err = rsa.VerifyPKCS1v15(publickKey, crypto.SHA256, hash[:], signnature)
	if err != nil {
		return err
	}

	return nil
}
