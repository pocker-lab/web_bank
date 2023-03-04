package main

import (
	bank "web_bank/bankcore"
)

var accounts = map[float64]*bank.Account{}

func main() {
	// https://learn.microsoft.com/zh-cn/training/modules/go-write-test-program/4-project-api
	// https://learn.microsoft.com/zh-cn/training/modules/go-write-test-program/5-challenge
	_ = accounts
}
