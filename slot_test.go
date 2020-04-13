package redis

import (
	"testing"
)

func TestSlotSetStatus(t *testing.T) {
	redisClient := NewClient(&Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})
	statusRs := redisClient.SlotsSetSlot(1, SlotStatusMigrating)
	t.Logf("\n++++++++++= SlotsSetSlot rs: %v", statusRs)

	slotInfoRs := redisClient.SlotsInfo(0, 10)
	t.Logf("\n++++++++++=  SlotsInfo rs: %v", slotInfoRs)

	mgrRs := redisClient.SlotsMgrtState(0, 10)
	t.Logf("\n++++++++++=  mgrRs rs: %v", mgrRs)
}
