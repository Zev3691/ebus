package main

import (
	"context"
	"ebus"
	"fmt"
)

const (
	Topic1 = "0101"
)

func main() {
	ebus.NewMgr()
	l := ListenerFactory("l")
	ll := ListenerFactory("ll")
	lll := ListenerFactory("lll")
	ebus.Subscribe(Topic1, l, ll, lll)
	ebus.Publish(context.Background(), Topic1, "for test", ebus.Sync)
	fmt.Println()
	ebus.Remove(Topic1, l)
	ebus.Publish(context.Background(), Topic1, "for remove", ebus.Sync)
}

func ListenerFactory(name string) *ebus.Listener {
	return ebus.NewListener(func(ctx context.Context, arg any) {
		fmt.Println("Factory: "+name, arg)
	})
}
