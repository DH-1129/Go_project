package funcs

import "golang.org/x/crypto/bcrypt"

func Encrypthon_PW(Old_PW string) (New_PW string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Old_PW), bcrypt.DefaultCost)
	if err != nil {
		Danger(err)
	}
	encodePW := string(hash)
	return encodePW
}
