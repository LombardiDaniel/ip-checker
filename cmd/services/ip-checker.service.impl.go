package services

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

const ipFilePath 		string = "./last_ip.txt"
const checkIpUrl	string = "https://checkip.amazonaws.com"

type IpCheckerServiceImpl struct {
	httpClient 			*http.Client
}

func NewIpCheckerImpl() IpCheckerService {
	return &IpCheckerServiceImpl{
		httpClient: &http.Client{},
	}
}

func (s *IpCheckerServiceImpl) GetCurrIp() (*string, error) {
	resp, err := s.httpClient.Get(checkIpUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        slog.Error(fmt.Sprintf("Error reading response body: '%s'", err))
        return nil, err
    }

	if resp.StatusCode != 200 {
		slog.Error(fmt.Sprintf("StatusCode != 200: '%d', body: %s", resp.StatusCode, string(body)))
        return nil, errors.New("return code != 200")
	}

	bodyStr := string(body[:len(body)-1]) // we cut out the '\n' from the msg
	return &bodyStr, nil
}

func (s *IpCheckerServiceImpl) StoreIp(ip string) error {
	err := os.WriteFile(ipFilePath, []byte(ip), 0644)
    if err!= nil {
        return err
    }

    return nil
}

func (s *IpCheckerServiceImpl) ReadOldIp() (*string, error) {

	_, err := os.Stat(ipFilePath)
	if errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(ipFilePath)
		if err != nil {
			err = errors.Join(err, errors.New("could not create file"))
			return nil, err
		}
	}

	data, err := os.ReadFile(ipFilePath)
    if err!= nil {
        return nil, err
    }

    ip := string(data)

    return &ip, nil
}
