package persist

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func init() {
	_, err := os.Stat("./data")
	if os.IsNotExist(err) {
		err := os.Mkdir("./data", 0777)
		if err != nil {
			return
		}
	}
	db, err := sql.Open("sqlite3", "./data/data.db")
	checkErr(err)
	_, err = db.Exec("create table if not exists request_info(id string primary key,search_key varchar, url varchar, request_type varchar, created_time date)")
	checkErr(err)
}
