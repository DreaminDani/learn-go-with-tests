package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	verifyBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		verifyBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		verifyBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(9)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(10)

		verifyBalance(t, wallet, Bitcoin(startingBalance))

		if err == nil {
			t.Error("Wanted an error but didn't get one")
		}
	})
}
