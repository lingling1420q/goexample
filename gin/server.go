package main

import (
	logs "github.com/yangaowei/gologs"
)

func main() {
	logs.Log.Debug("test")
	logs.Log.Warn("test")
	logs.Log.Debug("test")
}
