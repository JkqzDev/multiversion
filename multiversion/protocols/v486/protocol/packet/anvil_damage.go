package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AnvilDamage is sent by the client to request the dealing damage to an anvil. This packet is completely
// pointless and the server should never listen to it.
type AnvilDamage struct {
	Damage        uint8
	AnvilPosition protocol.BlockPos
}

// ID ...
func (*AnvilDamage) ID() uint32 {
	return IDAnvilDamage
}

// Marshal ...
func (pk *AnvilDamage) Marshal(io protocol.IO) {
	io.Uint8(&pk.Damage)
	io.UBlockPos(&pk.AnvilPosition)
}
