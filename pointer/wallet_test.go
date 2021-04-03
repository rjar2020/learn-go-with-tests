package pointer

import (
	"fmt"
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
		err := wallet.Withdraw(Bitcoin(10))
		asserts.NoError(err)
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

func BenchmarkWallet_Balance(b *testing.B) {
	wallet := Wallet{balance: Bitcoin(20)}
	for i := 0; i < b.N; i++ {
		wallet.Balance()
	}
}

func BenchmarkWallet_Deposit(b *testing.B) {

	b.Run("new ballet", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wallet := Wallet{balance: Bitcoin(20)}
			wallet.Deposit(Bitcoin(10))
		}
	})

	b.Run("increasing funds same wallet", func(b *testing.B) {
		wallet := Wallet{balance: Bitcoin(20)}
		for i := 0; i < b.N; i++ {
			wallet.Deposit(Bitcoin(10))
		}
	})
}

func BenchmarkWallet_Withdraw(b *testing.B) {
	b.Run("new ballet", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wallet := Wallet{balance: Bitcoin(20)}
			_ = wallet.Withdraw(Bitcoin(10))
		}
	})

	b.Run("decreasing funds same wallet", func(b *testing.B) {
		wallet := Wallet{balance: Bitcoin(9223372036854775807)}
		for i := 0; i < b.N; i++ {
			_ = wallet.Withdraw(Bitcoin(100))
		}
	})

	b.Run("insufficient funds", func(b *testing.B) {
		wallet := Wallet{balance: Bitcoin(10)}
		for i := 0; i < b.N; i++ {
			_ = wallet.Withdraw(Bitcoin(100))
		}
	})
}

func ExampleWallet_Balance() {
	wallet := Wallet{balance: Bitcoin(20)}
	fmt.Printf("%s", wallet.Balance())
	/* Output: 20 BTC*/
}

func ExampleWallet_Deposit() {
	wallet := Wallet{balance: Bitcoin(20)}
	wallet.Deposit(Bitcoin(10))
	fmt.Printf("%s", wallet.Balance())
	/* Output: 30 BTC*/
}

func ExampleWallet_Withdraw() {
	wallet := Wallet{balance: Bitcoin(20)}
	err := wallet.Withdraw(Bitcoin(10))
	fmt.Printf("%s \"", wallet.Balance())
	err = wallet.Withdraw(Bitcoin(100))
	if err != nil {
		fmt.Printf("%s\" ", err)
	}
	fmt.Printf("%s", wallet.Balance())
	/* Output: 10 BTC "cannot withdraw, insufficient funds" 10 BTC*/
}
