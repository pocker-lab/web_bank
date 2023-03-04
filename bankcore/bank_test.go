package bank

import "testing"

// 测试帐户
func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("can't create an Account object\n无法创建帐户对象")
	}
}

// 测试存款
func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit\n存款后未更新余额")
	}
}

// 测试存款为负数
func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "渣男",
			Address: "云南省昆明市",
			Phone:   "13812348765",
		},
		Number:  1001,
		Balance: 0,
	}
	if err := account.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit\n只允许正数存款")
	}
}

// 测试取款
func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "渣男",
			Address: "云南省昆明市",
			Phone:   "13812348765",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(10)
	account.Withdraw(10)
	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw\n取款后未更新余额")
	}
}

// 测试对账单方法
func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("statement doesn't have the proper format\n语句的格式不正确")
	}
}
