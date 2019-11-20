package main

//import (
//	"net/http"
//	"github.com/julienschmidt/httprouter"
//)
//
//func RegisterHandlers() *httprouter.Router{
//	router := httprouter.New()
//	router.POST("/user",CreateUser)
//	router.POST("/user/:user_name",Login)
//	return router
//}
import (
	"log"
)
func main(){

	//r := RegisterHandlers()
	//http.ListenAndServe(":8000",r)

	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)

	log.Println("飞雪无情的博客:","http://www.flysnow.org")
	log.Printf("飞雪无情的微信公众号：%s\n","flysnow_org")
}
