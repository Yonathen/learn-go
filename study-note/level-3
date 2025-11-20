# Level 3: Multiple Packages - True Modularity

## 🎯 The Big Leap: From Organization to Architecture

Congratulations on mastering Levels 1 and 2! Level 3 is where Go really shines. You'll learn to create **truly independent modules** that can be **reused**, **tested separately**, and **shared with others**.

---

## Chapter 1: Why Multiple Packages? (The Real-World Problem)

### 1.1 The Limitations You've Hit

At Level 2, you probably experienced:
- **"I want to reuse this in another project"** - but it's all tangled together
- **"This is getting complex"** - everything can still talk to everything
- **"Different teams need different parts"** - but they have to take everything
- **"I want to distribute just this feature"** - but it's not independent

**This is the perfect time to learn packages!**

### 1.2 What Multiple Packages Solve

**🎯 True Independence:** Each package can exist on its own
**🎯 Reusability:** Import just what you need in other projects
**🎯 Clear Interfaces:** Explicit contracts between components
**🎯 Better Testing:** Test each package independently
**🎯 Team Collaboration:** Different teams own different packages
**🎯 Distribution:** Share packages individually

**Real-world analogy:** Moving from a single workshop to a proper factory with specialized departments.

### 1.3 The Mental Shift

**Level 2 thinking:** "How do I organize functions in files?"
**Level 3 thinking:** "How do I design independent, reusable components?"

This is a **fundamental shift** from organization to architecture!

---

## Chapter 2: Understanding Package Boundaries

### 2.1 What Makes a Good Package?

A good package is like a **well-designed tool** - it:
- **Does ONE thing really well** (single responsibility)
- **Has a clear interface** (exported functions)
- **Hides implementation details** (unexported internals)
- **Can work independently** (minimal dependencies)
- **Is easy to understand** (clear purpose)

### 2.2 The Package Design Process

**Step 1: Identify Independent Concepts**
Look for parts of your code that could exist independently:
- Email sending logic
- File processing utilities
- User authentication
- Data validation
- Mathematical calculations

**Step 2: Define Clear Interfaces**
What would other packages need from this package?
```go
// What should an email package provide?
func SendEmail(to, subject, body string) error
func ValidateEmail(email string) bool

// What should a math package provide?
func Add(a, b int) int
func Average(numbers []int) float64
```

**Step 3: Design Dependencies**
What does this package need from others (if anything)?

### 2.3 Package Responsibility Matrix

| Package Type | Responsibility | Example |
|--------------|----------------|---------|
| **Utility** | Helper functions | math, strings, validation |
| **Domain** | Business logic | user, order, product |
| **Infrastructure** | External concerns | database, email, files |
| **Interface** | User interaction | handlers, cli, gui |

---

## Chapter 3: Creating Your First Package

### 3.1 The Practical Setup

Let's create a real example - a **math utilities** package:

**Directory structure:**
```
myapp/
├── main.go
└── mathutil/
    └── calculator.go
```

### 3.2 Creating the mathutil Package

**File: mathutil/calculator.go**
```go
package mathutil

import "errors"

// Add performs addition of two integers
func Add(a, b int) int {
    return a + b
}

// Multiply performs multiplication
func Multiply(a, b int) int {
    return a * b
}

// Average calculates the average of a slice of numbers
func Average(numbers []int) (float64, error) {
    if len(numbers) == 0 {
        return 0, errors.New("cannot calculate average of empty slice")
    }
    
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return float64(sum) / float64(len(numbers)), nil
}

// isValid checks if a number is valid (unexported - private)
func isValid(num int) bool {
    return num >= 0
}

// ValidateNumbers checks if all numbers are valid  
func ValidateNumbers(numbers []int) bool {
    for _, num := range numbers {
        if !isValid(num) {
            return false
        }
    }
    return true
}
```

### 3.3 Using the Package

**File: main.go**
```go
package main

import (
    "fmt"
    "myapp/mathutil"
)

func main() {
    // Use exported functions from mathutil
    result := mathutil.Add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
    
    numbers := []int{1, 2, 3, 4, 5}
    avg, err := mathutil.Average(numbers)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Average: %.2f\n", avg)
    }
    
    // This would NOT work - isValid is unexported:
    // valid := mathutil.isValid(5) // Compile error!
    
    // But this works - ValidateNumbers is exported:
    valid := mathutil.ValidateNumbers(numbers)
    fmt.Printf("Numbers are valid: %t\n", valid)
}
```

### 3.4 Key Concepts in Action

**✅ Package Declaration:** `package mathutil` creates a new namespace
**✅ Exported vs Unexported:** `Add()` vs `isValid()`
**✅ Import Path:** `import "myapp/mathutil"`
**✅ Qualified Calls:** `mathutil.Add()`
**✅ Error Handling:** Packages can return errors
**✅ Interface Design:** Clean, focused API

---

## Chapter 4: Building a Multi-Package Application

Let's build a complete application with multiple packages to see real architecture:

### 4.1 Application Design: Personal Finance Tracker

**Packages we'll create:**
- `main` - Entry point and coordination
- `account` - Account management
- `transaction` - Transaction processing  
- `report` - Reporting and analysis
- `fileutil` - File operations

**Directory structure:**
```
finance-tracker/
├── main.go
├── account/
│   └── account.go
├── transaction/
│   └── transaction.go
├── report/
│   └── report.go
└── fileutil/
    └── csv.go
```

### 4.2 The account Package

**File: account/account.go**
```go
package account

import (
    "errors"
    "fmt"
)

// Account represents a financial account
type Account struct {
    ID      int
    Name    string
    Balance float64
}

// accounts stores all accounts (unexported - internal state)
var accounts []Account
var nextID int = 1

// CreateAccount creates a new account
func CreateAccount(name string) *Account {
    account := Account{
        ID:      nextID,
        Name:    name,
        Balance: 0.0,
    }
    accounts = append(accounts, account)
    nextID++
    return &account
}

// GetAccount retrieves an account by ID
func GetAccount(id int) (*Account, error) {
    for i := range accounts {
        if accounts[i].ID == id {
            return &accounts[i], nil
        }
    }
    return nil, errors.New("account not found")
}

// GetAllAccounts returns all accounts
func GetAllAccounts() []Account {
    return accounts
}

// UpdateBalance updates an account's balance
func UpdateBalance(id int, amount float64) error {
    account, err := GetAccount(id)
    if err != nil {
        return err
    }
    account.Balance += amount
    return nil
}

// String provides a string representation of an account
func (a Account) String() string {
    return fmt.Sprintf("Account %d: %s (Balance: $%.2f)", a.ID, a.Name, a.Balance)
}
```

### 4.3 The transaction Package

**File: transaction/transaction.go**
```go
package transaction

import (
    "errors"
    "time"
)

// Transaction represents a financial transaction
type Transaction struct {
    ID        int
    AccountID int
    Amount    float64
    Category  string
    Date      time.Time
    Note      string
}

// transactions stores all transactions (unexported - internal state)
var transactions []Transaction
var nextID int = 1

// CreateTransaction creates a new transaction
func CreateTransaction(accountID int, amount float64, category, note string) *Transaction {
    transaction := Transaction{
        ID:        nextID,
        AccountID: accountID,
        Amount:    amount,
        Category:  category,
        Date:      time.Now(),
        Note:      note,
    }
    transactions = append(transactions, transaction)
    nextID++
    return &transaction
}

// GetTransactionsByAccount returns all transactions for an account
func GetTransactionsByAccount(accountID int) []Transaction {
    var accountTransactions []Transaction
    for _, t := range transactions {
        if t.AccountID == accountID {
            accountTransactions = append(accountTransactions, t)
        }
    }
    return accountTransactions
}

// GetTransactionsByCategory returns all transactions in a category
func GetTransactionsByCategory(category string) []Transaction {
    var categoryTransactions []Transaction
    for _, t := range transactions {
        if t.Category == category {
            categoryTransactions = append(categoryTransactions, t)
        }
    }
    return categoryTransactions
}

// GetAllTransactions returns all transactions
func GetAllTransactions() []Transaction {
    return transactions
}

// calculateTotal calculates total for a slice of transactions (unexported helper)
func calculateTotal(transactions []Transaction) float64 {
    total := 0.0
    for _, t := range transactions {
        total += t.Amount
    }
    return total
}

// GetTotalForAccount calculates total amount for an account
func GetTotalForAccount(accountID int) float64 {
    accountTransactions := GetTransactionsByAccount(accountID)
    return calculateTotal(accountTransactions)
}
```

### 4.4 The report Package

**File: report/report.go**
```go
package report

import (
    "fmt"
    "finance-tracker/account"
    "finance-tracker/transaction"
)

// GenerateAccountSummary creates a summary report for all accounts
func GenerateAccountSummary() {
    fmt.Println("=== ACCOUNT SUMMARY ===")
    accounts := account.GetAllAccounts()
    
    for _, acc := range accounts {
        fmt.Printf("%s\n", acc.String())
        
        transactions := transaction.GetTransactionsByAccount(acc.ID)
        fmt.Printf("  Transactions: %d\n", len(transactions))
        
        total := transaction.GetTotalForAccount(acc.ID)
        fmt.Printf("  Total Activity: $%.2f\n", total)
        fmt.Println()
    }
}

// GenerateCategoryReport creates a report by category
func GenerateCategoryReport() {
    fmt.Println("=== CATEGORY REPORT ===")
    
    // Get all unique categories
    categories := getUniqueCategories()
    
    for _, category := range categories {
        transactions := transaction.GetTransactionsByCategory(category)
        total := 0.0
        for _, t := range transactions {
            total += t.Amount
        }
        
        fmt.Printf("%s: %d transactions, $%.2f total\n", 
                   category, len(transactions), total)
    }
}

// getUniqueCategories extracts unique categories (unexported helper)
func getUniqueCategories() []string {
    categoryMap := make(map[string]bool)
    transactions := transaction.GetAllTransactions()
    
    for _, t := range transactions {
        categoryMap[t.Category] = true
    }
    
    var categories []string
    for category := range categoryMap {
        categories = append(categories, category)
    }
    return categories
}
```

### 4.5 The main Package (Orchestration)

**File: main.go**
```go
package main

import (
    "fmt"
    "finance-tracker/account"
    "finance-tracker/transaction" 
    "finance-tracker/report"
)

func main() {
    fmt.Println("Personal Finance Tracker")
    fmt.Println("========================")
    
    // Create accounts using account package
    checking := account.CreateAccount("Checking Account")
    savings := account.CreateAccount("Savings Account")
    
    // Create transactions using transaction package
    transaction.CreateTransaction(checking.ID, -50.0, "Food", "Grocery shopping")
    transaction.CreateTransaction(checking.ID, 1000.0, "Salary", "Monthly salary")
    transaction.CreateTransaction(savings.ID, 500.0, "Transfer", "Monthly savings")
    transaction.CreateTransaction(checking.ID, -30.0, "Transport", "Gas")
    
    // Update account balances based on transactions
    account.UpdateBalance(checking.ID, transaction.GetTotalForAccount(checking.ID))
    account.UpdateBalance(savings.ID, transaction.GetTotalForAccount(savings.ID))
    
    // Generate reports using report package
    report.GenerateAccountSummary()
    report.GenerateCategoryReport()
}
```

### 4.6 What We Achieved

**✅ Clear Separation:** Each package has a distinct responsibility
**✅ Reusable Components:** Any package can be used in other projects
**✅ Clean Interfaces:** Each package exposes only what others need
**✅ Testable:** Each package can be tested independently
**✅ Maintainable:** Changes in one package don't affect others (much)
**✅ Understandable:** Anyone can see what each package does

---

## Chapter 5: Package Design Principles

### 5.1 The Single Responsibility Principle

Each package should have ONE reason to change:

**Good:**
```go
// email package - only changes when email logic changes
package email
func SendEmail(to, subject, body string) error
func ValidateEmail(email string) bool
```

**Bad:**
```go
// utils package - changes for ANY reason
package utils
func SendEmail() error      // Email changes affect this
func ValidateInput() bool   // Validation changes affect this  
func CalculateTotal() int   // Math changes affect this
```

### 5.2 The Dependency Inversion Principle

High-level packages should not depend on low-level packages directly:

**Good Design:**
```
main → handlers → models
main → database
(main coordinates, handlers and database don't depend on each other)
```

**Bad Design:**
```
handlers → database
(handlers are tightly coupled to database)
```

### 5.3 The Interface Segregation Principle

Packages should expose small, focused interfaces:

**Good:**
```go
// user package
func CreateUser(name string) User
func ValidateUser(user User) error

// auth package  
func Login(username, password string) bool
func Logout(user User)
```

**Bad:**
```go
// userauth package (doing too much)
func CreateUser() User
func ValidateUser() error
func Login() bool
func Logout() 
func SendEmail()    // Why is this here?
func HashPassword() // This should be internal
```

---

## Chapter 6: Advanced Package Patterns

### 6.1 The Repository Pattern

Create packages that abstract data storage:

```go
// File: storage/user.go
package storage

type User struct {
    ID   int
    Name string
}

// UserRepository defines how to store/retrieve users
type UserRepository interface {
    Save(user User) error
    GetByID(id int) (User, error)
    GetAll() ([]User, error)
}

// MemoryUserRepository implements UserRepository using memory
type MemoryUserRepository struct {
    users []User
}

func (m *MemoryUserRepository) Save(user User) error {
    m.users = append(m.users, user)
    return nil
}

func (m *MemoryUserRepository) GetByID(id int) (User, error) {
    for _, user := range m.users {
        if user.ID == id {
            return user, nil
        }
    }
    return User{}, errors.New("user not found")
}
```

### 6.2 The Service Pattern

Create packages for business logic:

```go
// File: service/user.go
package service

import "myapp/storage"

type UserService struct {
    repo storage.UserRepository
}

func NewUserService(repo storage.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name string) error {
    // Business logic here
    if name == "" {
        return errors.New("name cannot be empty")
    }
    
    user := storage.User{
        ID:   generateID(), // Some ID generation logic
        Name: name,
    }
    
    return s.repo.Save(user)
}
```

### 6.3 The Factory Pattern

Create packages that construct complex objects:

```go
// File: factory/database.go
package factory

import "myapp/storage"

type DatabaseType string

const (
    Memory   DatabaseType = "memory"
    File     DatabaseType = "file"
    Postgres DatabaseType = "postgres"
)

func CreateUserRepository(dbType DatabaseType) storage.UserRepository {
    switch dbType {
    case Memory:
        return &storage.MemoryUserRepository{}
    case File:
        return &storage.FileUserRepository{}
    case Postgres:
        return &storage.PostgresUserRepository{}
    default:
        return &storage.MemoryUserRepository{}
    }
}
```

---

## Chapter 7: Testing Multiple Packages

### 7.1 Package-Level Testing

Each package should have its own tests:

```go
// File: mathutil/calculator_test.go
package mathutil

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

func TestAverage(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}
    result, err := Average(numbers)
    
    if err != nil {
        t.Errorf("Average(%v) returned error: %v", numbers, err)
    }
    
    expected := 3.0
    if result != expected {
        t.Errorf("Average(%v) = %f; want %f", numbers, result, expected)
    }
}

func TestAverageEmptySlice(t *testing.T) {
    _, err := Average([]int{})
    if err == nil {
        t.Error("Average([]) should return an error")
    }
}
```

### 7.2 Integration Testing

Test how packages work together:

```go
// File: integration_test.go
package main

import (
    "testing"
    "myapp/account"
    "myapp/transaction"
)

func TestAccountTransactionIntegration(t *testing.T) {
    // Create account
    acc := account.CreateAccount("Test Account")
    
    // Create transaction
    transaction.CreateTransaction(acc.ID, 100.0, "Test", "Test transaction")
    
    // Update balance
    err := account.UpdateBalance(acc.ID, 100.0)
    if err != nil {
        t.Errorf("Failed to update balance: %v", err)
    }
    
    // Verify balance
    updatedAcc, err := account.GetAccount(acc.ID)
    if err != nil {
        t.Errorf("Failed to get account: %v", err)
    }
    
    if updatedAcc.Balance != 100.0 {
        t.Errorf("Expected balance 100.0, got %f", updatedAcc.Balance)
    }
}
```

---

## Chapter 8: Package Documentation

### 8.1 Package-Level Documentation

Document what your package does:

```go
// Package mathutil provides mathematical utility functions for common
// calculations including basic arithmetic, averages, and validation.
//
// This package is designed to be simple and focused, providing only
// the most commonly needed mathematical operations with proper error
// handling.
//
// Example usage:
//
//     result := mathutil.Add(5, 3)
//     avg, err := mathutil.Average([]int{1, 2, 3, 4, 5})
//     if err != nil {
//         log.Fatal(err)
//     }
//
package mathutil
```

### 8.2 Function Documentation

Document exported functions clearly:

```go
// Add performs addition of two integers and returns the result.
// This function handles integer overflow according to Go's standard
// integer overflow behavior.
func Add(a, b int) int {
    return a + b
}

// Average calculates the arithmetic mean of a slice of integers.
// It returns an error if the slice is empty.
//
// The result is returned as a float64 to preserve precision.
//
// Example:
//     numbers := []int{1, 2, 3, 4, 5}
//     avg, err := Average(numbers)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Printf("Average: %.2f\n", avg) // Output: Average: 3.00
func Average(numbers []int) (float64, error) {
    // implementation...
}
```

### 8.3 Generating Documentation

Go can generate beautiful documentation:

```bash
# Generate documentation for your packages
go doc myapp/mathutil

# Start a documentation server
godoc -http=:6060
# Then visit http://localhost:6060/pkg/myapp/
```

---

## Chapter 9: Common Pitfalls and Solutions

### 9.1 Circular Dependencies

**Problem:**
```go
// package a imports package b
// package b imports package a
```

**Solution:**
Create a third package for shared types or redesign the dependencies.

### 9.2 Overly Large Packages

**Problem:**
Packages that do too many things.

**Solution:**
Split into smaller, more focused packages.

### 9.3 Too Many Small Packages

**Problem:**
Packages with only one or two functions.

**Solution:**
Combine related small packages into larger, cohesive ones.

### 9.4 Unclear Package Boundaries

**Problem:**
Not knowing what belongs in which package.

**Solution:**
Define clear responsibilities and stick to them.

---

## Chapter 10: Mastery and Next Steps

### 10.1 Mastery Checklist

- [ ] Can design packages with single responsibilities
- [ ] Understand exported vs unexported clearly
- [ ] Can create clean package interfaces
- [ ] Know how to avoid circular dependencies
- [ ] Can test packages independently
- [ ] Understand package documentation
- [ ] Can refactor monolithic code into packages
- [ ] Think in terms of reusable components

### 10.2 Advanced Topics to Explore

**Interfaces:** Define contracts between packages
**Dependency Injection:** Make packages more flexible
**Package Vendoring:** Managing external dependencies
**Module System:** Modern Go dependency management
**Design Patterns:** Common architectural patterns in Go

### 10.3 Real-World Practice

**Build These Projects:**
1. **CLI Tool with Packages** - Each command as a package
2. **Web API with Packages** - handlers, models, database as separate packages
3. **Library for Others** - Create a package others can import
4. **Microservice** - Multiple independent packages

---

## 🎯 Key Takeaways

1. **Packages enable true modularity** - independent, reusable components
2. **Design interfaces first** - think about what each package should provide
3. **Keep dependencies flowing one direction** - avoid circular dependencies
4. **Single responsibility per package** - one reason to change
5. **Test each package independently** - packages should be self-contained
6. **Document your packages** - others (including future you) will thank you

**The Level 3 achievement:** You now think in terms of **architecture** and **design**, not just organization! 🎉

---

## What's Next?

You've mastered the fundamentals! Next you can explore:
- **Advanced Go patterns** (interfaces, embedding, concurrency)
- **Web frameworks** (using package-based architecture)
- **Database integration** (with proper package separation)
- **Testing strategies** (unit, integration, e2e)
- **Deployment and distribution** (sharing your packages)

Congratulations on reaching Level 3! You're now thinking like a professional Go developer! 🚀
