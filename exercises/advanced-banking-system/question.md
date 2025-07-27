Let's level up. This next problem introduces **interface composition**, **pointer receivers**, **error handling**, and the use of a **type switch**, which are all common patterns in idiomatic Go.

### The Problem: Advanced Banking System 🏦

You are tasked with building the core logic for a banking system. The system needs to handle different types of accounts. All accounts must support basic transactions like deposits and withdrawals, but some special accounts also accrue interest.

Your goal is to create a system that can process a list of different account types, applying interest where applicable, and handling potential transaction errors.

---
### Requirements

1.  **Define a base `interface` named `Account`**. It must define methods for basic transactions:
    * `Deposit(amount float64)`
    * `Withdraw(amount float64) error` — This method should return an error if the withdrawal amount exceeds the balance.
    * `Balance() float64`

2.  **Define another `interface` named `InterestBearing`**.
    * `ApplyInterest()` — This method will calculate and add interest to the account's balance.

3.  **Compose the interfaces**. Define a new `interface` named `ManagedAccount` that includes all methods from *both* `Account` and `InterestBearing`.
    * *Hint: Go allows for embedding interfaces within other interfaces.*

4.  **Define two `struct` types**: `CheckingAccount` and `SavingsAccount`.
    * `CheckingAccount`: Should have an `Owner` (string) and a `CurrentBalance` (float64). It will only implement the `Account` interface.
    * `SavingsAccount`: Should have an `Owner` (string), a `CurrentBalance` (float64), and an `InterestRate` (float64, e.g., 0.02 for 2%). This struct must implement the `ManagedAccount` interface.

5.  **Implement the methods for the structs**.
    * **Important**: The methods that modify the balance (`Deposit`, `Withdraw`, `ApplyInterest`) **must use pointer receivers** (e.g., `func (a *SavingsAccount) Deposit(...)`) because they need to change the state of the original struct.
    * When implementing `Withdraw`, if the funds are insufficient, return an error (e.g., `fmt.Errorf("insufficient funds")`).
    * The `ApplyInterest` method for `SavingsAccount` should update the balance using the formula: `balance = balance * (1 + InterestRate)`.

6.  **Create a function named `ProcessTransactions`**.
    * This function should accept a slice of `Account` interfaces (`[]Account`).
    * Inside the function, iterate over the slice. For each `account` in the slice, use a **type switch** to determine its underlying concrete type.
        * If the account is a `InterestBearing`, cast it and call its `ApplyInterest()` method. You can print a message like "Applied interest to [Owner]'s savings account."
        * If the account is just a `CheckingAccount`, you can print a simple message like "Processing [Owner]'s checking account."
        * For all accounts, print their final balance after any processing.

7.  **In your `main` function**:
    * Create a slice of `Account`s.
    * Instantiate at least one `CheckingAccount` and one `SavingsAccount` and add their pointers to the slice.
    * Perform a successful `Deposit` and `Withdraw` on one of the accounts.
    * Attempt a `Withdraw` that will fail and correctly handle/print the returned error.
    * Finally, call `ProcessTransactions` with your slice of accounts to see the type switch in action.