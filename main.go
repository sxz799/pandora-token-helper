package main

import (
	"gpt-token-helper/services"
	"log"
)

func RefreshTokenBySessionToken(baseUrl, sessionToken, uniqueName string) {
	var s services.Service
	s.BaseUrl = baseUrl
	accessToken, err := s.GetAccessTokenBySessionToken(sessionToken)
	log.Println(accessToken)
	shareToken, err := s.RefreshShareToken(accessToken, uniqueName)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(shareToken)
}

func RefreshTokenByAccount(baseUrl, account, password, uniqueName string) {
	var s services.Service
	s.BaseUrl = baseUrl
	accessToken, err := s.GetAccessTokenByAccount(account, password)
	log.Println(accessToken)
	shareToken, err := s.RefreshShareToken(accessToken, uniqueName)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(shareToken)
}

func main() {
	RefreshTokenByAccount("", "", "", "")
}
