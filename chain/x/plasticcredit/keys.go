package plasticcredit

var (
	ParamsKey     = []byte{0x00}
	IDCountersKey = []byte{0x01}
	IssuerKey     = []byte{0x02}
	ApplicantKey  = []byte{0x03}
)

const (
	// ModuleName defines the module name
	ModuleName = "plasticcredit"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_plasticcredit"
)