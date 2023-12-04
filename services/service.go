package services

import (
	"encoding/json"
	"log"
	"net/url"
	"pandora-token-helper/utils"
)

type Service struct {
	BaseUrl string
}

func (s *Service) GetAccessTokenBySessionToken(sessionToken string) (string, error) {

	values := url.Values{}
	values.Add("session_token", sessionToken)
	apiUrl := s.BaseUrl + "/auth/session"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	request, err := utils.SendRequest("POST", apiUrl, headers, []byte(values.Encode()))
	if err != nil {
		return "", err
	}
	var m map[string]interface{}
	err = json.Unmarshal(request, &m)
	if err != nil {
		return "", err
	}
	accessToken := m["access_token"].(string)
	return accessToken, nil
}

func (s *Service) GetAccessTokenByAccount(account, password string) (string, error) {

	values := url.Values{}
	values.Add("username", account)
	values.Add("password", password)
	// 构建一个切片用于保存键值对的字符串
	apiUrl := s.BaseUrl + "/auth/login"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	request, err := utils.SendRequest("POST", apiUrl, headers, []byte(values.Encode()))
	if err != nil {
		return "", err
	}
	log.Println(string(request))
	var m map[string]interface{}
	err = json.Unmarshal(request, &m)
	if err != nil {
		return "", err
	}
	accessToken := m["access_token"].(string)
	return accessToken, nil
}

func (s *Service) RefreshShareToken(accessToken, uniqueName string) (string, error) {

	values := url.Values{}
	values.Add("access_token", accessToken)
	values.Add("unique_name", uniqueName)

	apiUrl := s.BaseUrl + "/token/register"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	request, err := utils.SendRequest("POST", apiUrl, headers, []byte(values.Encode()))
	if err != nil {
		return "", err
	}
	return string(request), nil

}
