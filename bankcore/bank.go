package bank

import (
	"errors"
	"fmt"
)

/*
1. 为客户和账户创建结构
先创建一个 Customer 结构，其中将包含要成为银行客户的人员的姓名、地址和电话号码。
此外，我们还需要 Account 数据的结构。 由于一个客户可以有多个账户，因此让我们将客户信息嵌入到账户对象中
*/

type Customer struct {
	Name    string // 姓名
	Address string // 地址
	Phone   string //电话号码
}

type Account struct {
	Customer         // 嵌套Customer结构体
	Number   int32   // 账号
	Balance  float64 //余额
}

/*
2. 实现存款方法
创建 Account 结构的 Deposit 方法。
如果收到的金额等于或小于零，该方法会返回一个错误。 否则，只需将收到的金额添加到账户的余额。
*/

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero\n存款金额应大于零")
	}
	a.Balance += amount
	return nil
}

/*
3. 实现取款方法
实现 Withdraw 方法的逻辑，该方法将账户余额减少的金额就是以参数方式收到的金额。
需要验证收到的数字是否大于零，
以及账户中的余额是否足够。
*/

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero\n提取的金额应大于零")
	}
	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance\n取款金额应大于账户余额")
	}
	a.Balance -= amount
	return nil
}

/*
4.实现对账单方法
我们将编写一种方法来输出对账单，其中包含账户名称、账号和余额。
*/

func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}
