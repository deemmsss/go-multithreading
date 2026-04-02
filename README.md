# Go Multithreading Banking Lab

## About

This project demonstrates concurrency concepts in Go using a banking simulation. It progressively builds from basic sequential operations to goroutines, channels, and mutex-based synchronization — showing how real-world banking systems handle simultaneous transactions safely.

## Folder Structure

```
go-multithreading/
├── go.mod
├── README.md
├── images/                     # Screenshots of terminal output for each part
│   ├── part3_output.png
│   ├── part4_output.png
│   ├── part5_output.png
│   └── part6_output.png
├── part3_sequential/
│   └── main.go                 # Basic withdrawal function, no concurrency
├── part4_race/
│   └── main.go                 # Goroutines without protection (race condition)
├── part5_channel/
│   └── main.go                 # Channel-based transaction processor (double-entry)
└── part6_mutex/
    └── main.go                 # Mutex + WaitGroup for safe concurrency, deposit added
```

## Prerequisites

- Go 1.21+ installed ([download here](https://go.dev/dl/))
- Verify installation: `go version`

## How to Run

From the `go-multithreading/` root directory, run each part individually:

```bash
# Part 3 – Sequential operations
go run part3_sequential/main.go

# Part 4 – Race condition demo
go run part4_race/main.go

# Part 4 – Run with Go's race detector to confirm the race condition
go run -race part4_race/main.go

# Part 5 – Channel-based processing
go run part5_channel/main.go

# Part 6 – Mutex-synchronized concurrency
go run part6_mutex/main.go
```

## What Each Part Demonstrates

### Part 3: Sequential Operations
A basic `withdraw()` function that reduces the customer balance one transaction at a time. No concurrency is involved — transactions run in order. Starting balance is $1,000; a $700 phone transfer succeeds, then a $500 ATM withdrawal is declined due to insufficient funds.

### Part 4: Concurrent Operations (Race Condition)
The same two withdrawals now run as goroutines (concurrently). Without any locking, both goroutines read the balance as $1,000 before either writes, so both withdrawals succeed — resulting in a negative or incorrect final balance. This demonstrates a classic race condition.

### Part 5: Channel-Based Transaction Processing
Instead of goroutines modifying the balance directly, each transaction is sent through a Go channel to a single processor goroutine. The processor handles all balance updates sequentially, implementing a double-entry model where every customer withdrawal increases the bank ledger and vice versa. This mirrors how real banking systems centralize ledger updates to prevent inconsistencies.

### Part 6: Mutex Synchronization
Three concurrent goroutines (Phone Transfer, ATM Withdrawal, Salary Deposit) run simultaneously, but access to the shared balance is protected by `sync.Mutex`. A `sync.WaitGroup` ensures the main function waits for all goroutines to finish before printing the final balance. This guarantees correct output regardless of execution order.

## Observation: With vs. Without Mutex

**Without Mutex (Part 4):** Both goroutines read `customerBalance = 1000` before either modifies it. Both pass the `if customerBalance >= amount` check, and both withdraw — leaving the balance at -$200 (or some other incorrect value). The outcome varies between runs because goroutine scheduling is non-deterministic. Running with `go run -race` confirms the data race.

**With Mutex (Part 6):** Only one goroutine can access the balance at a time. The first goroutine to acquire the lock completes its transaction fully before the next one reads the balance. This ensures the balance is always accurate and the insufficient-funds check works correctly.

## Concepts Covered

- **Goroutines** — lightweight concurrent functions in Go
- **Race conditions** — what happens when goroutines share data without synchronization
- **Channels** — Go's built-in mechanism for safe communication between goroutines
- **sync.Mutex** — mutual exclusion lock to protect shared variables
- **sync.WaitGroup** — mechanism to wait for a group of goroutines to finish
- **Double-entry bookkeeping** — centralized transaction processing pattern
