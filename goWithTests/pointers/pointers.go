package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in deposit is %p \n", &wallet.balance)

	wallet.balance += amount
}

func(wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if wallet.balance < amount {
		return ErrInsufficientFunds
	}

	wallet.balance -= amount
	return nil
}