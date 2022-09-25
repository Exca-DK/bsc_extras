package eth

import (
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/metadata"
)

func TestPacketIncrement(t *testing.T) {

	var (
		size  = 50
		iters = 50
	)
	meta := NewEthMetadata(uint64(size))

	var wg sync.WaitGroup
	for i := 0; i < iters; i++ {
		wg.Add(1)
		func() {
			meta.IncrementPacket(metadata.PacketStatType(NewBlockHashesPacketStat))
			wg.Done()
		}()
	}

	wg.Wait()

	packet, err := meta.GetPacket(metadata.PacketStatType(NewBlockHashesPacketStat))
	if err != nil {
		t.Fatal(err)
	}

	value := packet.GetValue()
	if value != uint64(iters) {
		t.Fatalf("expected %v, got %v", iters, value)
	}
}

func TestDataPush(t *testing.T) {

	var (
		size  = 50
		iters = 50
	)
	meta := NewEthMetadata(uint64(size))
	item := metadata.PacketItem{
		Data: common.Hash{},
		Ts:   time.Now().Unix(),
	}
	var wg sync.WaitGroup
	for i := 0; i < iters*2; i++ {
		wg.Add(1)
		func() {
			meta.UpdatePacketData(NewBlockHashesPacketStat, item)
			wg.Done()
		}()
	}

	wg.Wait()

	packet, err := meta.GetPacket(NewBlockHashesPacketStat)
	if err != nil {
		t.Fatal(err)
	}

	data := packet.GetData()
	if len(data) != iters {
		t.Fatalf("expected %v, got %v", iters, len(data))
	}
}
