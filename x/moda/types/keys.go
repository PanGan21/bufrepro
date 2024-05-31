package types

const (
	// ModuleName defines the module name
	ModuleName = "moda"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_moda"
)

var (
	ParamsKey = []byte("p_moda")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
