package examples

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func NotifyContext() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(sigCh)

	select {
	case sig := <-sigCh:
		switch sig {
		case os.Interrupt:
			fmt.Println("Received Ctrl+C")
		case syscall.SIGTERM:
			fmt.Println("Received SIGTERM")
		case syscall.SIGHUP:
			fmt.Println("Received SIGHUP")
		default:
			fmt.Println("Received signal:", sig)
		}
	case <-ctx.Done():
		fmt.Println("Context done:", context.Cause(ctx))
	}
}
