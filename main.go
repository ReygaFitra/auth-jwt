package main

import (
	authDatabase "github.com/ReygaFitra/auth-jwt/database"
)

func main() {
	authDatabase.ConnectDB()
}