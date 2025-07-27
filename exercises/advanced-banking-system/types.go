package advancedbankingsystem

import "sync"

type CheckingAccount struct {
	Owner string
	CurrentBalance float64
}

type SavingsAccount struct { 
	CheckingAccount
	InterestRate float64
	accountMutex sync.Mutex
}
