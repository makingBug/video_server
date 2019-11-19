package dbops

import (
	"testing"
)

func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate commnets")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M){
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
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