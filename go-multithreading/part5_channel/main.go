package main
 
import (
    "fmt"
    "sync"
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
 
    fmt.Println("\n=== Part 5: Channel-Based Transaction Processing ===")
 
    customerBalance = 1000
    bankBalance = 0
 
    txChannel := make(chan Transaction)
 
    // Start processor goroutine
    go transactionProcessor(txChannel)
 
    // Send transactions through the channel
    txChannel <- Transaction{Amount: 700, Source: "Phone Transfer", CustomerID: "1"}
    txChannel <- Transaction{Amount: 500, Source: "ATM Withdrawal", CustomerID: "1"}
    txChannel <- Transaction{Amount: 200, Source: "Online Purchase", CustomerID: "1"}
    txChannel <- Transaction{Amount: -1500, Source: "Salary Deposit", CustomerID: "1"}
 
    // Close the channel after sending all transactions
    close(txChannel)
}
 
// Goroutine to process all transactions
func transactionProcessor(ch <-chan Transaction) {
    for tx := range ch {
        fmt.Printf("[Processor] Processing %s for customer %s: $%d\n", tx.Source, tx.CustomerID, tx.Amount)
 
        mu.Lock()
        if tx.Amount > 0 { // withdrawal
            if customerBalance >= tx.Amount {
                customerBalance -= tx.Amount
                bankBalance += tx.Amount
                fmt.Printf("[Processor] Withdrawal approved. Customer balance: $%d, Bank ledger: $%d\n",
                    customerBalance, bankBalance)
            } else {
                fmt.Printf("[Processor] Withdrawal declined - Insufficient funds.\n")
            }
        } else { // deposit
            depositAmount := -tx.Amount
            customerBalance += depositAmount
            bankBalance -= depositAmount
            fmt.Printf("[Processor] Deposit processed. Customer balance: $%d, Bank ledger: $%d\n",
                customerBalance, bankBalance)
        }
        mu.Unlock()
    }
 
    fmt.Printf("Final Customer Balance: $%d, Bank Ledger: $%d\n", customerBalance, bankBalance)
}