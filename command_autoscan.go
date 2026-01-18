package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	inputoutput "github.com/mortalglitch/hashwood/internal/input_output"
)

type AutoScanManager struct {
	mu          sync.Mutex
	activeScans map[string]context.CancelFunc
}

func NewAutoScanManager() *AutoScanManager {
	return &AutoScanManager{
		activeScans: make(map[string]context.CancelFunc),
	}
}

// Commands available
// autoscan start <time> <directory>
// autoscan stop <directory>
func (asm *AutoScanManager) CommandAutoScan(words []string, cfg *appConfig) {

	if len(words) == 4 {
		asm.mu.Lock()
		selectedTime := words[2]

		selectedTimeInt, err := strconv.Atoi(selectedTime)
		if err != nil {
			fmt.Println(err)
			return
		}

		targetDirectory, err := filepath.Abs(words[3])
		if err != nil {
			fmt.Println(err)
			return
		}

		if cancel, exists := asm.activeScans[targetDirectory]; exists {
			cancel()
		}

		ctx, cancel := context.WithCancel(context.Background())
		asm.activeScans[targetDirectory] = cancel
		asm.mu.Unlock()

		go func(ctx context.Context) {
			ticker := time.NewTicker(time.Duration(selectedTimeInt) * time.Second)
			defer ticker.Stop()

			defer func() {
				asm.mu.Lock()
				delete(asm.activeScans, targetDirectory)
				asm.mu.Unlock()
			}()

			fmt.Printf("Starting autoscan on %s every %s seconds\n", targetDirectory, selectedTime)

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Stopping scan on %s...\n", targetDirectory)
					return
				case <-ticker.C:
					scanDirectory(targetDirectory, cfg)
					fmt.Print("> ")
				}
			}
		}(ctx)

		return
	} else if len(words) == 3 {
		stopCommand := words[1]

		if stopCommand == "stop" {
			asm.mu.Lock()
			defer asm.mu.Unlock()

			targetDirectory, err := filepath.Abs(words[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			if cancel, exists := asm.activeScans[targetDirectory]; exists {
				cancel()
				fmt.Printf("Scan for %s halted.\n", targetDirectory)
			} else {
				fmt.Printf("No scans found for %s\n", targetDirectory)
			}
		} else {
			inputoutput.PrintAutoScan()
		}
		return
	}
}
