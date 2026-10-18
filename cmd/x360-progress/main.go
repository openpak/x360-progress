// Command x360-progress serves the Xbox 360 progress REST API (see internal/progress).
package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/openpak/x360-common/xstoken"
	"github.com/openpak/x360-progress/internal/progress"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))
	addr := os.Getenv("X360_PROGRESS_ADDR")
	if addr == "" {
		addr = ":10041"
	}
	keyFile := os.Getenv("X360_PROGRESS_XSTS_KEY") // the token service's signing key, read-only
	key, err := os.ReadFile(keyFile)
	if err != nil || len(key) < 16 {
		log.Error("xsts signing key (X360_PROGRESS_XSTS_KEY)", "err", err, "len", len(key))
		os.Exit(2)
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Error("listen", "err", err)
		os.Exit(1)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	log.Info("x360-progress listening", "addr", ln.Addr())
	srv := &http.Server{Handler: &progress.Handler{Signer: &xstoken.Signer{Key: key}, Log: log}, ReadHeaderTimeout: 10 * time.Second}
	go func() {
		<-ctx.Done()
		sc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(sc)
	}()
	if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("serve", "err", err)
		os.Exit(1)
	}
}
