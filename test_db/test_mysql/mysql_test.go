package test_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type DbWorker struct {
	Dsn      string
	Db       *sql.DB
	UserInfo userTB
}
type userTB struct {
	Id   int
	Name sql.NullString
	Age  sql.NullInt64
}

func init() {

}

func (dbw *DbWorker) insertData() {
	stmt, _ := dbw.Db.Prepare(`INSERT INTO user (name, age) VALUES (?, ?)`)
	defer stmt.Close()

	ret, err := stmt.Exec("xys", 23)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

func (dbw *DbWorker) QueryDataPre() {
	dbw.UserInfo = userTB{}
}

func (dbw *DbWorker) queryData() {
	stmt, _ := dbw.Db.Prepare(`SELECT * From user where age >= ? AND age < ?`)
	defer stmt.Close()

	dbw.QueryDataPre()

	rows, err := stmt.Query(20, 30)
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&dbw.UserInfo.Id, &dbw.UserInfo.Name, &dbw.UserInfo.Age)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		if !dbw.UserInfo.Name.Valid {
			dbw.UserInfo.Name.String = ""
		}
		if !dbw.UserInfo.Age.Valid {
			dbw.UserInfo.Age.Int64 = 0
		}
		fmt.Println("get data, id: ", dbw.UserInfo.Id, " name: ", dbw.UserInfo.Name.String, " age: ", int(dbw.UserInfo.Age.Int64))
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func Test_001(t *testing.T) {
	var err error
	dbw := DbWorker{
		Dsn: "root:123456@tcp(113.102.163.179:4306)/testdb1?charset=utf8mb4",
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer dbw.Db.Close()

	dbw.insertData()
	dbw.queryData()
}

/*
// 1. 创建 testdb1 数据库
create database testdb1;

// 2. 创建 user 表
CREATE TABLE IF NOT EXISTS `user`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `name` VARCHAR(100) NOT NULL,
   `age` INT UNSIGNED NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
