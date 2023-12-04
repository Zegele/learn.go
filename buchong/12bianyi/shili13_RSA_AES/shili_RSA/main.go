//RSA
//www.kancloud.cn/imdszxs/golang/1509752

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// AES一般用于加密密文，而RSA算法一般用来加解密密码 ，如下：

//可通过openssl产生
// openssl genrsa -out rsa_private_key.pem 1024
// 关于怎样生成.pem文件，参考9bingfa/shili29_jiamitongxun

var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC2AmxPsJUSr6XclbVatRU8hWuRgSlkr+dW/ya4iC87ThjmLhiu
XQSM3hJ5Wv+jlvchvfvcuztSzLF3ziAndJxCYNxiZNFe1/kUbpnvg3hwTU3g0o+y
AkAaKn2f5yS5MrLqXfSmMFxi5wTaZEhZDhdk1IQkfWZkHKpbyC5rFuBrgwIDAQAB
AoGAbsp4zeVqSO6covKKa0WxQJ7ihLa/0Kd2xbEZC9jyD0KRy4OhyHmcrHb9fduT
oyvhESwJZSqfiN5K5OeXEQ1dk8903DMocGKynstDYjqZLiBEGzXJV2XcEKTGRzYi
aQ0xJyFltk3AcR9Oo7IBjaZDA3NDVDBpXdFnAzkl7naAqMECQQDcCOEZPM1208bh
efp6lgS3mDAiUrjxnceLTCJTfynURX9x02e7CBHQ5uI46B5noCu72f7kHBPsEI3Z
33+jJ3DhAkEA08JrGR/eROJtLqZx6bHEXq5YjKHX2TMahiXKyRJyAAKllAKmGvvU
zOLKUIyHbSmfQX+cag5u5cCHKvMSxg3U4wJAH5K2miR7Zx2kYB5crtJwAtg3r/Um
zKTTSU23bzvECM2gJ/kp3VCfHdbDh17nXf9Bx0bUqG7O/QSzFhxX+FVkgQJBALCj
4MqYdsRnTEQ5u5QBIZkJPfwxlOgtZNa596o6pLW1f4EzpHOl1iECnVkLoHxC2AG+
/S7K0177dsUw9lwXdZkCQCgDPOoZP0LaVd9MQyZ/GI1vtYgYVPdld0Yp2SkjzrxK
22EM/GG6tPVtGztq068geKIvcrqBQfxnj6EFZyElnqE=
-----END RSA PRIVATE KEY-----
`)

// 上面的private.pem是shengchengpem文件夹的代码生成的。

// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBALYCbE+wlRKvpdyVtVq1FTyFa5GBKWSv51b/JriILztOGOYuGK5dBIze
Enla/6OW9yG9+9y7O1LMsXfOICd0nEJg3GJk0V7X+RRume+DeHBNTeDSj7ICQBoq
fZ/nJLkysupd9KYwXGLnBNpkSFkOF2TUhCR9ZmQcqlvILmsW4GuDAgMBAAE=
-----END RSA PUBLIC KEY-----`)

// 上面这个命令是什么？没懂
// 上面的public.pem是shengchengpem文件夹的代码生成的。

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	//解析公钥
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	//解析格式必须和加密格式一致
	//ParsePKCS1PublicKey返回的直接是*rsa.PublicKey类型

	//pubInterface2, err := x509.ParsePKIXPublicKey(block.Bytes)
	//ParsePKIXPublicKey返回值是一个接口类型，后面要断言是不是*rsa.PublicKey类型

	if err != nil {
		return nil, err
	}
	// 类型断言
	//pub := pubInterface2.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func main() {
	//fmt.Println("lkklkdfyit")
	data, _ := RsaEncrypt([]byte("golanggolanggolang")) //原文加密，获得密文
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	origData, _ := RsaDecrypt(data) // 密文解密，获得原文
	fmt.Println(string(origData))
	//为什么打印不出任何东西？？？
}
