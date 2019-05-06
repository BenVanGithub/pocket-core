package util

import (
	"errors"
	"github.com/pokt-network/pocket-core/config"
	"net/http"
	"strings"
	"time"
)

const (
	HTTPS = "https://"
	HTTP  = "http://"
)

func URLProto(query string) (string, error) {
	if strings.Contains(query, "//") {
		query = strings.TrimLeft(query, "//")
	}
	_, err := Ping(HTTPS + query)
	if err != nil {
		_, err := Ping(HTTP + query)
		if err != nil {
			return "", errors.New(err.Error())
		}
		return HTTP + query, nil
	}
	return HTTPS + query, nil
}

func Ping(url string) (int, error) {
	client := http.Client{}
	client.Timeout = time.Duration(config.GlobalConfig().RequestTimeout) * time.Millisecond
	req, err := http.NewRequest("HEAD", url, http.NoBody)
	if err != nil {
		return 0, err
	}
<<<<<<< 57ceb161d287776fc08ba212726bb3bf39a278c6
	if req!=nil {
		if req.Body!=nil{
			defer req.Body.Close()
		}
	}
=======
	defer req.Body.Close()
>>>>>>> fixed nil pointer error
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
<<<<<<< 57ceb161d287776fc08ba212726bb3bf39a278c6
	if resp != nil {
		if resp.Body!=nil{
			defer resp.Body.Close()
		}
	}
=======
	defer resp.Body.Close()
>>>>>>> fixed nil pointer error
	return resp.StatusCode, nil
}