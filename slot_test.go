package redis

import (
	"fmt"
	"testing"
)

func TestSlotSetStatus(t *testing.T) {
	srcHost := "10.77.9.52"
	srcPort := 6380
	redisClient := NewClient(&Options{
		Addr:     fmt.Sprintf("%s:%d", srcHost, srcPort),
		Password: "",
		DB:       0,
	})
	// statusRs := redisClient.SlotsSetSlot(1, SlotStatusStable)
	statusRs := redisClient.SlotsSetSlot(1, SlotStatusMigrating)
	t.Logf("\n++++++++++= SlotsSetSlot rs: %v", statusRs)

	destHost := srcHost
	destPort := 6390
	mgrtRs := redisClient.SlotsMgrtSlot(destHost, destPort, 100, 0, 5)
	t.Logf("\n++++++++++= SlotsMgrtSlot rs: %v", mgrtRs)

	hashKeyrs := redisClient.SlotsHashKey("k1", "k2")
	t.Logf("\n++++++++++= SlotsHashKey rs: %v", hashKeyrs)

	slotInfoRs := redisClient.SlotsInfo(0, 10)
	t.Logf("\n++++++++++=  SlotsInfo rs: %v", slotInfoRs)

	mgrRs := redisClient.SlotsMgrtState(0, 10)
	t.Logf("\n++++++++++=  mgrRs rs: %v", mgrRs)
}
