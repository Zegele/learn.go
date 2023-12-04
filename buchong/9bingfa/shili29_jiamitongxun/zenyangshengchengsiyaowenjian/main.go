// 怎样生成公私钥文件
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

//www.zongscan.com/demo333/96009.html
// 如何将RSA私钥和公钥保存到文件中？
// 本例使用RSA，也可以保存其他类型的加密系统的公私钥
//参考资料 :
//
//code.google.com/p/go/source/browse/src/pkg/encoding/pem/pem_test.go

func main() {
	// generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var publickey *rsa.PublicKey
	publickey = &privatekey.PublicKey
	//  save private and public key separately
	privatekeyfile, err := os.Create("private.key") // 生成的该文件，直接在大项目目录下
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	privatekeyencoder := gob.NewEncoder(privatekeyfile)
	privatekeyencoder.Encode(privatekey)
	privatekeyfile.Close()
	publickeyfile, err := os.Create("public.key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	publickeyencoder := gob.NewEncoder(publickeyfile)
	publickeyencoder.Encode(publickey)
	publickeyfile.Close()

	// save PEM file
	pemfile, err := os.Create("private.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// http://golang.org/pkg/encoding/pem/#Block
	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey),
	}
	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pemfile.Close()

}

//private.pem的输出样本:
//
//-----BEGIN RSA PRIVATE KEY-----
//MIICXAIBAAKBgQC9/H127GOFgwROUPTGdpw6fCL5DGvpldLALkr6o7ia/6mvioVT
//B+ivUwgo/WXyIHHPEtKNtnKH4OUcav/TNiXDqV+q9ybZ6SCodHPtryJzyy3iMhRz
//9WxfAlKrF+XmiC8hzJsXjfE3o0sjacrEdlxn+zmxu5H2weSFh4ZjXCBtzQIDAQAB
//AoGAe1QfYga76ByPvAMjkn3GltSko0Uj/CMNB0JF3ARRvxR9430pZSf6LW3aGzm7
//Zv0WxBR06Bdqq7gbImJ3JXW99vAqUUseLuR6KQ78YvZDkNz4aKnXCFBvJmtCVTj9
//SPyY0KoKjeR7slgdik0CbrisqrlFOk+eO9Bj7Wd40p14SGECQQDCUp2/t9qS5v1O
//C+1xkRZ+BDJ1+WBoGHDXFPVchbfYdZSHtlurTEdN8g4MaHohY7d7/QFUJSXOh4M8
//utS0TmTpAkEA+kmGLrrAO+SDNJMcNi/w+m1qO1o3acewUbZI04tyqon583O+rVzB
//Lo3iTbErKTK/1HA0+Brqjp7xe6lS0PyDRQJARjdWGw2TJFvlEcuLi+rSRsy7cxee
//N18FfyJqmnkS+ltaRUOmkhoo9chOPTuPTftbNKkyTrZxl9Qtnscfzts46QJBAIf2
///Q/Rn7BpmOUsrXy6WnyQh88qWUP7mMsq7TEOZgJC5ifczs66vq8doLx37Gx7Bz7O
//ndfSN2225pQ5DaY+JskCQHIrr7XOX0Ka0FmZcyvsDI0YzySvATrDnYxa3gj4YBSW
//eaV15LvMgibF3khLjXsuR8kUNd58NB/uAyeXuGrF2sI=
//-----END RSA PRIVATE KEY-----
