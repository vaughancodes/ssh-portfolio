package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

const host = "0.0.0.0"

func main() {
	port := flag.Int("p", 22, "port to listen on")
	flag.Parse()

	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, *port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatalf("Could not create server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	log.Printf("Starting SSH server on %s:%d", host, *port)
	log.Printf("Connect with: ssh -p %d localhost", *port)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-done
	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	w := pty.Window.Width
	h := pty.Window.Height
	if w == 0 {
		w = 80
	}
	if h == 0 {
		h = 24
	}
	renderer := bubbletea.MakeRenderer(s)
	m := newModel(w, h, renderer)
	return m, []tea.ProgramOption{tea.WithAltScreen(), tea.WithMouseAllMotion()}
}
