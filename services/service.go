package services

import (
	"gpt-token-helper/utils"
)

type Service struct{}

var baseUrl = "https://ai.fakeopen.com"

func (s *Service) GetAccessTokenBySessionToken(sessionToken string) (string, error) {

	// 构建一个切片用于保存键值对的字符串
	body := []byte("session_token=" + sessionToken)
	url := baseUrl + "/auth/session"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	request, err := utils.SendRequest("POST", url, headers, body)
	if err != nil {
		return "", err
	}
	return string(request), nil
}

func (s *Service) RefreshShareToken(accessToken, uniqueName string) (string, error) {

	// 构建一个切片用于保存键值对的字符串
	body := []byte("access_token=" + accessToken + "&unique_name=" + uniqueName)
	url := baseUrl + "/token/register"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	request, err := utils.SendRequest("POST", url, headers, body)
	if err != nil {
		return "", err
	}
	return string(request), nil

}
