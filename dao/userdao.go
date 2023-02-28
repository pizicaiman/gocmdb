package dao

import (
	"cmdb/model"
	"cmdb/utils"
)

// 验证用户名和密码
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sqlstr := "select id, username, password from users where username=? and password=?"
	row := utils.Db.QueryRow(sqlstr, username, password)
	user := &model.User{}
	err := row.Scan(&user.Uid, &user.Username, &user.Password)
	//fmt.Printf("查询用户名和密码%s, %s\n", username, password)
	return user, err
}

// 验证用户名
func CheckUserName(username string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id, username, password from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.Uid, &user.Username, &user.Password)
	return user, nil
}

// 插入
func SaveUser(username string, password string, email string) {

	//sql语句
	sqlStr := "Insert into users(username,password,email) values (?,?,?)"

	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		panic(err.Error())
	}
}
