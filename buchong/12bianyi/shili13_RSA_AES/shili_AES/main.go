// Go语言实现RSA和AES加解密
// www.kancloud.cn/imdszxs/golang/1509752
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//AES 加解密
// AES加密分为ECB，CBC，CFB，OFB等

/*
//1. CBC加密
func main() {
	orig := "1kjghkhnh'"
	key := "1234567890123456" //len不能少于16？？？
	fmt.Println("原文：", orig)
	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)
	decryptCode := AesDecrypt(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
	//orig := "http://c.biancheng.net/golang/"
	//key := "123456781234567812345678"
	//fmt.Println("原文：", orig)
	//encryptCode := AesEncrypt(orig, key)
	//fmt.Println("密文：", encryptCode)
	//decryptCode := AesDecrypt(encryptCode, key)
	//fmt.Println("解密结果：", decryptCode)
}

func AesEncrypt(orig string, key string) string {
	// 转成字节切片
	origData := []byte(orig)
	fmt.Println("origDataLen:", len(origData), "origLen:", len(orig))
	k := []byte(key)
	// 分组密钥
	block, _ := aes.NewCipher(k)
	//获取密钥块的长度
	blockSize := block.BlockSize()
	fmt.Println(blockSize)
	// 补充全码
	origData = PKCS7Padding(origData, blockSize)
	//加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建切片
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

func AesDecrypt(cryted string, key string) string {
	// 转成字节切片
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组密钥
	block, _ := aes.NewCipher(k)
	// 获取密钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建切片
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	// 16-10%16
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去补码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

*/

//2.CFB加密（非官方示例）

/*
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
func main() {
	//需要加密的字符串
	plaintext := []byte("golang")
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}
	//aes的加密字符串
	key_text := "asgdlhgaosidlkbvnb;lakdjf;ljge32" // 必须是32位？

	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)
	//解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}

*/

//2.2CFB加密（官方示例）

// 加密
func ExampleNewCFBEncrypter() {
	//生成key
	key, _ := hex.DecodeString("12345678901234567890123456789012") // 这个key要和解密的key一致 32位的

	//加密的内容
	plaintext := []byte("golanggolanggolanggolanggolang")
	block, err := aes.NewCipher(key) //把key打包？
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	//reader := rand.Reader // 是加密安全随机数生成的全局共享实例。
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		//io.ReadFull(rand.Reader, iv) 阅读器Reader，随机加密的reader把iv装满，这个iv应该就是所谓的摘要?
		//参考：blog.csdn.net/huang_2016/article/details/124639763
		panic(err)
	}
	fmt.Println(iv)
	stream := cipher.NewCFBEncrypter(block, iv)                // 把key的打包，把前缀（iv）一起加密
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext) //对内容开始加密
	fmt.Printf("%x\n", ciphertext)                             //ciphertext就是网络传输中的密文？
	// 每运行一次，因为有rand.Reader，所以每次生成的ciphertext都不一样，但是如果可以正确解密，都会解出同样的明文。
}

// 解密
func ExampleNewCFBDecrypter() {
	key, _ := hex.DecodeString("12345678901234567890123456789012") // 这个key要和加密的key一致
	//ciphertext, _ := hex.DecodeString("0015953b4c911b10b9745b626a2ffe69b004521c24760748d81e51abbecf60297255cf3f69932ae3ce97fc5a5f83")
	ciphertext, _ := hex.DecodeString("47d43b4d3932d3c93aeb114724b3dea62032ee5dca10390587827f310d8b1996a38553c78bdec553e58e0914ef95")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize { // 因为ciphertext的长度是aes.BlockSize+明文长度，所以ciphertext不可能小于aes.BlockSize
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv) //通过block和前缀（iv）生成解密的stream

	//stream.XORKeyStream(ciphertext, ciphertext)//一般这样写
	//fmt.Printf("%s\n", ciphertext)

	a := []byte("golanggolanggolanggolanggolanl123456") // 123456是超过ciphertext的部分
	stream.XORKeyStream(a, ciphertext)                  //是将密文ciphertext，解密后写入a。 对a的要求是最好和ciphertext的len一样长。不能短。但如果长度超过ciphertext，则再解密后，a中的超过部分还在。
	fmt.Printf("%s\n", a)
}

func main() {
	ExampleNewCFBEncrypter()
	ExampleNewCFBDecrypter()
}
