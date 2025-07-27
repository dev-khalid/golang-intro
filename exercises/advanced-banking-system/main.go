package advancedbankingsystem

import (
	"fmt"
)

func AdvancedBanking() {
	// Create a slice of Account interfaces
	var accounts []Account

	// Instantiate CheckingAccount and SavingsAccount
	checkingAccount := &CheckingAccount{
		Owner:          "Khalid",
		CurrentBalance: 1000.0,
	}

	savingsAccount := &SavingsAccount{
		CheckingAccount: CheckingAccount{
			Owner:          "Akash",
			CurrentBalance: 5000.0,
		},
		InterestRate: 0.25, // 25% interest rate
	}

	// Add their pointers to the slice
	accounts = append(accounts, checkingAccount)
	accounts = append(accounts, savingsAccount)

	// Perform a successful Deposit and Withdraw on one of the accounts
	fmt.Println("=== Performing successful transactions ===")
	checkingAccount.Deposit(200.0)
	err := checkingAccount.Withdraw(170.0)
	if err != nil {
		fmt.Printf("Withdrawal failed: %v\n", err)
	}

	// Attempt a Withdraw that will fail and handle the error
	fmt.Println("\n=== Attempting withdrawal that will fail ===")
	err = savingsAccount.Withdraw(10000.0) // This should fail - insufficient funds
	if err != nil {
		fmt.Printf("Withdrawal failed: %v\n", err)
	}

	// Call ProcessTransactions with the slice of accounts
	fmt.Println("\n=== Processing all transactions ===")
	ProcessTransactions(accounts)
}