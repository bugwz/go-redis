package redis

import (
	"fmt"
	"testing"
)

func TestSlotCmd(t *testing.T) {
	t.Log("will start test slot cmds!")
	srcHost := "10.77.9.52"
	srcPort := 6380
	redisClient := NewClient(&Options{
		Addr:     fmt.Sprintf("%s:%d", srcHost, srcPort),
		Password: "",
		DB:       0,
	})
	// for i := 0; i < 1024; i++ {
	// 	key := fmt.Sprintf("k_%d", i)
	// 	value := fmt.Sprintf("v_%d", i)
	// 	setRs := redisClient.Set(key, value, time.Duration(1000*1000*1000*100))
	// 	t.Logf("set %s %s, rs: %v", key, value, setRs)
	// }
	// statusRs := redisClient.SlotsSetSlot(1, SlotStatusStable)
	statusRs := redisClient.SlotsSetSlot(1, SlotStatusMigrating)
	t.Logf("\n++++++++++= SlotsSetSlot rs: %v", statusRs)

	destHost := srcHost
	destPort := 6390
	mgrtRs := redisClient.SlotsMgrtSlot(destHost, destPort, 100, 0, 5)
	t.Logf("\n++++++++++= SlotsMgrtSlot rs: %v", mgrtRs)

	hashKeyrs := redisClient.SlotsHashKey("k1", "k2")
	t.Logf("\n++++++++++= SlotsHashKey rs: %v", hashKeyrs)

	slotInfoRs := redisClient.SlotsInfo(0, 1023)
	t.Logf("\n++++++++++=  SlotsInfo rs: %v", slotInfoRs)

	mgrRs := redisClient.SlotsMgrtState(0, 10)
	t.Logf("\n++++++++++=  mgrRs rs: %v", mgrRs)
}
