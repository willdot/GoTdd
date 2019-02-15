package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is currency
type Bitcoin int

// Stringer is
type Stringer interface {
	String() string
}

// Wallet is a struct for a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit takes an amount and adds it to the wallet balance
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance returns the wallets balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds is an error for when the user tries to withdraw more than they are allowed
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw takes and amount and subtracts it from the balance
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// CalculateIntrest takes a percent and returns the interest
func (w *Wallet) CalculateIntrest(percent int) Bitcoin {

	var i = int(w.balance)

	return Bitcoin(i * percent)
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
