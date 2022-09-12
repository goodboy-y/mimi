package core

import (
	"fmt"
	"github.com/google/uuid"
	"mimi/internal/persist"
	"mimi/pkg"
	"regexp"
	"strings"
	"time"
)

func NewIntercept(port, key string, isPersist bool) *pkg.Intercept {
	re := regexp.MustCompile(key)
	return &pkg.Intercept{
		KeyWord: key,
		Ip:      "0.0.0.0:" + port,
		HttpPackageFunc: func(pack *pkg.HttpPackage) {
			var requestInfo string
			if strings.Index(pack.ContentType, "json") != -1 && re.MatchString(pack.Json()) {
				requestInfo = "json"
				fmt.Printf("json: url = %s\r\n ", pack.Url.String())
			}
			if strings.Index(pack.ContentType, "html") != -1 && re.MatchString(pack.Html()) {
				requestInfo = "html"
				fmt.Printf("html: url = %s\r\n ", pack.Url.String())
			}
			if isPersist && requestInfo != "" {
				newUUID, _ := uuid.NewUUID()
				requestInfo := &persist.RequestInfo{
					Id:          newUUID.String(),
					SearchKey:   key,
					Url:         pack.Url.String(),
					RequestType: requestInfo,
					CreatedTime: time.Now(),
				}
				persist.Insert(requestInfo)
			}

		},
	}
}
