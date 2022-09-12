package persist

import (
	"database/sql"
	"time"
)

type RequestInfo struct {
	Id          string
	Url         string
	SearchKey   string
	RequestType string
	CreatedTime time.Time
}

func Insert(requestInfo *RequestInfo) {
	db, err := sql.Open("sqlite3", "./data/data.db")

	//插入数据
	stmt, err := db.Prepare("INSERT INTO request_info(id,search_key, url, request_type, created_time) values (?,?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(requestInfo.Id, requestInfo.SearchKey, requestInfo.Url, requestInfo.RequestType, requestInfo.CreatedTime)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
