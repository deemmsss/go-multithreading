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
    customerBalance = 1000
 
    fmt.Println("========================================")
    fmt.Println("    Go Multithreading Banking Lab")
    fmt.Println("========================================")
 
    fmt.Println("\n=== Part 4: Concurrent Operations (Race Condition) ===")
 
    // Two goroutines accessing the same balance concurrently — race condition
    go withdrawNoLock(700, "Phone Transfer")
    go withdrawNoLock(500, "ATM Withdrawal")
 
    // Wait for goroutines to finish
    time.Sleep(1 * time.Second)
 
    fmt.Printf("Final Balance: $%d (May be incorrect due to race condition)\n", customerBalance)
}
 
// Function without mutex (intentionally causes a race condition)
func withdrawNoLock(amount int, source string) {
    fmt.Printf("[%s] Attempting to withdraw $%d\n", source, amount)
 
    if customerBalance >= amount {
        time.Sleep(10 * time.Millisecond)
        customerBalance -= amount
        fmt.Printf("[%s] Withdrawal successful. New balance: $%d\n", source, customerBalance)
    } else {
        fmt.Printf("[%s] Insufficient funds for withdrawal of $%d\n", source, amount)
    }
}