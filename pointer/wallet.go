package pointer

import (
	"errors"
	"fmt"
)

//ErrInsufficientFunds is thronw when a withdraw can't be done in a wallet as there aren't enough funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Bitcoin helps to make the wallet implementation more specific
type Bitcoin int

//String implements Stringer for Bitcoin type
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Wallet models a digital wallet for storing bitcoins
type Wallet struct {
	balance Bitcoin
}

//Deposit allows a wallet to increase its balance by receiving funds
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance allows a client to query a wallet's balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Withdraw decrease a Wallets balance when a client takes some bitcoins
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
