package module1

import (
	"log"
	"testing"
	"time"
)

func TestRun1(t *testing.T) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 3)
		log.Println("TestRun1:::" + string(i))
	}
}

func TestRun2(t *testing.T) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 3)
		log.Println("TestRun2:::" + string(i))
	}
}
