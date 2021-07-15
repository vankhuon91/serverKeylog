package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/kataras/iris/v12"
	_ "github.com/lib/pq"
)

type (
	request struct {
		MAC     string `json:"MAC"`
		IP      string `json:"IP"`
		Message string `json:"Message"`
	}
	RComputer struct {
		ID           int    `json:ID`
		MAC          string `json:"MAC"`
		IP           string `json:"IP"`
		ComputerName string `json:"ComputerName"`
		LastSeen     string `json:"LastSeen"`
		Created_ON   string `json:"Created_ON"`
	}
)

var db *sql.DB

func main() {
	connectionDB()
	app := iris.New()
	app.Handle("POST", "/api/computers", addNewComputer)
	var port_number = os.Getenv("PORT")
	if port_number == "" {
		port_number = "8080"
	}
	app.Listen(":" + port_number)
}

func addNewComputer(ctx iris.Context) {
	var req request
	ctx.ReadJSON(&req)
	timeNow := time.Now().Format("2006-01-02")
	keyMd5_byte := md5.Sum([]byte(req.MAC + timeNow))
	keyMd5_str := fmt.Sprintf("%x", keyMd5_byte)
	old_message := existsMD5(keyMd5_str)

	if old_message != "" {
		_, err := db.Exec(`update messages set message=$1 where md5=$2`, old_message+req.Message, keyMd5_str)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := db.Exec(`insert into messages (mac,message,date,md5) values ($1,$2,$3,$4)`, req.MAC, req.Message, timeNow, keyMd5_str)
		if err != nil {
			panic(err)
		}
	}
	ctx.JSON(req)
}

func existsMD5(MD5 string) string {
	var existsMD5 string
	row := db.QueryRow(`SELECT message from messages where md5=$1`, MD5)
	err := row.Scan(&existsMD5)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		} else {
			panic(err)
		}
	}
	return existsMD5
}

func connectionDB() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://ozlkvuxnytxeny:37cdd35ab7424cc55f5a788d40bb1a3f073b20f35f9eaa785c7cdf049aed3c01@ec2-52-23-45-36.compute-1.amazonaws.com:5432/d3ao1vte8slkbt"
	}
	db, _ = sql.Open("postgres", connStr)
}
