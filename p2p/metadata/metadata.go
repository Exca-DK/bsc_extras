package metadata

import (
	"time"

	"github.com/ethereum/go-ethereum/log"
)

type IMetadata interface {
	ConnectionTime() uint64
	IncrementPacket(stat PacketStatType)
	UpdatePacketData(stat PacketStatType, data PacketItem)
	GetPacket(stat PacketStatType) (IPacketMeta, PacketError)
	GetPackets() []IPacketMeta
	RegisterPacket(stat PacketStatType, packet IPacketMeta) PacketError
}

type Metadata struct {
	ConnectedAt time.Time                      `json:"connectedAt"`
	Packets     map[PacketStatType]IPacketMeta `json:"packets"`
}

func (p *Metadata) ConnectionTime() uint64 {
	return uint64(time.Since(p.ConnectedAt).Seconds())
}

func (p *Metadata) IncrementPacket(stat PacketStatType) {
	packet, err := p.GetPacket(stat)
	if err != nil {
		log.Warn(err.Error(), "packet", stat)
		return
	}
	packet.Increment()
}

func (p *Metadata) UpdatePacketData(stat PacketStatType, data PacketItem) {
	packet, err := p.GetPacket(stat)
	if err != nil {
		log.Warn(err.Error(), "packet", stat)
		return
	}
	packet.UpdateData(data)
}

func (p *Metadata) GetPacket(stat PacketStatType) (IPacketMeta, PacketError) {
	packet, ok := p.Packets[stat]
	if !ok {
		return nil, ErrPacketNotFound
	}
	return packet, nil
}

func (p *Metadata) GetPackets() []IPacketMeta {
	packets := []IPacketMeta{}
	for _, packet := range p.Packets {
		packets = append(packets, packet)
	}
	return packets
}

func (p *Metadata) RegisterPacket(stat PacketStatType, packet IPacketMeta) PacketError {
	_, err := p.GetPacket(stat)
	if err == nil {
		return ErrPackedRegistered
	}
	p.Packets[stat] = packet
	return nil
}

func NewMetadata() IMetadata {
	return &Metadata{
		ConnectedAt: time.Now(),
		Packets:     make(map[PacketStatType]IPacketMeta),
	}
}
