package main

import (
	"context"
	"os"
	"os/signal"
	"role-manager-bot/internal/bot"
	"role-manager-bot/internal/database"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	stopBot := bot.Init(ctx)

	if err := database.Connect(ctx); err != nil {
		panic(err)
	}
	waitExitSignal()
	stopBot()
	cancel()
}
func waitExitSignal() os.Signal {
	ch := make(chan os.Signal, 3)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-ch
}
