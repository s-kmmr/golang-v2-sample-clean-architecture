package mysql

import (
	m "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	conn *gorm.DB
}

func (c *Client) Conn() *gorm.DB {
	return c.conn
}

func NewClient() *Client {
	// TODO configから設定できるようにする
	dsn := "docker:docker@tcp(127.0.0.1:3306)/pracv2?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(m.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db = db.Debug()

	return &Client{
		conn: db,
	}
}
