package main

import (
	"log"
	"pandora-token-helper/services"
)

func RefreshShareTokenBySessionToken(baseUrl, sessionToken, uniqueName string) {
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

func RefreshShareTokenByAccount(baseUrl, account, password, uniqueName string) {
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
	baseUrl := "https://yourip.or.domain:port/<your_proxy_api_prefix>"
	uniqueName := "12587587"


	sessionToken := "eyJ"
	RefreshShareTokenBySessionToken(baseUrl, sessionToken, uniqueName)

	account:="your_account"
	password:="your_password"
	RefreshShareTokenByAccount(baseUrl, account, password,uniqueName)
}
