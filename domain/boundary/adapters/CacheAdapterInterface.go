package adapters

// CacheAdapterInterface is implemented by all cache adapters.
type CacheAdapterInterface interface {
	Get(key string) []byte
	Set(key string, value []byte)
	Del(key string)
	Destruct()
}
