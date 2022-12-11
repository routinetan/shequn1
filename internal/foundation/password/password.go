package password

import (
	"crypto/md5"
	"fmt"
)

// Verify 密码hash验证
func Verify(password, hash string) bool {
	hexStr := md5.Sum([]byte(password))
	md5Str := fmt.Sprintf("%x\n", hexStr)
	if md5Str == hash {
		return true
	} else {
		return false
	}
}

func Hash(in string) string {
	hexStr := md5.Sum([]byte(in))
	return fmt.Sprintf("%x\n", hexStr)
}
