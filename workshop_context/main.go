package main

import (
	"context"
	"log"
	"time"

	"github.com/arminaaki/LearnGolang/workshop_context/client"
)

const defaultSleepTime = 3 * time.Second
const defaultContextTimeout = 2 * time.Second

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), defaultContextTimeout)
	defer cancel()

	cli := client.NewClient(client.WithSleepTime(defaultSleepTime), client.WithContext(ctx))

	if err := cli.Sleep("Hello from %s\n", "World"); err != nil {
		log.Fatal(err)
	}
}
