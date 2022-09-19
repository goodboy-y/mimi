package persist

import (
	"database/sql"
	"time"
)

type RequestInfo struct {
	Id          string    `json:"id,omitempty"`
	Url         string    `json:"url,omitempty"`
	SearchKey   string    `json:"searchKey,omitempty"`
	RequestType string    `json:"requestType,omitempty"`
	CreatedTime time.Time `json:"createdTime"`
}

func Insert(requestInfo *RequestInfo) {
	db, err := sql.Open("sqlite3", "./data/data.db")

	//插入数据
	stmt, err := db.Prepare("INSERT INTO request_info(id,search_key, url, request_type, created_time) values (?,?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(requestInfo.Id, requestInfo.SearchKey, requestInfo.Url, requestInfo.RequestType, requestInfo.CreatedTime)
	checkErr(err)
}

func Search(pagesize, pageNo int) []RequestInfo {
	var result []RequestInfo
	db, err := sql.Open("sqlite3", "./data/data.db")

	//插入数据
	stmt, err := db.Prepare("select id,search_key, url, request_type, created_time from request_info order by created_time desc")
	checkErr(err)

	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		var id string
		var searchKey string
		var url string
		var requestType string
		var created time.Time
		err = rows.Scan(&id, &searchKey, &url, &requestType, &created)
		checkErr(err)
		result = append(result, RequestInfo{
			Id:          id,
			Url:         url,
			SearchKey:   searchKey,
			RequestType: requestType,
			CreatedTime: created,
		})
	}
	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
