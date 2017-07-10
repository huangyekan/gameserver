package md5

import "crypto/md5"



func Encode(arg string) string {
	h := md5.New()
	h.Write([]byte(arg))
	result := h.Sum(nil)
	return string(result)
}