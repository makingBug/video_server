package dbops

import (
	"fmt"
	"testing"
	"strconv"
	"time"
)
var tempvid string

func clearTables()  {
	// truncate user是mysql中的操作,清空users表
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M){
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
	clearTables()
	t.Run("add",testAddUser)
	t.Run("get",testGetUser)
	t.Run("delete",testDeleteUser)
	t.Run("reget",testRegetUser)
}

func testAddUser(t *testing.T)  {
	err := AddUserCredential("avenssi","123")
	if err!=nil{
		t.Errorf("Error of AddUser: %v",err)
	}
}

func testGetUser(t *testing.T)  {
	pwd,err := GetUserCredential("avenssi")
	if err !=nil || pwd != "123"{
		t.Errorf("Error of GetUser")
	}

}

func testDeleteUser(t *testing.T)  {
	err := DeleteUser("avenssi","123")
	if err!=nil{
		t.Errorf("Error of DeleteUser: %v",err)
	}
}

func testRegetUser(t *testing.T)  {
	pwd,err :=GetUserCredential("avenssi")
	if err!=nil{
		t.Errorf("Error of RegetUser: %v",err)
	}
	if(pwd!=""){
		t.Errorf("Deleting user test failed")
	}
}

func TestVideoWorkFlow(t *testing.T)  {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id//视频id是随机生成的,为了知道这个id多少,需要把它存下来
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T)  {
	clearTables()
	t.Run("AddUser",testAddUser)
	t.Run("AddComments",testAddComments)
	t.Run("ListComments",testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid,aid,content)
	if err != nil{
		t.Errorf("Error of AddComments: %v",err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to,_ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000,10))

	res,err := ListComments(vid,from,to)
	if err!= nil{
		t.Errorf("Error of ListComments: % v",err)
	}

	for i,ele := range res {
		fmt.Printf("comment: %d, %v \n",i,ele)
	}
}
