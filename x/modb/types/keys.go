package types

const (
	// ModuleName defines the module name
	ModuleName = "modb"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_modb"
)

var (
	ParamsKey = []byte("p_modb")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
