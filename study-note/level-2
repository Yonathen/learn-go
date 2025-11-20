# Level 2: Multiple Files, Same Package - Better Organization

## 🎯 Evolution: From Simple to Organized

You've mastered Level 1 - now let's evolve! Level 2 is about **organization without complexity**. You'll learn to structure code better while keeping the simplicity of a single package.

---

## Chapter 1: The Organization Challenge

### 1.1 The Growing Pains

As your Level 1 projects grow, you might have experienced:
- **Long files** - scrolling forever to find functions
- **Mixed concerns** - user input mixed with calculations
- **Hard to navigate** - where did I put that function?
- **Difficult maintenance** - changing one thing affects everything

**This is normal!** It's a sign you're ready for better organization.

### 1.2 The Level 2 Solution

Level 2 keeps the **same package benefits** but adds **logical organization**:
- ✅ Still one package (simplicity maintained)
- ✅ Functions still call each other freely
- ✅ No complex imports
- ✅ BUT: organized by purpose and responsibility

**Think of it like:** Organizing your workshop into tool stations while keeping everything accessible.

---

## Chapter 2: Understanding File-Based Organization

### 2.1 The Core Principle

**Same Package, Different Files = Organized Workshop**

Instead of one big messy workbench, you have:
- **Input Station** (`input.go`) - tools for getting data
- **Processing Station** (`process.go`) - tools for transforming data
- **Output Station** (`output.go`) - tools for displaying results
- **Control Station** (`main.go`) - coordinates everything

### 2.2 Why This Works So Well

**For Your Brain:**
- Easy to find things (logical locations)
- Clear separation of concerns
- Reduced cognitive load

**For Collaboration:**
- Different people can work on different files
- Fewer merge conflicts
- Clear ownership of functionality

**For Maintenance:**
- Changes are localized
- Easier to test specific areas
- Cleaner code organization

---

## Chapter 3: Designing Your File Structure

### 3.1 The "By Purpose" Approach

Organize files by what they DO, not by what they ARE:

**❌ Bad (by type):**
```
structs.go    // All structs
functions.go  // All functions  
variables.go  // All variables
```

**✅ Good (by purpose):**
```
user.go       // Everything about users
auth.go       // Everything about authentication
database.go   // Everything about data storage
```

### 3.2 Common File Organization Patterns

**Pattern 1: By Feature**
```
main.go           // Entry point and coordination
user-management.go // User creation, updating, deletion
file-processing.go // File reading, writing, parsing
reporting.go      // Generate reports and summaries
```

**Pattern 2: By Layer**
```
main.go       // Entry point
input.go      // Data collection and validation
business.go   // Core logic and calculations  
output.go     // Presentation and formatting
```

**Pattern 3: By Domain**
```
main.go       // Orchestration
customers.go  // Customer-related functionality
orders.go     // Order processing
inventory.go  // Stock management
```

### 3.3 The "Single Responsibility" Rule

Each file should have ONE main job:

**Good example:**
```go
// file: email.go
// Purpose: Everything about sending emails

func sendWelcomeEmail(user User) error { ... }
func sendPasswordResetEmail(email string) error { ... }
func validateEmailAddress(email string) bool { ... }
func formatEmailTemplate(template string, data map[string]string) string { ... }
```

**Why this works:** When you need email functionality, you know exactly where to look.

---

## Chapter 4: Practical File Organization Strategies

### 4.1 The Dependency Flow Strategy

Organize files so dependencies flow in one direction:

```
main.go           ← Entry point (depends on everything)
    ↓
coordination.go   ← High-level coordination
    ↓  
business.go      ← Business logic
    ↓
data.go          ← Data operations (depends on nothing)
```

**Why this matters:** Clean dependency flow makes code easier to understand and test.

### 4.2 The Interface-First Strategy

Start by designing what each file will PROVIDE:

```go
// What should user.go provide?
func CreateUser(name, email string) User
func ValidateUser(user User) error
func DisplayUser(user User)

// What should auth.go provide?
func Login(username, password string) bool
func Logout(user User)
func CheckPermission(user User, action string) bool
```

**Process:**
1. Design the interface (what functions you need)
2. Organize functions into logical files
3. Implement the functions

### 4.3 The Evolution Strategy

Start simple, then split when files get too big:

**Stage 1:** Everything in main.go
**Stage 2:** Split by obvious boundaries (input, processing, output)
**Stage 3:** Split further as complexity grows

**Rule of thumb:** If a file is over 200 lines, consider splitting it.

---

## Chapter 5: Communication Patterns Between Files

### 5.1 Function Calls Across Files

This is the bread and butter of Level 2:

```go
// In user.go
func CreateUser(name, email string) User {
    user := User{Name: name, Email: email}
    if ValidateEmail(email) {  // Function from validation.go
        return user
    }
    return User{}
}

// In validation.go
func ValidateEmail(email string) bool {
    return strings.Contains(email, "@")
}

// In main.go
func main() {
    user := CreateUser("John", "john@example.com")  // From user.go
    DisplayUser(user)  // From display.go
}
```

**The beauty:** Functions call each other naturally across files!

### 5.2 Shared Variables and Constants

Some data needs to be shared across files:

```go
// In config.go
var AppName = "My Application"
var Version = "1.0.0"
const MaxUsers = 100

// In main.go
func main() {
    fmt.Printf("Starting %s v%s\n", AppName, Version)
}

// In user.go  
func CreateUser(name string) error {
    if userCount >= MaxUsers {
        return fmt.Errorf("maximum users reached")
    }
    // ... create user
}
```

**Best practice:** Keep shared data in a dedicated file (like `config.go` or `globals.go`).

### 5.3 Type Definitions and Structs

Define types where they make the most sense:

```go
// In types.go (or models.go)
type User struct {
    ID    int
    Name  string
    Email string
}

type Order struct {
    ID       int
    UserID   int
    Products []Product
}

// Now use these types in other files
// In user.go
func CreateUser(name, email string) User { ... }

// In orders.go  
func CreateOrder(userID int, products []Product) Order { ... }
```

---

## Chapter 6: Real-World Example - Building a Task Manager

Let's build a complete example to see Level 2 in action:

### 6.1 Project Structure
```
task-manager/
├── main.go          // Entry point and menu
├── types.go         // Data structures
├── tasks.go         // Task operations
├── input.go         // User input handling
├── display.go       // Output formatting
├── storage.go       // File save/load
└── utils.go         // Helper functions
```

### 6.2 types.go - Define Your Data
```go
package main

import "time"

type Task struct {
    ID          int
    Title       string
    Description string
    Done        bool
    CreatedAt   time.Time
}

type TaskList struct {
    Tasks []Task
    nextID int
}
```

### 6.3 tasks.go - Core Business Logic
```go
package main

import "time"

func (tl *TaskList) AddTask(title, description string) {
    task := Task{
        ID:          tl.nextID,
        Title:       title,
        Description: description,
        Done:        false,
        CreatedAt:   time.Now(),
    }
    tl.Tasks = append(tl.Tasks, task)
    tl.nextID++
}

func (tl *TaskList) CompleteTask(id int) bool {
    for i := range tl.Tasks {
        if tl.Tasks[i].ID == id {
            tl.Tasks[i].Done = true
            return true
        }
    }
    return false
}

func (tl *TaskList) GetPendingTasks() []Task {
    var pending []Task
    for _, task := range tl.Tasks {
        if !task.Done {
            pending = append(pending, task)
        }
    }
    return pending
}
```

### 6.4 input.go - User Interaction
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func getUserChoice() int {
    fmt.Print("Your choice: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    choice, _ := strconv.Atoi(strings.TrimSpace(input))
    return choice
}

func getTaskDetails() (string, string) {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("Task title: ")
    title, _ := reader.ReadString('\n')
    title = strings.TrimSpace(title)
    
    fmt.Print("Task description: ")
    description, _ := reader.ReadString('\n')
    description = strings.TrimSpace(description)
    
    return title, description
}

func getTaskID() int {
    fmt.Print("Enter task ID: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(input))
    return id
}
```

### 6.5 display.go - Pretty Output
```go
package main

import (
    "fmt"
    "strings"
)

func showMenu() {
    fmt.Println("\n" + strings.Repeat("=", 30))
    fmt.Println("       TASK MANAGER")
    fmt.Println(strings.Repeat("=", 30))
    fmt.Println("1. Add Task")
    fmt.Println("2. List Tasks")
    fmt.Println("3. Complete Task")
    fmt.Println("4. Save & Exit")
    fmt.Println(strings.Repeat("=", 30))
}

func displayTasks(tasks []Task) {
    if len(tasks) == 0 {
        fmt.Println("No tasks found!")
        return
    }
    
    fmt.Println("\nYour Tasks:")
    fmt.Println(strings.Repeat("-", 50))
    for _, task := range tasks {
        status := "[ ]"
        if task.Done {
            status = "[✓]"
        }
        fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
        fmt.Printf("    %s\n", task.Description)
        fmt.Printf("    Created: %s\n", task.CreatedAt.Format("2006-01-02 15:04"))
        fmt.Println(strings.Repeat("-", 50))
    }
}

func showMessage(message string) {
    fmt.Printf("\n>>> %s <<<\n", message)
}
```

### 6.6 main.go - Orchestration
```go
package main

func main() {
    taskList := &TaskList{nextID: 1}
    loadTasks(taskList) // From storage.go
    
    for {
        showMenu()        // From display.go
        choice := getUserChoice()  // From input.go
        
        switch choice {
        case 1:
            title, desc := getTaskDetails()  // From input.go
            taskList.AddTask(title, desc)    // From tasks.go
            showMessage("Task added successfully!")
            
        case 2:
            displayTasks(taskList.Tasks)     // From display.go
            
        case 3:
            id := getTaskID()                // From input.go
            if taskList.CompleteTask(id) {   // From tasks.go
                showMessage("Task completed!")
            } else {
                showMessage("Task not found!")
            }
            
        case 4:
            saveTasks(taskList)              // From storage.go
            showMessage("Tasks saved. Goodbye!")
            return
            
        default:
            showMessage("Invalid choice!")
        }
    }
}
```

### 6.7 What We Achieved

**✅ Clear Organization:** Each file has a specific purpose
**✅ Easy to Understand:** Anyone can see where functionality lives
**✅ Easy to Extend:** Want to add categories? Create `categories.go`
**✅ Easy to Test:** Can test each file's functions independently
**✅ Still Simple:** No complex imports or package management

---

## Chapter 7: Best Practices for Level 2

### 7.1 File Naming Conventions

**Do:**
- Use descriptive names: `user-management.go`, `email-sender.go`
- Be consistent: if you use dashes, use them everywhere
- Use lowercase with dashes or underscores
- Make purpose immediately clear

**Don't:**
- Use generic names: `utils.go`, `helpers.go`, `misc.go`
- Mix naming styles: `userMgmt.go` and `email-sender.go`
- Use abbreviations unless very common

### 7.2 Function Organization Within Files

**Group related functions together:**
```go
// user.go

// User creation functions
func CreateUser(name, email string) User { ... }
func ValidateUser(user User) error { ... }

// User query functions  
func FindUserByEmail(email string) (User, error) { ... }
func GetAllUsers() []User { ... }

// User update functions
func UpdateUser(user User) error { ... }
func DeleteUser(id int) error { ... }
```

### 7.3 Documentation and Comments

**File-level comments:**
```go
// Package main implements a task management system.
// 
// This file (tasks.go) contains all core task operations including
// creation, completion, and querying of tasks.
package main
```

**Function-level comments:**
```go
// AddTask creates a new task and adds it to the task list.
// It automatically assigns an ID and sets the creation time.
func (tl *TaskList) AddTask(title, description string) {
```

### 7.4 Error Handling Patterns

**Consistent error handling across files:**
```go
// In user.go
func CreateUser(email string) (User, error) {
    if !isValidEmail(email) {
        return User{}, fmt.Errorf("invalid email: %s", email)
    }
    return User{Email: email}, nil
}

// In main.go
user, err := CreateUser("invalid-email")
if err != nil {
    showMessage("Error: " + err.Error())
    continue
}
```

---

## Chapter 8: Common Patterns and Anti-Patterns

### 8.1 Good Patterns

**The Coordinator Pattern:**
```go
// main.go coordinates but doesn't do the work
func main() {
    data := loadData()      // From storage.go
    results := process(data) // From processor.go
    display(results)        // From display.go
}
```

**The Feature Module Pattern:**
```go
// user.go - everything about users
// order.go - everything about orders  
// product.go - everything about products
```

**The Layer Pattern:**
```go
// data.go - data operations
// business.go - business logic
// presentation.go - user interface
```

### 8.2 Anti-Patterns to Avoid

**The God File:**
```go
// helpers.go - contains everything that doesn't fit elsewhere
// This becomes a dumping ground!
```

**The Circular Dependency:**
```go
// user.go calls functions from order.go
// order.go calls functions from user.go
// Creates confusion about responsibility
```

**The Scattered Feature:**
```go
// User functions spread across multiple files
// Hard to find all user-related functionality
```

---

## Chapter 9: Transitioning from Level 1

### 9.1 Refactoring Strategy

**Step 1: Identify Natural Boundaries**
Look for functions that naturally group together:
- All functions dealing with files
- All functions dealing with user input
- All functions doing calculations

**Step 2: Create New Files**
Move related functions to new files:
```bash
# Before
main.go (500 lines)

# After
main.go (50 lines - coordination)
input.go (100 lines - user input)
calculations.go (150 lines - math operations)
output.go (100 lines - display)
files.go (100 lines - file operations)
```

**Step 3: Test After Each Move**
Make sure everything still works after each file creation.

### 9.2 Gradual Migration

**Don't try to reorganize everything at once!**

**Week 1:** Move obvious groups (file operations)
**Week 2:** Move user interface functions  
**Week 3:** Move business logic functions
**Week 4:** Clean up and optimize

---

## Chapter 10: When You're Ready for Level 3

### 10.1 Signs You've Mastered Level 2

- [ ] You can organize functions into logical files
- [ ] You understand file-based responsibility separation
- [ ] You can navigate your codebase easily
- [ ] You think in terms of "what does this file do?"
- [ ] You can work on one file without affecting others
- [ ] Your files have clear, single purposes

### 10.2 Signs You Need Level 3

You might be ready for Level 3 when:
- **Reusability:** "I want to use this in another project"
- **Team Collaboration:** "Different teams need different parts"  
- **Clear Boundaries:** "These features are completely separate"
- **Testing:** "I want to test this independently"
- **Distribution:** "I want to share just this part"

### 10.3 Preparing for Level 3

**Mental preparation:**
- Start thinking about "what could be a separate package?"
- Identify which functions are truly independent
- Think about clear interfaces between different areas

**Technical preparation:**
- Practice writing functions that don't depend on global variables
- Write clear, documented interfaces for your file groups
- Think about error handling between different components

---

## 🎯 Key Takeaways

1. **Organization improves understanding** - logical file structure makes code easier to navigate
2. **Same package benefits remain** - functions still call each other easily
3. **Single responsibility per file** - each file should have one clear purpose
4. **Start simple, evolve gradually** - don't over-organize from the beginning
5. **Good organization enables growth** - well-structured code can grow larger without becoming messy

**The Level 2 sweet spot:** All the simplicity of a single package with much better organization!

---

## What's Next?

- **Master Level 2** with the task manager project
- **Practice file organization** with your own projects
- **When ready**: Check out **Level 3 Guide** (Multiple Packages)

Remember: Level 2 is where many successful Go projects live happily. Don't rush to Level 3 unless you really need it! 🚀
