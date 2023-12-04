package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

// 生成private.key(私钥文件），public.key(公钥文件)，以及private.pem和public.pem文件
func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var publicKey *rsa.PublicKey
	publicKey = &privateKey.PublicKey
	privateKeyFile, err := os.Create("E:/Geek/src/learn.go/buchong/12bianyi/shili13_RSA_AES/shili_RSA/shengchengpem/priveta.key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	privateKeyEncoder := gob.NewEncoder(privateKeyFile)
	privateKeyEncoder.Encode(privateKey)
	privateKeyFile.Close()
	publicKeyFile, err := os.Create("E:/Geek/src/learn.go/buchong/12bianyi/shili13_RSA_AES/shili_RSA/shengchengpem/public.key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	publicKeyEncoder := gob.NewEncoder(publicKeyFile)
	publicKeyEncoder.Encode(publicKey)
	publicKeyFile.Close()

	// 生成private.pem
	pemPrivateFile, err := os.Create("E:/Geek/src/learn.go/buchong/12bianyi/shili13_RSA_AES/shili_RSA/shengchengpem/private.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var pemPrivateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(pemPrivateFile, pemPrivateKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemPrivateFile.Close()

	// 生成public.pem(自己尝试模仿上面的)
	pemPublicFile, err := os.Create("E:/Geek/src/learn.go/buchong/12bianyi/shili13_RSA_AES/shili_RSA/shengchengpem/public.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var pemPublicKey = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	err = pem.Encode(pemPublicFile, pemPublicKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemPublicFile.Close()
}
