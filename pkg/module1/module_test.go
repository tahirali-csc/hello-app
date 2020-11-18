package module1

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestRun1(t *testing.T) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 3)
		log.Println(fmt.Sprintf("TestRun1:::%d", i))
	}
}

func TestRun2(t *testing.T) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 3)
		log.Println(fmt.Sprintf("TestRun2:::%d", i))
	}
}
