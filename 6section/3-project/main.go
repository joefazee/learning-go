package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	acc.Balance += amount
	fmt.Printf("Deposited $%.2f to %s. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if acc.Balance < amount {
		return fmt.Errorf("insufficient funds in %s. Balance: $%.2f, Tried to withdraw: $%.2f",
			acc.AccountNumber, acc.Balance, amount)
	}
	acc.Balance -= amount
	fmt.Printf("Withdrew $%.2f from %s. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) GetBalance() float64 {
	return acc.Balance
}

func (acc *Account) String() string {
	return fmt.Sprintf("Account [%s] Owner: %s, Balance: $%.2f",
		acc.AccountNumber, acc.OwnerName, acc.Balance)
}

type SavingsAccount struct {
	Account              // Embed Account struct (anonymous field)
	InterestRate float64 // e.g., 0.02 for 2%
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate // Accesses promoted Balance field
	fmt.Printf("Adding interest $%.2f to savings account %s. ", interest, sa.AccountNumber)
	err := sa.Deposit(interest) // Uses promoted Deposit method
	if err != nil {
		fmt.Printf("AddInterst: Error depositing $%.2f to savings account. %v\n", interest, err)
	}
}

type OverdraftAccount struct {
	Account        // Embed Account struct
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	// Allow withdrawal up to Balance + OverdraftLimit
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("withdrawal of $%.2f exceeds overdraft limit for %s. Available including overdraft: $%.2f",
			amount, oa.AccountNumber, oa.Balance+oa.OverdraftLimit)
	}
	oa.Balance -= amount // Balance can go negative
	fmt.Printf("Withdrew $%.2f from overdraft account %s. New Balance: $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {

	fmt.Println("--- Bank Account System ---")

	savAcc := SavingsAccount{
		Account: Account{ // Initialize the embedded Account
			AccountNumber: "SAV001",
			Balance:       1000.00,
			OwnerName:     "Alice Saver",
		},
		InterestRate: 0.02, // 2%
	}
	fmt.Println("\n--- Savings Account Operations ---")
	fmt.Println(savAcc.Account.String())

	err := savAcc.Deposit(200.00)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %v\n", 200.00, err)
	}
	savAcc.AddInterest()
	err = savAcc.Withdraw(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Savings Details:", savAcc.Account.String())

	ovdAcc := OverdraftAccount{
		Account: Account{
			AccountNumber: "OVD002",
			Balance:       100.00,
			OwnerName:     "Bob Spender",
		},
		OverdraftLimit: 200.00,
	}

	fmt.Println("\n--- Overdraft Account Operations ---")
	fmt.Println(ovdAcc.Account.String())

	err = ovdAcc.Deposit(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(200.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error (expected for overdraft limit):", err)
	}

	fmt.Println("Final Overdraft Details:", ovdAcc.Account.String())
}
