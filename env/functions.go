package env

import (
	"os"
	"server/apis/jwt"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Name        string
	Port        string
	Downloads   string
	BodyLimit   int
	Concurrency int
)

func LoadEnv() (err error) {
	err = godotenv.Load("./.env")
	if err != nil {
		return
	}
	Name = os.Getenv("NAME")
	Port = os.Getenv("PORT")
	Downloads = os.Getenv("DOWNLOADS")
	Concurrency, _ = strconv.Atoi(os.Getenv("CONCURRENCY"))
	BodyLimit, _ = strconv.Atoi(os.Getenv("BODYLIMIT"))
	err = jwt.Init()
	if err != nil {
		return
	}
	err = os.MkdirAll(Downloads, os.ModePerm)
	if err != nil {
		return
	}
	return
}
