package pointer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWallet(t *testing.T) {

	asserts := require.New(t)

	t.Run("Deposit should increase the balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		asserts.Equal(Bitcoin(10), wallet.Balance())
	})

	t.Run("Withdraw should decrease the balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		asserts.Equal(Bitcoin(10), wallet.Balance())
	})

	t.Run("Withdraw should throw an error when insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(Bitcoin(100))
		asserts.Error(ErrInsufficientFunds, err)
		asserts.Equal(initialBalance, wallet.Balance())
	})
}
