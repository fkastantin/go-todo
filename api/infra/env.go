package infra

import (
    "log"
    "github.com/joho/godotenv"
)


func LoadEnvVariables() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Can not load environment variables!")
        panic("Can not load environment variables!")
    }
}