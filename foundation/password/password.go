package password

import (
	"golang.org/x/crypto/bcrypt"
	"quanzi1/foundation/app"
	"quanzi1/foundation/server"
)

var passwordToken = server.Config.PasswordToken

// Hash 密码hash
func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordToken+password), bcrypt.DefaultCost)
	if err != nil {
		app.Logger().WithField("log_type", "foundation.password.password").Error(err)
		return ""
	}

	return string(bytes)
}

// Verify 密码hash验证
func Verify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordToken+password))
	return err == nil
}
