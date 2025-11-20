# Learning Go: From Basic to Advanced - REST API Journey

## 🧠 **Learning-Focused Approach**

### **What Changed:**
- ❌ Removed overwhelming code examples
- ✅ Added **WHY** explanations for every concept
- ✅ Progressive learning from basic to advanced
- ✅ Real-world analogies (restaurant, library)
- ✅ Common mistakes and why they happen
- ✅ Mental models for thinking about packages

### **Structure of Learning:**

#### **Chapter 1-2: Foundation**
- Why packages exist (organization problem)
- Basic rules and reasoning behind them
- Import system explanation

#### **Chapter 3-4: Core Concepts** 
- Public vs Private (and why Go chose capitalization)
- Package design principles
- Dependency direction

#### **Chapter 5-6: Architecture**
- Why web apps use layered structure
- Mental models (restaurant/library analogies)
- How to think about package responsibilities

#### **Chapter 7-8: Practical Wisdom**
- Learning progression (5 levels)
- Common beginner mistakes
- Why those mistakes happen

#### **Chapter 9-10: Advanced Path**
- Practical exercises
- Advanced patterns when you're ready

### **Key Learning Principles:**
1. **Start with WHY** - Every concept explained with reasoning
2. **Use analogies** - Complex ideas made simple
3. **Progressive complexity** - Build understanding step by step
4. **Learn from mistakes** - Common pitfalls and solutions
5. **Practical focus** - How to apply knowledge

### **For You as a Beginner:**
- **No overwhelming code dumps**
- **Clear explanations** of why things work the way they do
- **Learning path** from where you are now to advanced
- **Practical exercises** to build understanding
- **Mental models** to think like a Go developer

## Chapter 1: Understanding the Problem (Why Do We Need Packages?)

### The Single File Problem
When you first learned Go, everything was in one file:
```go
package main
func main() { ... }
```

**This works for tiny programs, but imagine:**
- 50 functions in one file
- User logic mixed with database logic
- Hard to find anything
- Hard to test individual parts
- Impossible to reuse code

**Real-world analogy:** It's like putting your entire house (kitchen, bedroom, bathroom) in one room!

### The Solution: Organization
**Packages** are like rooms in a house - each has a specific purpose:
- `kitchen/` - cooking functions
- `bedroom/` - sleeping functions  
- `bathroom/` - cleaning functions

In programming:
- `models/` - data structures
- `handlers/` - web request processing
- `database/` - data storage operations

---

## Chapter 2: Basic Package Concepts (The Foundation)

### 2.1 What is a Package?
A **package** is simply a folder with related Go files that work together.

**Key Rules:**
1. All `.go` files in the same folder = same package
2. Package name should match folder name (except `main`)
3. `main` package = executable program

### 2.2 Why These Rules Exist

**Why same folder = same package?**
- Go needs to know which files belong together
- Makes it easy to organize related functionality
- Compiler can optimize files together

**Why package name should match folder?**
- Makes imports predictable: `import "myapp/users"` loads the `users/` folder
- Other developers know what they're getting

**Why `main` package is special?**
- Go needs to know where your program starts
- `main()` function is the entry point
- Only `main` packages become executable programs

### 2.3 The Import System (How Packages Talk)

```go
import "fmt"          // Standard library package
import "myapp/users"  // Your own package
```

**What really happens:**
1. Go looks for the package
2. Loads all `.go` files in that package
3. Makes exported items available
4. Creates a namespace to avoid conflicts

---

## Chapter 3: Public vs Private (Access Control)

### 3.1 The Capitalization Rule
Go uses a simple rule for access control:
- **Uppercase first letter** = Public (exported)
- **Lowercase first letter** = Private (unexported)

### 3.2 Why This System?

**Traditional languages use keywords:**
```java
public class User { ... }
private void validate() { ... }
```

**Go uses capitalization because:**
- Simpler - no keywords to remember
- Visual - you can see it immediately
- Consistent - same rule for everything

### 3.3 Real-World Example

```go
// In users package
func CreateUser(name string) User {    // PUBLIC - other packages can use
    user := User{Name: name}
    if validate(user) {                // PRIVATE - only this package can call
        return user
    }
}

func validate(u User) bool {           // PRIVATE - implementation detail
    return u.Name != ""
}
```

**Why this matters:**
- `CreateUser` is the **interface** - what other packages need
- `validate` is **implementation** - internal logic that might change
- Other packages depend on the interface, not the implementation

---

## Chapter 4: Package Design Principles (Thinking Like a Pro)

### 4.1 Single Responsibility Principle
Each package should have ONE main job.

**Bad Design:**
```
utils/
├── math.go        // Math functions
├── database.go    // Database stuff  
├── email.go       // Email sending
└── files.go       // File operations
```
**Problem:** "utils" tells you nothing. What does this package do?

**Good Design:**
```
math/
├── calculator.go
└── geometry.go

email/
├── sender.go
└── templates.go

storage/
├── files.go
└── database.go
```
**Why better:** Each package has a clear, focused purpose.

### 4.2 Dependency Direction

**Rule:** High-level packages import low-level packages, not the other way around.

**Good:**
```
main.go imports handlers
handlers imports models
models imports nothing (or just standard library)
```

**Bad:**
```
models imports handlers (circular dependency!)
```

**Why this matters:**
- Creates a clear hierarchy
- Prevents circular dependencies
- Makes testing easier
- Makes code more modular

### 4.3 Package Interfaces

Think of packages as having **interfaces** (what they provide) and **implementations** (how they do it).

**Example:**
```go
// Package interface (what it promises)
func CreateUser(name, email string) (*User, error)
func GetUser(id int) (*User, error)
func DeleteUser(id int) error

// Implementation (how it works inside - can change!)
// - Database queries
// - Validation logic
// - Error handling
```

---

## Chapter 5: Web Application Architecture (Why This Structure?)

### 5.1 The Layered Approach

```
main.go          ← Entry point
handlers/        ← Web layer (HTTP requests/responses)
models/          ← Data layer (structures, validation)
database/        ← Storage layer (database operations)
utils/           ← Helper layer (common functions)
```

**Why this structure?**

**Separation of Concerns:** Each layer has one job:
- Handlers: "How do we process web requests?"
- Models: "What data do we work with?"
- Database: "How do we store/retrieve data?"
- Utils: "What common operations do we need?"

**Independence:** You can change one layer without breaking others:
- Switch from PostgreSQL to MySQL? Only change `database/`
- Change API format? Only change `handlers/`
- Add new validation? Only change `models/`

### 5.2 Why Not Put Everything in Main?

**Beginner thinking:** "Why not just put all functions in `main.go`?"

**Problems that emerge:**
1. **File becomes huge** (1000+ lines)
2. **Hard to test** (can't test individual functions easily)
3. **Hard to reuse** (can't import parts in other projects)
4. **Hard to collaborate** (multiple people editing same file)
5. **Hard to understand** (mixing web logic with database logic)

**Professional approach:** Separate concerns into packages.

---

## Chapter 6: Building Mental Models (How to Think About Packages)

### 6.1 The Restaurant Analogy

Think of a restaurant:
- **Kitchen** (models/) - prepares the food (data)
- **Waiters** (handlers/) - take orders and serve customers (handle requests)
- **Storage** (database/) - stores ingredients (data persistence)
- **Utilities** (utils/) - shared tools like knives, cleaning (common functions)

Each area has:
- **Specific responsibilities**
- **Clear interfaces** (menu, order forms)
- **Internal operations** customers don't see

### 6.2 The Library Analogy

A library has:
- **Public areas** (reading rooms) - anyone can use
- **Private areas** (staff offices) - only staff can access
- **Catalog system** (imports) - how you find books
- **Sections** (packages) - fiction, science, history

Same with Go packages:
- **Exported functions** = public areas
- **Unexported functions** = private areas
- **Import statements** = catalog system
- **Packages** = sections

---

## Chapter 7: Progressive Complexity (Learning Path)

### 7.1 Level 1: Single Package (Where You Are)
```go
package main
func main() { ... }
func helper() { ... }
```
**Learn:** Basic Go syntax, functions, variables

### 7.2 Level 2: Multiple Files, Same Package
```
main.go
helpers.go
```
**Learn:** How files in same package share functions

### 7.3 Level 3: Multiple Packages
```
main.go
utils/
└── math.go
```
**Learn:** Import system, exported/unexported

### 7.4 Level 4: Structured Application
```
main.go
models/
handlers/
database/
utils/
```
**Learn:** Architecture, dependency management

### 7.5 Level 5: Advanced Patterns
- Interfaces
- Dependency injection
- Testing strategies
- Error handling patterns

---

## Chapter 8: Common Beginner Mistakes (And Why They Happen)

### 8.1 Mistake: Making Everything Public
```go
// Beginner code
func CreateUser() { ... }  // OK - needed by other packages
func ValidateEmail() { ... } // BAD - should be private
func HashPassword() { ... }  // BAD - should be private
```

**Why this happens:** Fear that they might need it later.

**Better approach:** Start private, make public only when needed.

### 8.2 Mistake: Circular Imports
```go
// models/user.go imports handlers/auth.go
// handlers/auth.go imports models/user.go
```

**Why this happens:** Not understanding dependency direction.

**Solution:** Dependencies should flow in one direction (down the hierarchy).

### 8.3 Mistake: God Packages
```go
// Everything in utils/
utils/
├── everything.go  // 500 lines of random functions
```

**Why this happens:** Not knowing where to put things.

**Solution:** Create specific packages for specific purposes.

---

## Chapter 9: Practical Exercise Progression

### 9.1 Exercise 1: Convert Single File to Packages
**Start with:** Everything in `main.go`
**Goal:** Split into logical packages
**Learn:** Package organization, imports

### 9.2 Exercise 2: Add a New Feature
**Task:** Add user authentication
**Goal:** Understand which package new code belongs in
**Learn:** Package responsibilities

### 9.3 Exercise 3: Refactor for Testing
**Task:** Make code testable
**Goal:** Design clean interfaces between packages
**Learn:** Dependency management, interfaces

---

## Chapter 10: Advanced Concepts (When You're Ready)

### 10.1 Interfaces Between Packages
Instead of depending on concrete types, depend on interfaces:
```go
// Good: handler depends on interface
type UserStorage interface {
    CreateUser(user User) error
    GetUser(id int) (*User, error)
}

// Bad: handler depends on concrete database package
```

### 10.2 Dependency Injection
Pass dependencies into packages instead of hardcoding them:
```go
// Good: flexible
func NewHandler(storage UserStorage) *Handler

// Bad: rigid
func NewHandler() *Handler {
    storage := database.NewPostgresStorage() // hardcoded!
}
```

### 10.3 Package-Level Design Patterns
- Repository pattern (for data access)
- Service pattern (for business logic)
- Factory pattern (for object creation)

---

## 🎯 Key Takeaways

1. **Packages are about organization** - like rooms in a house
2. **Start simple, add complexity gradually** - don't over-engineer
3. **Think about dependencies** - who needs what from whom?
4. **Public vs private is about interfaces** - what do others need to see?
5. **Each package should have a clear purpose** - single responsibility
6. **Architecture emerges from organization** - good packages lead to good architecture

Remember: You don't need to understand everything at once. Start with basic package organization and build understanding gradually!

---

## Next Steps for You

1. **Experiment with your current hello.go/variable.go** - try splitting into packages
2. **Read other Go projects** - see how they organize packages
3. **Start small projects** - practice package organization
4. **Don't fear refactoring** - packages can be reorganized as you learn

The best way to learn is by doing! 🚀

---

## Step 1: Initialize the Project

### 1.1 Create Project Directory
```bash
mkdir webapp
cd webapp
```

### 1.2 Initialize Go Module
```bash
go mod init webapp
```

### 1.3 Create Directory Structure
```bash
mkdir handlers models database utils
```

---

## Step 2: Create the Models Package

The models package defines our data structures.

### 2.1 Create `models/user.go`
```go
package models

import (
    "time"
    "errors"
)

// User represents a user in our system
type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"` // Don't include in JSON responses
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// UserRequest represents user input for registration/updates
type UserRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Validate validates user input
func (ur *UserRequest) Validate() error {
    if ur.Username == "" {
        return errors.New("username is required")
    }
    if ur.Email == "" {
        return errors.New("email is required")
    }
    if ur.Password == "" {
        return errors.New("password is required")
    }
    if len(ur.Password) < 6 {
        return errors.New("password must be at least 6 characters")
    }
    return nil
}
```

### 2.2 Create `models/post.go`
```go
package models

import "time"

// Post represents a blog post or content item
type Post struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    AuthorID  int       `json:"author_id"`
    Author    *User     `json:"author,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// PostRequest represents post input for creation/updates
type PostRequest struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

// Validate validates post input
func (pr *PostRequest) Validate() error {
    if pr.Title == "" {
        return errors.New("title is required")
    }
    if pr.Content == "" {
        return errors.New("content is required")
    }
    return nil
}
```

---

## Step 3: Create the Utils Package

The utils package provides helper functions used across the application.

### 3.1 Create `utils/validation.go`
```go
package utils

import (
    "regexp"
    "strings"
)

// IsValidEmail checks if an email address is valid
func IsValidEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}

// IsValidUsername checks if a username is valid
func IsValidUsername(username string) bool {
    // Username must be 3-20 characters, alphanumeric and underscores only
    if len(username) < 3 || len(username) > 20 {
        return false
    }
    usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
    return usernameRegex.MatchString(username)
}

// SanitizeString removes leading/trailing whitespace and converts to lowercase
func SanitizeString(s string) string {
    return strings.ToLower(strings.TrimSpace(s))
}
```

### 3.2 Create `utils/crypto.go`
```go
package utils

import (
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPassword compares a password with its hash
func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// GenerateRandomToken generates a random token for authentication
func GenerateRandomToken(length int) (string, error) {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

// HashString creates a SHA256 hash of a string
func HashString(s string) string {
    hash := sha256.Sum256([]byte(s))
    return hex.EncodeToString(hash[:])
}
```

---

## Step 4: Create the Database Package

The database package handles all database operations.

### 4.1 Create `database/connection.go`
```go
package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// Config holds database configuration
type Config struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

// Connect establishes a database connection
func Connect(config Config) error {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        return fmt.Errorf("failed to open database: %v", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %v", err)
    }

    log.Println("Successfully connected to database")
    return nil
}

// Close closes the database connection
func Close() {
    if DB != nil {
        DB.Close()
    }
}

// InitTables creates necessary tables if they don't exist
func InitTables() error {
    userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    postTable := `
    CREATE TABLE IF NOT EXISTS posts (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        author_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    if _, err := DB.Exec(userTable); err != nil {
        return fmt.Errorf("failed to create users table: %v", err)
    }

    if _, err := DB.Exec(postTable); err != nil {
        return fmt.Errorf("failed to create posts table: %v", err)
    }

    log.Println("Database tables initialized")
    return nil
}
```

### 4.2 Create `database/queries.go`
```go
package database

import (
    "database/sql"
    "webapp/models"
    "time"
)

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
    query := `INSERT INTO users (username, email, password, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
    
    err := DB.QueryRow(query, user.Username, user.Email, user.Password, 
                      time.Now(), time.Now()).Scan(&user.ID)
    return err
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (*models.User, error) {
    query := `SELECT id, username, email, password, created_at, updated_at 
              FROM users WHERE id = $1`
    
    user := &models.User{}
    err := DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email,
                                      &user.Password, &user.CreatedAt, &user.UpdatedAt)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return user, err
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*models.User, error) {
    query := `SELECT id, username, email, password, created_at, updated_at 
              FROM users WHERE username = $1`
    
    user := &models.User{}
    err := DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email,
                                            &user.Password, &user.CreatedAt, &user.UpdatedAt)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return user, err
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]models.User, error) {
    query := `SELECT id, username, email, created_at, updated_at FROM users`
    
    rows, err := DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Username, &user.Email, 
                        &user.CreatedAt, &user.UpdatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

// CreatePost inserts a new post into the database
func CreatePost(post *models.Post) error {
    query := `INSERT INTO posts (title, content, author_id, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
    
    err := DB.QueryRow(query, post.Title, post.Content, post.AuthorID,
                      time.Now(), time.Now()).Scan(&post.ID)
    return err
}

// GetPostByID retrieves a post by ID with author information
func GetPostByID(id int) (*models.Post, error) {
    query := `SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at,
                     u.id, u.username, u.email, u.created_at, u.updated_at
              FROM posts p
              JOIN users u ON p.author_id = u.id
              WHERE p.id = $1`
    
    post := &models.Post{Author: &models.User{}}
    err := DB.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content,
                                      &post.AuthorID, &post.CreatedAt, &post.UpdatedAt,
                                      &post.Author.ID, &post.Author.Username, &post.Author.Email,
                                      &post.Author.CreatedAt, &post.Author.UpdatedAt)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return post, err
}

// GetAllPosts retrieves all posts with author information
func GetAllPosts() ([]models.Post, error) {
    query := `SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at,
                     u.id, u.username, u.email, u.created_at, u.updated_at
              FROM posts p
              JOIN users u ON p.author_id = u.id
              ORDER BY p.created_at DESC`
    
    rows, err := DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        post.Author = &models.User{}
        
        err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID,
                        &post.CreatedAt, &post.UpdatedAt,
                        &post.Author.ID, &post.Author.Username, &post.Author.Email,
                        &post.Author.CreatedAt, &post.Author.UpdatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    return posts, nil
}
```

---

## Step 5: Create the Handlers Package

The handlers package contains HTTP request handlers for our API endpoints.

### 5.1 Create `handlers/user.go`
```go
package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"

    "webapp/database"
    "webapp/models"
    "webapp/utils"

    "github.com/gorilla/mux"
)

// CreateUser handles user registration
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var userReq models.UserRequest
    
    if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Validate input
    if err := userReq.Validate(); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Additional validation
    if !utils.IsValidEmail(userReq.Email) {
        http.Error(w, "Invalid email format", http.StatusBadRequest)
        return
    }

    if !utils.IsValidUsername(userReq.Username) {
        http.Error(w, "Invalid username format", http.StatusBadRequest)
        return
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(userReq.Password)
    if err != nil {
        http.Error(w, "Error processing password", http.StatusInternalServerError)
        return
    }

    // Create user
    user := &models.User{
        Username: utils.SanitizeString(userReq.Username),
        Email:    utils.SanitizeString(userReq.Email),
        Password: hashedPassword,
    }

    if err := database.CreateUser(user); err != nil {
        if strings.Contains(err.Error(), "duplicate key") {
            http.Error(w, "Username or email already exists", http.StatusConflict)
            return
        }
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    // Return user without password
    user.Password = ""
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// GetUser handles retrieving a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := database.GetUserByID(id)
    if err != nil {
        http.Error(w, "Error retrieving user", http.StatusInternalServerError)
        return
    }

    if user == nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Don't return password
    user.Password = ""
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// GetAllUsers handles retrieving all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := database.GetAllUsers()
    if err != nil {
        http.Error(w, "Error retrieving users", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
```

### 5.2 Create `handlers/auth.go`
```go
package handlers

import (
    "encoding/json"
    "net/http"
    "webapp/database"
    "webapp/models"
    "webapp/utils"
)

// LoginRequest represents login credentials
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// LoginResponse represents login response
type LoginResponse struct {
    User  *models.User `json:"user"`
    Token string       `json:"token"`
}

// Login handles user authentication
func Login(w http.ResponseWriter, r *http.Request) {
    var loginReq LoginRequest
    
    if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if loginReq.Username == "" || loginReq.Password == "" {
        http.Error(w, "Username and password are required", http.StatusBadRequest)
        return
    }

    // Get user from database
    user, err := database.GetUserByUsername(loginReq.Username)
    if err != nil {
        http.Error(w, "Error retrieving user", http.StatusInternalServerError)
        return
    }

    if user == nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Check password
    if !utils.CheckPassword(loginReq.Password, user.Password) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Generate token (simple token for demo)
    token, err := utils.GenerateRandomToken(32)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    // Don't return password
    user.Password = ""

    response := LoginResponse{
        User:  user,
        Token: token,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Logout handles user logout (simple implementation)
func Logout(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}
```

---

## Step 6: Create the Main Application

### 6.1 Create `main.go`
```go
package main

import (
    "log"
    "net/http"
    "os"

    "webapp/database"
    "webapp/handlers"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    // Database configuration
    dbConfig := database.Config{
        Host:     getEnv("DB_HOST", "localhost"),
        Port:     5432,
        User:     getEnv("DB_USER", "postgres"),
        Password: getEnv("DB_PASSWORD", "password"),
        DBName:   getEnv("DB_NAME", "webapp"),
        SSLMode:  getEnv("DB_SSLMODE", "disable"),
    }

    // Connect to database
    if err := database.Connect(dbConfig); err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer database.Close()

    // Initialize database tables
    if err := database.InitTables(); err != nil {
        log.Fatal("Failed to initialize database tables:", err)
    }

    // Create router
    router := mux.NewRouter()

    // API routes
    api := router.PathPrefix("/api/v1").Subrouter()

    // User routes
    api.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    api.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    api.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")

    // Auth routes
    api.HandleFunc("/auth/login", handlers.Login).Methods("POST")
    api.HandleFunc("/auth/logout", handlers.Logout).Methods("POST")

    // Health check
    router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    }).Methods("GET")

    // CORS middleware
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
    })

    handler := c.Handler(router)

    port := getEnv("PORT", "8080")
    log.Printf("Server starting on port %s", port)
    
    if err := http.ListenAndServe(":"+port, handler); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}

// getEnv gets environment variable with default fallback
func getEnv(key, defaultVal string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultVal
}
```

---

## Step 7: Add Dependencies

### 7.1 Add required dependencies
```bash
go get github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/rs/cors
go get golang.org/x/crypto/bcrypt
```

### 7.2 Tidy dependencies
```bash
go mod tidy
```

---

## Step 8: Database Setup

### 8.1 Install PostgreSQL
For macOS:
```bash
brew install postgresql
brew services start postgresql
```

For Ubuntu/Debian:
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
```

### 8.2 Create Database
```bash
sudo -u postgres createuser --interactive
sudo -u postgres createdb webapp
```

Or use SQL:
```sql
CREATE DATABASE webapp;
CREATE USER webapp_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE webapp TO webapp_user;
```

---

## Step 9: Environment Configuration

### 9.1 Create `.env` file (optional)
```env
DB_HOST=localhost
DB_USER=webapp_user
DB_PASSWORD=your_password
DB_NAME=webapp
DB_SSLMODE=disable
PORT=8080
```

---

## Step 10: Run the Application

### 10.1 Start the server
```bash
go run main.go
```

### 10.2 Test the API

**Health Check:**
```bash
curl http://localhost:8080/health
```

**Create User:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "password": "password123"
  }'
```

**Get All Users:**
```bash
curl http://localhost:8080/api/v1/users
```

---

## Step 11: Next Steps and Enhancements

### 11.1 Add Post Handlers
Create handlers for posts in `handlers/post.go`:
- `CreatePost`
- `GetPost`
- `GetAllPosts`
- `UpdatePost`
- `DeletePost`

### 11.2 Add Middleware
- Authentication middleware
- Logging middleware
- Rate limiting

### 11.3 Add Tests
Create test files:
- `handlers/user_test.go`
- `utils/validation_test.go`
- `database/queries_test.go`

### 11.4 Add Configuration Management
- Use `viper` for configuration
- Support multiple environments

### 11.5 Add API Documentation
- Use `swagger` for API documentation
- Add OpenAPI specifications

---

## Package Concepts Demonstrated

1. **Package Organization**: Clear separation of concerns across packages
2. **Exported/Unexported**: Public APIs vs internal implementation
3. **Import Management**: How packages depend on each other
4. **Interface Design**: Clean APIs between packages
5. **Modularity**: Each package has a specific responsibility

This project structure demonstrates professional Go development practices and provides a solid foundation for larger applications!

---

## Troubleshooting

### Common Issues:
1. **Database Connection Issues**: Check PostgreSQL is running and credentials are correct
2. **Import Path Issues**: Ensure go.mod is properly initialized
3. **Port Already in Use**: Change PORT environment variable
4. **Missing Dependencies**: Run `go mod tidy`

### Useful Commands:
```bash
go mod tidy          # Clean up dependencies
go fmt ./...         # Format all code
go vet ./...         # Check for issues
go test ./...        # Run all tests
```
