package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"sync"
	"time"
)

func processTask(id int) error {
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)

	if rand.Float32() < 0.2 {
		return fmt.Errorf("erro ao processar tarefa %d", id)
	}
	fmt.Printf("Tarefa %d concluída em %v\n", id, delay)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxConcurrency := 5
	semaphore := make(chan struct{}, maxConcurrency)

	var g errgroup.Group
	var mu sync.Mutex
	totalTasks := 20
	completedTasks := 0

	for i := 1; i <= totalTasks; i++ {
		i := i
		g.Go(func() error {
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				if err := processTask(i); err != nil {
					cancel()
					return err
				}

				mu.Lock()
				completedTasks++
				mu.Unlock()

				return nil
			}
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Processamento interrompido devido a erro: %v\n", err)
	} else {
		fmt.Println("Todas as tarefas foram concluídas com sucesso.")
	}

	fmt.Printf("Total de tarefas concluídas: %d/%d\n", completedTasks, totalTasks)
}
