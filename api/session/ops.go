package session

import (
	"time"
	"sync"
	"video_server/api/defs"
	"video_server/api/dbops"
	"video_server/api/utils"
)


//这个线程安全的map需要好好研究一下
var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano()/1000000
}

func deleteExpireSession(sid string)  {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

//将session中数据从数据拿出来方在sessionMap中
func LoadSessionsFromDB(){
	r,err := dbops.RetrieveAllSessions()
	if err !=nil{
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss:= v.(*defs.SimpleSession)
		sessionMap.Store(k,ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id,_:= utils.NewUUID()
	ct := nowInMilli() //精度 ms
	ttl := ct + 30*60*1000//session在本地的过期时间,30min
	ss:= &defs.SimpleSession{Username:un,TTL:ttl}
	sessionMap.Store(id,ss)
	dbops.InserSession(id,ttl,un)
	return id
}

//判断session是否过期
func IsSessionExpired(sid string)(string,bool){
	ss,ok := sessionMap.Load(sid)
	if ok{
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct{
			//delete expired session
			deleteExpireSession(sid) //删除过期的session
			return "",true
		}
		return ss.(*defs.SimpleSession).Username,false
	}
	return "",true
}
