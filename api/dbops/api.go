package dbops

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func AddUserCredential(loginName string,pwd string) error  {
	stmtIns,err:=dbConn.Prepare( "INSERT INTO users(login_name,pwd) VALUES(?,?)")
	if err != nil{
		return err
	}
	//将这两个参数按顺序放入问号中
	stmtIns.Exec(loginName,pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)(string,error)  {
	stmtOut,err := dbConn.Prepare("SELECT pwd FROM users where login_name=?")
	if(err!= nil){
		log.Printf("%s",err)
		return "",err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string,pwd string) error{
	stmtDel,err := dbConn.Prepare("DELETE from users where login_name=? and pwd=?")
	if err!= nil{
		log.Printf("DeleteUser error: %s",err)
		return err
	}
	stmtDel.Exec(loginName,pwd)
	stmtDel.Close()
	return nil
}