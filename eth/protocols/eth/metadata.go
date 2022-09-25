package eth

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/metadata"
)

var (
	NewBlockHashesPacketStat             metadata.PacketStatType = "NewBlockHashesPacket"
	NewBlockPacketStat                   metadata.PacketStatType = "NewBlockPacket"
	NewPooledTransactionHashesPacketStat metadata.PacketStatType = "NewPooledTransactionHashesPacket"
	TransactionsPacketStat               metadata.PacketStatType = "TransactionsPacket"
	PooledTransactionsPacketStat         metadata.PacketStatType = "PooledTransactionsPacket"

	mapping map[int]metadata.PacketStatType = map[int]metadata.PacketStatType{
		NewBlockHashesMsg:             NewBlockHashesPacketStat,
		NewBlockMsg:                   NewBlockPacketStat,
		NewPooledTransactionHashesMsg: NewPooledTransactionHashesPacketStat,
		TransactionsMsg:               TransactionsPacketStat,
		PooledTransactionsMsg:         PooledTransactionsPacketStat,
	}
)

func MsgToStat(msgType int) (metadata.PacketStatType, metadata.PacketError) {
	stat, ok := mapping[msgType]
	if !ok {
		return "", metadata.ErrPacketNotFound
	}
	return stat, nil
}

type EthPacketData struct {
	Cache []metadata.PacketItem `json:"cache"`
	size  int                   `json:"-"`
	mu    *sync.RWMutex         `json:"-"`
	Last  int64                 `json:"lastActivity"`
}

func (d *EthPacketData) Push(data metadata.PacketItem) {
	d.mu.Lock()
	defer d.mu.Unlock()

	switch v := data.Data.(type) {
	case []common.Hash:
		for _, hash := range v {
			d.Cache = append(d.Cache, metadata.PacketItem{Data: hash, Ts: data.Ts})
		}
	case common.Hash:
		d.Cache = append(d.Cache, data)
	}

	length := len(d.Cache)
	if length > d.size {
		d.Cache = d.Cache[length-d.size:]
	}
}

func (d *EthPacketData) Reset() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Cache = make([]metadata.PacketItem, d.size)
	d.Last = time.Now().Unix()
}

func (d *EthPacketData) Get() []metadata.PacketItem {
	d.mu.RLock()
	defer d.mu.RUnlock()

	cache := make([]metadata.PacketItem, 0, len(d.Cache))
	cache = append(cache, d.Cache...)
	return cache
}

func NewPacketMetadata(size int, _type metadata.PacketStatType) metadata.IPacketMeta {
	data := EthPacketData{
		Cache: make([]metadata.PacketItem, 0, size),
		size:  size,
		mu:    &sync.RWMutex{},
		Last:  0,
	}
	return metadata.NewPacketMeta(_type, &data)
}

func NewEthMetadata(size uint64) metadata.IMetadata {
	meta := metadata.NewMetadata()

	for _, pType := range mapping {
		packet := NewPacketMetadata(int(size), pType)
		err := meta.RegisterPacket(pType, packet)
		if err != nil {
			log.Warn(err.Error(), "packet", pType)
		}
	}

	return meta
}
