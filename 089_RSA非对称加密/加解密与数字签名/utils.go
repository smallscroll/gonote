package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

//ReadRSAPublicKey 从文件获取公钥
func ReadRSAPublicKey(filename string) (*rsa.PublicKey, error) {

	//读取公钥文件
	info, err := ioutil.ReadFile(filename) //读取全部文件内容
	if err != nil {
		return nil, err
	}
	//解码获取block
	block, _ := pem.Decode(info)
	//获取DER
	der := block.Bytes
	//获取公钥
	pub, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, err
	}
	//断言获得公钥
	publickKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("publickKey not ok")
	}

	return publickKey, nil
}

//ReadRSAPrivateKey 从文件获取私钥
func ReadRSAPrivateKey(filename string) (*rsa.PrivateKey, error) {

	//读取私钥文件
	info, err := ioutil.ReadFile(filename) //读取全部文件内容
	if err != nil {
		return nil, err
	}
	//解码获取block
	block, _ := pem.Decode(info)
	//获取DER
	der := block.Bytes
	//获取私钥
	key, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}
	//断言获得私钥
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("privateKey not ok")
	}

	return privateKey, nil
}
