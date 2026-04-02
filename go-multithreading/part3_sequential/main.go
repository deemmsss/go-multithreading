package main
 
import (
    "fmt"
    "sync"
)
 
// Shared data and mutex
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
 
    fmt.Println("\n=== Part 3: Basic Sequential Operations ===")
    withdrawSequential(700, "Phone Transfer")
    withdrawSequential(500, "ATM Withdrawal")
    fmt.Printf("Final Balance: $%d\n", customerBalance)
}
 
// Sequential withdrawal function (no concurrency)
func withdrawSequential(amount int, source string) {
    fmt.Printf("[%s] Attempting to withdraw $%d\n", source, amount)
 
    mu.Lock() // Lock to simulate safe access (optional for sequential but good practice)
    defer mu.Unlock()
 
    if customerBalance >= amount {
        customerBalance -= amount
        fmt.Printf("[%s] Withdrawal successful. New balance: $%d\n", source, customerBalance)
    } else {
        fmt.Printf("[%s] Insufficient funds for withdrawal of $%d\n", source, amount)
    }
}