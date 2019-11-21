package main
//这个文件是整个流控机制
import (
	"log"
)

type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}



// buffer token算法完全不懂
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn:cc,
		bucket:make(chan int,cc),

	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn{
		log.Printf("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}


func (cl *ConnLimiter)ReleaseConn () {
	c := <- cl.bucket
	log.Printf("New connction coming: %d",c)
}
