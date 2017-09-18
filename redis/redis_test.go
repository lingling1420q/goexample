package redis

import (
	"fmt"
	logs "github.com/yangaowei/gologs"
	"testing"
)

func TestBase(t *testing.T) {
	conn, err := Dial("tcp", "127.0.0.1:6379")
	if err == nil {
		logs.Log.Debug("conn %v", conn)
		// result, e := conn.Do("auth", "waqu@test")
		// fmt.Println("result:", result)
		// fmt.Println("e:", e)
		result, e := conn.Do("keys", "*")
		fmt.Println("result:", result)
		fmt.Println("e:", e)

		conn.Close()
	} else {
		logs.Log.Debug("error %v", err)
	}
}

func dailPool() (Conn, error) {
	conn, err := Dial("tcp", "127.0.0.1:6379")
	return conn, err
}

func TestPool(t *testing.T) {
	pool := NewPool(dailPool, 3)
	fmt.Println("pool:", pool)

	conn, err := pool.Get()
	logs.Log.Debug("conn %v", conn)
	logs.Log.Debug("err %v", err)
	result, e := conn.Do("keys", "*")
	fmt.Println("result:", result)
	fmt.Println("e:", e)
	result, e = conn.Do("get", "fdsdaf")
	fmt.Println("result:", string(result.([]byte)))
	fmt.Println("e:", e)

	fmt.Println("action:", pool.ActiveCount())
	fmt.Println("IdleCount:", pool.IdleCount())
}
