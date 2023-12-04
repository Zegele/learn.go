// 怎样加载公私钥文件
package main

import (
	"bufio"
	"crypto/rsa"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

//参考：www.zongscan.com/demo333/96075.html

func main() {
	//Load Private Key
	privatekeyfile, err := os.Open("private.key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	decoder := gob.NewDecoder(privatekeyfile)
	var privatekey rsa.PrivateKey
	err = decoder.Decode(&privatekey)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	privatekeyfile.Close()
	fmt.Printf("Private Key :\n%x\n", privatekey)

	//Public Key
	publickeyfile, err := os.Open("public.key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	decoder = gob.NewDecoder(publickeyfile)
	var publickey rsa.PublicKey
	err = decoder.Decode(&publickey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	publickeyfile.Close()
	fmt.Printf("Public key : \n%x\n", publickey)

	// Load PEM
	pemfile, err := os.Open("private.Pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// need to convert pemfile to []byte for decoding
	pemfileinfo, _ := pemfile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	//read pemfile content into pembytes
	buffer := bufio.NewReader(pemfile)
	_, err = buffer.Read(pembytes)
	// proper decoding now
	data, _ := pem.Decode([]byte(pembytes))
	pemfile.Close()
	fmt.Printf("PEM Type :\n%s\n", data.Type)
	fmt.Printf("PEM Headers :\n%s\n", data.Headers)
	fmt.Printf("PEM Bytes :\n%x\n", string(data.Bytes))
}
