package v419

import _ "embed"

var (
	//go:embed assets/required_item_list.json
	requiredItemList []byte
	//go:embed assets/item_runtime_ids.nbt
	itemRuntimeIDData []byte
	//go:embed assets/block_states.nbt
	blockStateData []byte
	//go:embed assets/entity_identifiers.nbt
	entityIdentifierData []byte
)

const (
	// ProtocolID represents the version id of the protocol
	ProtocolID      = 419
	ProtocolVersion = "1.16.100"
)

const (
	// ItemVersion represents the item version used for upgrading/downgrading.
	ItemVersion = 21
	// BlockVersion is the version of blocks (states) of the game. This version is composed
	// of 4 bytes indicating a version, interpreted as a big endian int. The current version represents
	// 1.19.70.15 {1, 19, 70, 15}.
	BlockVersion int32 = (1 << 24) | (12 << 16) | (0 << 8) | 1
)
