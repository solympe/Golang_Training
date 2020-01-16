package proxy

// Proxy ...
type Proxy interface {
	Get() (mySpecificData string)
	Set(mySpecificData string)
}

type proxy struct {
	cache    Proxy
	dataBase Proxy
}

// Set ...
func (p *proxy) Set(mySpecificData string) {
	p.dataBase.Set(mySpecificData)
	p.cache.Set(mySpecificData)
}

// Get returns data from cacheц
func (p *proxy) Get() (mySpecificData string) {
	mySpecificData = p.cache.Get()
	if mySpecificData == "" {
		mySpecificData = p.dataBase.Get()
	}
	return
}

// NewProxy returns new instance of proxy
func NewProxy(dataBase Proxy, cache Proxy) Proxy {
	return &proxy{
		dataBase: dataBase,
		cache:    cache,
	}
}
