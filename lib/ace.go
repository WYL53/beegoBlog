package lib

import (
	//"crypto/aes"
	"bytes"
	//"encoding/base64"
)

//func AesEncryption(bytes,key []byte) ([]byte,error) {
//	cip,err := aes.NewCipher(key)
//	if err != nil{
//		return nil,err
//	}
//	padding := cip.BlockSize()
//	raw := pkcs5Padding(bytes,padding)
//	dest := make([]byte,len(raw))
//	cip.Encrypt(dest,raw)
//	base64.StdEncoding.Encode(dest,raw)
//	return dest,nil
//}

//func AesDecryption(src,key []byte) []byte {
//	cip,err := aes.NewCipher(key)
//	if err != nil{
//		return nil
//	}
//	cip
//}

func pkcs5Padding(src []byte, blockSize int) []byte {
	p := blockSize - (len(src) % blockSize)
	return append(src,bytes.Repeat([]byte(p),p)...)
}

func unpkcs5Padding(src []byte) []byte {
	p := src[len(src)-1]
	return src[:len(src)-int(p)]
}