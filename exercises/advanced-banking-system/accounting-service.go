package advancedbankingsystem

import (
	"fmt"
)

// CheckingAccount methods implementing Account interface
func (account *CheckingAccount) Deposit(amount float64) {
	account.CurrentBalance += amount
	fmt.Printf("Successfully deposited %f bucks.\n", amount)
}

func (account *CheckingAccount) Withdraw(amount float64) error {
	if account.CurrentBalance < amount {
		err := fmt.Errorf("insufficient funds")
		fmt.Println(err)
		return err
	}
	
	account.CurrentBalance -= amount
	fmt.Printf("Successful withdrawal of %f bucks.\n", amount)
	return nil
}

func (account *CheckingAccount) Balance() float64 {
	fmt.Printf("Available balance %f bucks.\n", account.CurrentBalance)
	return account.CurrentBalance
}

/** Note: We don't need the following code. Cause SavingsAccount already has CheckingAccount and that CheckingAccount already implemented Account interface. So SavingsAccount will inherit that from CheckingAccount.

func (account *SavingsAccount) Deposit(amount float64) {
	account.accountMutex.Lock()
	account.CurrentBalance += amount
	account.accountMutex.Unlock()

	fmt.Printf("Successfully deposited %f bucks.\n", amount)
}

func (account *SavingsAccount) Withdraw(amount float64) error {
	defer account.accountMutex.Unlock()

	account.accountMutex.Lock() 
	if(account.CurrentBalance < amount) { 
		err := fmt.Errorf("insufficient funds")
		fmt.Println(err)
		
		return err
	}
	
	account.CurrentBalance -= amount; 

	fmt.Printf("Successful withdrawal of %f bucks.\n", amount)
	return nil
}

func (account *SavingsAccount) Balance() float64 {
	fmt.Printf("Available balance %f bucks.\n", account.CurrentBalance)
	return account.CurrentBalance
}
*/

// SavingsAccount methods implementing ManagedAccount interface
func (account *SavingsAccount) ApplyInterest() { 
	account.accountMutex.Lock()
	account.CurrentBalance = account.CurrentBalance * (1 + account.InterestRate)
	account.accountMutex.Unlock()

	fmt.Printf("%.0f%% interest applied.\n", account.InterestRate*100)
}

func ProcessTransactions(accounts []Account) {
	for _, account := range accounts {

		// Check if the account implements InterestBearing interface
		if interestAccount, ok := account.(InterestBearing); ok {
			// Apply interest for accounts that support it
			interestAccount.ApplyInterest()
			savingsAcc, isSavings := account.(*SavingsAccount); if isSavings {
				fmt.Printf("Applied interest to %s's savings account.\n", savingsAcc.Owner)
			}
			
		} else {
			// For accounts that don't support interest (like CheckingAccount)
			if checkingAcc, isChecking := account.(*CheckingAccount); isChecking {
				fmt.Printf("Processing %s's checking account.\n", checkingAcc.Owner)
			}
		}
		
		// All Account interfaces have Balance method, so we can call it
		account.Balance()
	}
}