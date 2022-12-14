package types

const (
	// ModuleName defines the module name
	ModuleName = "atomicchocolate"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_atomicchocolate"

	// Keep track of the index of posts
	ProjectKey      = "Project/value/"
	ProjectCountKey = "Project/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
