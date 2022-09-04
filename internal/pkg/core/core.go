package core

import (
	"fmt"
	"mimi/pkg"
	"regexp"
	"strings"
)

func NewIntercept(port, key string) *pkg.Intercept {
	re := regexp.MustCompile(key)
	return &pkg.Intercept{
		KeyWord: key,
		Ip:      "0.0.0.0:" + port,
		HttpPackageFunc: func(pack *pkg.HttpPackage) {
			if strings.Index(pack.ContentType, "json") != -1 && re.MatchString(pack.Json()) {
				fmt.Printf("json: url = %s\r\n ", pack.Url.String())
			}
			if strings.Index(pack.ContentType, "html") != -1 && re.MatchString(pack.Html()) {
				fmt.Printf("html: url = %s\r\n ", pack.Url.String())
			}
		},
	}
}
