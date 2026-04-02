package main
 
import (
    "fmt"
    "sync"
    "time"
)
 
// Shared struct and data
type Transaction struct {
    Amount     int
    Source     string
    CustomerID string
}
 
var customerBalance = 1000
var bankBalance = 0
var mu sync.Mutex
 
func main() {
    fmt.Println("========================================")
    fmt.Println("    Go Multithreading Banking Lab")
    fmt.Println("========================================")
 
    fmt.Println("\n=== Part 6: Concurrent Operations with Mutex Synchronization ===")
 
    customerBalance = 1000
 
    var wg sync.WaitGroup
    wg.Add(3)
 
    // Concurrent goroutines with proper synchronization
    go withdrawWithLock(&wg, 700, "Phone Transfer")
    go withdrawWithLock(&wg, 500, "ATM Withdrawal")
    go depositWithLock(&wg, 1500, "Salary Deposit")
 
    wg.Wait() // wait for all goroutines to finish
 
    fmt.Printf("Final Balance: $%d (Guaranteed correct with mutex)\n", customerBalance)
}
 
// Thread-safe withdrawal using mutex
func withdrawWithLock(wg *sync.WaitGroup, amount int, source string) {
    defer wg.Done()
    mu.Lock()
    defer mu.Unlock()
 
    fmt.Printf("[%s] Attempting to withdraw $%d\n", source, amount)
    if customerBalance >= amount {
        customerBalance -= amount
        fmt.Printf("[%s] Withdrawal successful. New balance: $%d\n", source, customerBalance)
    } else {
        fmt.Printf("[%s] Insufficient funds for withdrawal of $%d\n", source, amount)
    }
}
 
// Thread-safe deposit using mutex
func depositWithLock(wg *sync.WaitGroup, amount int, source string) {
    defer wg.Done()
    mu.Lock()
    defer mu.Unlock()
 
    fmt.Printf("[%s] Depositing $%d\n", source, amount)
    customerBalance += amount
    fmt.Printf("[%s] Deposit successful. New balance: $%d\n", source, customerBalance)
    time.Sleep(5 * time.Millisecond)
}