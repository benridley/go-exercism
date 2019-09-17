package account

import "sync"

// Account represents a bank account, with balance and active status
type Account struct {
	mu      sync.Mutex
	balance int
	active  bool
}

// New creates a bank account.
func New(amt int) *Account {
	if amt < 0 {
		return nil
	}

	acc := Account{
		balance: amt,
		active:  true,
	}
	return &acc
}

// Open is an alias for New
func Open(amt int) *Account {
	return New(amt)
}

// Balance returns the balance of a given bank account
func (acc *Account) Balance() (int, bool) {
	if !acc.active {
		return 0, false
	}
	return acc.balance, true
}

// Close deactivates a bank account and prevents further transactions.
func (acc *Account) Close() (int, bool) {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if !acc.active {
		return 0, false
	}
	acc.active = false
	return acc.balance, true
}

// Deposit adds funds to a bank account balance
func (acc *Account) Deposit(funds int) (int, bool) {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if !acc.active {
		return 0, false
	}

	if (acc.balance + funds) < 0 {
		return 0, false
	}

	acc.balance += funds
	return acc.balance, true
}

// Withdrawal reduces funds in a bank account
func (acc *Account) Withdrawal(funds int) (int, bool) {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	_, ok := acc.Deposit(-funds)
	if !ok {
		return acc.balance, false
	}
	return acc.balance, true
}
