package md5

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
	"os"
	"reflect"
)

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 将任意类型的变量进行md5摘要(注意map等非排序变量造成的不同结果)
func Encode(v interface{}) string {
	h := md5.New()
	if "string" == reflect.TypeOf(v).String() {
		h.Write([]byte(v.(string)))
	} else {
		b, err := json.Marshal(v)
		if err != nil {
			return ""
		} else {
			h.Write(b)
		}
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 将字符串进行MD5哈希摘要计算
func EncodeString(v string) string {
	h := md5.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 将文件内容进行MD5哈希摘要计算
func EncodeFile(path string) string {
	f, e := os.Open(path)
	if e != nil {
		fmt.Errorf("%s", e)
	}
	h := md5.New()
	_, e = io.Copy(h, f)
	if e != nil {
		fmt.Errorf("%s", e)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// GenRandomString 生成随机字符串
// length 生成长度
// specialChar 是否生成特殊字符
func GenRandomString(length int, specialChar bool) string {

	letterBytes := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special := "!@#%$*.="

	if specialChar {
		letterBytes = letterBytes + special
	}

	chars := []byte(letterBytes)

	if length == 0 {
		return ""
	}

	clen := len(chars)
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			return ""
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

// CryptPassword 加密密码
func CryptPassword(password, salt string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	return fmt.Sprintf("%x", dk)
}
