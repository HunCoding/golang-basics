package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var (
	visitors = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func getIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 100)
		visitors[ip] = limiter
	}

	return limiter
}

func rateLimitByIP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		limiter := getVisitor(ip)
		if !limiter.Allow() {
			log.Printf("[BLOCKED] Request was block using IP: %s\n", ip)
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		log.Printf("[ALLOWD] Request was allowd using IP: %s\n", ip)
		h.ServeHTTP(w, r)
	})
}

func doRequestUsingDifferentIPs() {
	ips := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"}
	wg := sync.WaitGroup{}

	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				sendRequest(ip)
				time.Sleep(300 * time.Millisecond)
			}
		}(ip)
	}

	wg.Wait()
	fmt.Println("Simulação concluída.")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	go func() {
		http.ListenAndServe(":8080", rateLimitByIP(mux))
	}()

	doRequestUsingDifferentIPs()
}

func sendRequest(ip string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	req.RemoteAddr = ip + ":12345" // Define o IP simulado
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Falha ao enviar requisição para IP %s: %v", ip, err)
		return
	}
	defer resp.Body.Close()

	log.Printf("[RESPONSE] IP: %s - Status: %s", ip, resp.Status)
}
