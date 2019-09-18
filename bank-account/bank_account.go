package account

import (
	"sync"
)

// Account represents a bank account, with balance and active status
type Account struct {
	mu              sync.Mutex
	balance         int
	active          bool
	closeRequests   chan closeRequest
	depositRequests chan depositRequest
}

type closeRequest struct {
	responder chan bankResponse
}

type depositRequest struct {
	responder chan bankResponse
	amt       int
}

type bankResponse struct {
	ok      bool
	balance int
}

// New creates a bank account.
func New(amt int) *Account {
	if amt < 0 {
		return nil
	}

	acc := Account{
		balance:         amt,
		active:          true,
		closeRequests:   make(chan closeRequest),
		depositRequests: make(chan depositRequest),
	}

	go func() {
		for {
			closeRequest := <-acc.closeRequests
			if acc.active {
				acc.active = false
				closeRequest.responder <- bankResponse{ok: true, balance: acc.balance}
			} else {
				closeRequest.responder <- bankResponse{ok: false, balance: 0}
			}
		}
	}()

	// Handle deposit requests
	go func() {
		for {
			depositReq := <-acc.depositRequests
			if !acc.active {
				depositReq.responder <- bankResponse{ok: false, balance: 0}
			} else if (acc.balance + depositReq.amt) < 0 {
				depositReq.responder <- bankResponse{ok: false, balance: 0}
			} else {
				acc.balance += depositReq.amt
				depositReq.responder <- bankResponse{ok: true, balance: acc.balance}
			}
		}
	}()

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
	responseChan := make(chan bankResponse)
	acc.closeRequests <- closeRequest{responder: responseChan}
	resp := <-responseChan
	return resp.balance, resp.ok
}

// Deposit adds funds to a bank account balance
func (acc *Account) Deposit(funds int) (int, bool) {
	responseChan := make(chan bankResponse)
	acc.depositRequests <- depositRequest{responder: responseChan, amt: funds}
	resp := <-responseChan
	return resp.balance, resp.ok
}

// Withdrawal reduces funds in a bank account
func (acc *Account) Withdrawal(funds int) (int, bool) {
	responseChan := make(chan bankResponse)
	acc.depositRequests <- depositRequest{responder: responseChan, amt: -funds}
	resp := <-responseChan
	return resp.balance, resp.ok
}
