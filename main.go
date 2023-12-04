package main

import (
	"log"
	"pandora-token-helper/services"
)

func RefreshTokenBySessionToken(baseUrl, sessionToken, uniqueName string) {
	var s services.Service
	s.BaseUrl = baseUrl
	accessToken, err := s.GetAccessTokenBySessionToken(sessionToken)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(accessToken)
	shareToken, err := s.RefreshShareToken(accessToken, uniqueName)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(shareToken)
}

func RefreshTokenByAccount(baseUrl, account, password, uniqueName string) {
	var s services.Service
	s.BaseUrl = baseUrl
	accessToken, err := s.GetAccessTokenByAccount(account, password)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("accessToken:", accessToken)
	shareToken, err := s.RefreshShareToken(accessToken, uniqueName)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("shareToken,", shareToken)
}

func main() {
	baseUrl := "https://dd.18.1/pr156456"
	sessionToken := "eyJ"
	uniqueName := "12587587"

	RefreshTokenBySessionToken(baseUrl, sessionToken, uniqueName)
	//RefreshTokenByAccount("", "", "", "")
}
