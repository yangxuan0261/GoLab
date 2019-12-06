package test_encrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"testing"
)

func Test_hmac_01(t *testing.T) {
	// 明文
	in := []byte("Sample message for keylen<blocklen")
	// 密钥
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f}
	// 输出
	out := "a28cf43130ee696a98f14a37678b56bcfcbdd9e5cf69717fecf5480f0ebdf790"
	// 使用HMAC-SHA256得到哈希运算消息认证码，其他如MD5, SHA1, SHA224, SHA384, SHA512
	h := hmac.New(sha256.New, key)
	n, err := h.Write(in)
	if n != len(in) || err != nil {
		fmt.Printf("--- Write(%d) = %d, %v", len(in), n, err)
	}
	sum := fmt.Sprintf("%x", h.Sum(nil))
	if sum != out {
		fmt.Println("--- Sum error.")
	}
}

func Test_hmac_02(t *testing.T) {

}
