package wallet

import "fmt"

// Bitcoin represents a number of Bitcoins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin someone owns
type Wallet struct {
	balance Bitcoin
}

// Withdraw will remove some bitcoin from a wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.balance -= amount
	return nil
}

// Deposit will add some Bitcoin to a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the number of Bitcoin a wallet has
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
