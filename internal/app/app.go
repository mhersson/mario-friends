package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
)

var (
	sigChannel      = make(chan os.Signal, 1)
	messageFile     = os.Getenv("MESSAGE_FILE")
	nextHopHostname = os.Getenv("NEXT_HOP_HOSTNAME")
	nextHopPort     = os.Getenv("NEXT_HOP_PORT")
)

func checkError(w *http.ResponseWriter, msg string, err error) {
	if err != nil {
		fmt.Fprintf(*w, "%s: %s\n", msg, err)
	}
}

func getMessageFromNextHop() (string, error) {
	if nextHopHostname == "" {
		return "", nil
	}

	client := http.Client{
		Timeout: 250 * time.Millisecond, //nolint: go-lint
	}

	url := fmt.Sprintf("http://%s:%s", nextHopHostname, nextHopPort)

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return string(body), nil
}

func getMessageFromFile() string {
	c, _ := os.ReadFile(messageFile)

	return string(c)
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}

	return r.RemoteAddr
}

func SignalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not supported, I think")

		return
	}

	log.Printf("New %s request from %s\n", r.Method, getIP(r))

	sigChannel <- syscall.SIGTERM
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead && r.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not supported")
	}
}

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 Game Over")

		return
	}

	log.Printf("New %s request from %s\n", r.Method, getIP(r))

	currentHostMessage := getMessageFromFile()
	fmt.Fprint(w, currentHostMessage)

	if currentHostMessage == "" {
		hostname, err := os.Hostname()
		checkError(&w, "Failed to get hostname", err)

		interfaces, err := net.Interfaces()
		checkError(&w, "Failed to get interfaces", err)

		fmt.Fprintf(w, "Hello!\nMy name is %s!\n", hostname)
		fmt.Fprintln(w, "You can find me here:")

		for _, iface := range interfaces {
			addr, _ := iface.Addrs()
			if len(addr) > 0 && iface.Name != "lo" {
				fmt.Fprintf(w, "%s - %s\n", iface.Name, addr[0])
			}
		}
	}

	nextHopMessage, err := getMessageFromNextHop()
	checkError(&w, "Failed to get message from next hop", err)

	if nextHopMessage != "" {
		fmt.Fprint(w, "\nEntering level "+nextHopHostname+"...\n")
		fmt.Fprint(w, "\n\n"+nextHopMessage)
	}
}

func Run() {
	hostname, _ := os.Hostname()

	log.Printf("Started %s\n", hostname)

	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", SimpleHandler)
	http.HandleFunc("/signal", SignalHandler)
	http.HandleFunc("/.well-known/health", HealthHandler)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("An unexpected error occurred: %v\n", err)
		}
	}()

	<-sigChannel

	log.Println("Received SIGTERM stopping server")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("An error occurred when stopping the server: %v\n", err)
	}
}
