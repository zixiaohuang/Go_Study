package main

//Question: 在操作数据库的时候，比如dao层中当遇到一个sql.ErrNoRows的时候，是否应该Wrap这个error，抛给上层。为什么，应该怎么做请写出代码？
// 个人觉得应该包装错误，方便上层区分哪个具体数据库相关（mysql/postgres/mongdb）,并定位到具体出现在哪个error

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// 连接时提供的用户信息
type UserInfo struct {
	account string // 访问数据库的账号
	passwd string // 访问数据库的密码
	ip_address string // ip地址
	ip_port string	// 端口
	select_db string // database
}

// table 信息
type TbInfo struct {
	uid int
	name string
	phone string
}

type MyDB struct {
	db *sql.DB
	e error
}

type dbMethod interface {
	initDB()
	queryRow()
	queryMultiRow()
}

func (dbCon *MyDB) initDB(userinfo UserInfo) error {
	// 初始化数据库链接
	var (
		err error
		connect_info string
	)
	connect_info = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userinfo.account,
		userinfo.passwd, userinfo.ip_address, userinfo.ip_port, userinfo.select_db)
	dbCon.db, err = sql.Open("mysql", connect_info)
	if err != nil {
		return  errors.Wrap(err, "connect db failed")
	}
	return nil
}

func (dbCon *MyDB) ReleaseDB() {
	// 释放数据库链接
	dbCon.db.Close()
}

func (dbCon *MyDB) queryRow(id int, tbinfo TbInfo) (TbInfo, error) {
	// 查询单行数据

	// 调用Scan，否则持有数据库的链接不会释放
	err := dbCon.db.QueryRow("select uid, name, phone from `user` where uid=? limit 1", id).Scan(&tbinfo.uid, &tbinfo.name, &tbinfo.phone)
	if err != nil {
		if errors.Is(err,sql.ErrNoRows) {
			return tbinfo, errors.Wrap(err,"query row failed")
		}
	}
	return tbinfo, nil
}

func (dbCon *MyDB)queryMultiRow(request string, tbinfo TbInfo) ([]string, error){
	// 访问多行数据
	// @request 需要执行的数据库的访问操作
	// @tbinfo 需要访问数据库表的信息
	var (
		single_res string
		result = make([]string, 10)
	)

	rows, err := dbCon.db.Query(request)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tbinfo.uid, &tbinfo.name, &tbinfo.phone)
		if err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		single_res = fmt.Sprintf("uid: %d name: %s phone: %s\n", tbinfo.uid, tbinfo.name, tbinfo.phone)
		result = append(result, single_res)
	}
	return result, nil
}

func main() {
	var (
		info UserInfo
		mydb MyDB
		request string
		tbinfo TbInfo
	)

	info.account = "user"
	info.passwd = "password"
	info.ip_address = "ip"
	info.ip_port = "port"
	info.select_db = "db name"

	err := mydb.initDB(info)
	if err != nil {
		fmt.Printf("init error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		return
	}

	// 再判一下防止db创建有误
	if mydb.db == nil {
		fmt.Printf("init error: db is nil")
		return
	}

	defer mydb.ReleaseDB()

	if mydb.db == nil {
		fmt.Printf("db nil\n")
	}

	request = "select uid, name, phone from `user`"
	res, err := mydb.queryMultiRow(request, tbinfo)
	if err != nil {
		fmt.Printf("query error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		return
	}

	fmt.Printf("multi query result:%v", res)

	tb_res, err := mydb.queryRow(6, tbinfo)
	if err != nil {
		fmt.Printf("query row error: %T %v\n", errors.Cause(err), errors.Cause(err))
		return
	}

	fmt.Printf("query row result:%+v", tb_res)
}
