package metadata

import (
	"sync/atomic"
	"time"
)

type PacketStatType string

type IPacketMeta interface {
	Add(value uint64)
	Increment()
	GetValue() uint64
	GetLastActivity() time.Time
	UpdateData(data PacketItem)
	GetData() []PacketItem
	Reset()
}

type PacketItem struct {
	Data interface{} `json:"data"`
	Ts   int64       `json:"timestamp"`
}

type IPacketData interface {
	Push(data PacketItem)
	Reset()
	Get() []PacketItem
}

type PacketMeta struct {
	Value uint64      `json:"value"`
	Data  IPacketData `json:"data"`
	Last  int64       `json:"lastActivity"`
}

func (p *PacketMeta) Add(value uint64) {
	atomic.AddUint64(&p.Value, value)
	atomic.StoreInt64(&p.Last, time.Now().Unix())
}

func (p *PacketMeta) Increment() {
	atomic.AddUint64(&p.Value, 1)
	atomic.StoreInt64(&p.Last, time.Now().Unix())
}

func (p *PacketMeta) GetValue() uint64 {
	return atomic.LoadUint64(&p.Value)
}

func (p *PacketMeta) GetLastActivity() time.Time {
	last := atomic.LoadInt64(&p.Last)
	return time.Unix(last, 0)
}

func (p *PacketMeta) Reset() {
	atomic.StoreUint64(&p.Value, 0)
	atomic.StoreInt64(&p.Last, time.Now().Unix())
	p.Data.Reset()
}

func (p *PacketMeta) UpdateData(data PacketItem) {
	p.Data.Push(data)
}

func (p *PacketMeta) GetData() []PacketItem {
	return p.Data.Get()
}

func NewPacketMeta(name PacketStatType, data IPacketData) IPacketMeta {
	return &PacketMeta{
		Value: 0,
		Data:  data,
		Last:  0,
	}
}
