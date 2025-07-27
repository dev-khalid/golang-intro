package advancedbankingsystem

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

type InterestBearing interface { 
	ApplyInterest()
}

type ManagedAccount interface {
	Account
	InterestBearing
}
