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

func BenchmarkBalance(b *testing.B) {
	wallet := Wallet{balance: Bitcoin(20)}
	for i := 0; i < b.N; i++ {
		wallet.Balance()
	}
}

func BenchmarkDepositNewWallet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Deposit(Bitcoin(10))
	}
}

func BenchmarkDepositIncreasingFunds(b *testing.B) {
	wallet := Wallet{balance: Bitcoin(20)}
	for i := 0; i < b.N; i++ {
		wallet.Deposit(Bitcoin(10))
	}
}

func BenchmarkWithdrawNewWallet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			fmt.Printf("%s\" ", err)
		}
	}
}

func BenchmarkWithdrawDecreasingFunds(b *testing.B) {
	wallet := Wallet{balance: Bitcoin(9223372036854775807)}
	for i := 0; i < b.N; i++ {
		err := wallet.Withdraw(Bitcoin(100))
		if err != nil {
			fmt.Printf("%s\" ", err)
		}
	}
}

func BenchmarkWithdrawInsufficientFunds(b *testing.B) {
	wallet := Wallet{balance: Bitcoin(10)}
	for i := 0; i < b.N; i++ {
		err := wallet.Withdraw(Bitcoin(100))
		if err != nil {
			//Do nothing, as it will pollute benchmark's output
		}
	}
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
