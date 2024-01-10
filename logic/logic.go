package logic

import (
	"effective-garbanzo/logic/database"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LogicModule struct {
	Database database.DatabaseModule
}

func GenUUID() string {
	return uuid.NewString()
}

func PasswordHash(source string) (res string, err error) {
	resbyte, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	res = string(resbyte)
	return
}
func ComparePasswordHash(sourcehash string, input string) (res bool) {
	err := bcrypt.CompareHashAndPassword([]byte(sourcehash), []byte(input))
	res = (err == nil)
	return
}
