package auth

import (
	"errors"
	"net/http"
	"strings"
)


func GetAPIKey(headers http.Header)(string, error){
	val := headers.Get("Authorization");
	if val == ""{
		return "", errors.New("no Authorization header found");
	}
	vals := strings.Split(val, " ");
	if len(vals) != 2 {
		
		return "", errors.New("authurization header malformed");
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("authurization header malformed");
	}
	return vals[1], nil;
}