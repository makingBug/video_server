package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"video_server/api/defs"
	"time"
	"video_server/api/utils"

	//"video_server/api/utils"
)


/***************************操作用户表**************************************/
func AddUserCredential(loginName string,pwd string) error  {
	stmtIns,err:=dbConn.Prepare("INSERT INTO users(login_name,pwd) VALUES(?,?)")
	if err != nil{
		return err
	}
	//将这两个参数按顺序放入问号中
	_,err=stmtIns.Exec(loginName,pwd)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)(string,error)  {
	stmtOut,err := dbConn.Prepare("SELECT pwd FROM users where login_name=?")
	if(err!= nil){
		log.Printf("%s",err)
		return "",err
	}
	var pwd string

	//单行查询,如果查询结果最多返回一行,可以使用QueryRow查询
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows{
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string,pwd string) error{
	stmtDel,err := dbConn.Prepare("DELETE from users where login_name=? and pwd=?")
	if err!= nil{
		log.Printf("DeleteUser error: %s",err)
		return err
	}
	_,err = stmtDel.Exec(loginName,pwd)
	if err != nil{
		return err
	}
	defer stmtDel.Close()
	return nil
}


/***************************操作视频表**************************************/
func AddNewVideo(aid int,name string)(*defs.VideoInfo,error){
	//create uuid
	vid,err := utils.NewUUID()
	if err != nil{
		return nil,err
	}
	t := time.Now()
	//时间字符不能更改,可以该格式,比如把:改成-,但是不能改数字,05不能改成06,据说是go诞生日期
	ctime := t.Format("Jan 02 2006,15:04:05")
	stmtIns,err := dbConn.Prepare("INSERT INTO video_info (id,author_id,name,display_ctime) values(?,?,?,?)")
	if err != nil{
		return nil,err
	}
	_,err = stmtIns.Exec(vid,aid,name,ctime)
	if err != nil{
		return nil,err
	}

	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayCtime:ctime}

	defer stmtIns.Close()
	return res,nil
}

func GetVideoInfo(vid string)(*defs.VideoInfo,error)  {
	stmtOut,err := dbConn.Prepare("select author_id,name,display_ctime from video_info where id=?")

	var aid int
	var dct string
	var name string

	//这是一个有疑问的地方,数据库中剩下的那个数create_time可以不获得,数据是按列传送的吗
	err = stmtOut.QueryRow(vid).Scan(&aid,&name,&dct)
	if err != nil && err != sql.ErrNoRows{
		return nil,err
	}

	if err == sql.ErrNoRows{
		return nil,nil
	}

	defer stmtOut.Close()
	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayCtime:dct}

	return res,nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

/***************************操作评论表**************************************/
func AddNewComments(vid string,aid int,content string) error {
	id,err := utils.NewUUID()
	if(err != nil){
		return err
	}

	stmtIns,err := dbConn.Prepare("insert into comments(id,video_id,author_id,content)values (?,?,?,?)")
	if err!=nil{
		return err
	}
	_,err = stmtIns.Exec(id,vid,aid,content)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string,from,to int)([] *defs.Comment,error){
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment
	rows,err := stmtOut.Query(vid,from,to)
	if err != nil{
		return res,err
	}

	for rows.Next(){
		var id,name,content string
		if err := rows.Scan(&id,&name,&content);err != nil{
			return res,err
		}

		c := &defs.Comment{Id:id,VideoId:vid,Author:name,Content:content}
		res = append(res,c)
	}

	defer stmtOut.Close()
	return res,nil
}