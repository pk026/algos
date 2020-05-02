package main
import "fmt"

type Account struct {
	balance int
}

func (a *Account) Deposit(amount int) {
	a.balance = a.balance + amount
}

func (a *Account) Balance() {
	fmt.Printf("%v \n", a.balance)
}

func main() {
	account := Account{balance: 0}
	for i:= 0; i <= 2; i++  {
		go account.Deposit(10)
		account.Balance()
	}
}

/* here in above example 
Deposit method simply adds amount into balance
and Balance method simply retruns the balance

In normal scenarios if we access these method indivisually the will simply serve the purpose
But when we call in them concurrently
lets say Ram and Rahim trying to put some money into account and then checking balance
It is possible that Ram read from balance and did not finish updating balance by then Rahim read the balance too
and they read the same balance in account. When they update the balance they will it to their respective amount plus what they rad
and endup having wrong final value for balance. As it should be initial + Ram's amount + Rahim's amount
and endup updating to initial + Ram's amount/Rahim's amount
This is race
*/
