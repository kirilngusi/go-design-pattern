package strategy

type Cache struct {
	storage map[string]string
	evictionAlgo EvictionAlgo
	capacity int
	maxCapacity int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make([])
}
