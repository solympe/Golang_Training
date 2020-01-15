package proxy

// Proxy ...
type Proxy interface {
	GetData() string
	Send(data string)
}

type proxy struct {
	cache    Proxy
	dataBase Proxy
}

// Send updates data in the main database and cache
func (p *proxy) Send(data string) {
	Proxy.Send(p.cache, data)
	Proxy.Send(p.dataBase, data)
}

// GetData returns 'fresh' data from cache
func (p *proxy) GetData() string {
	freshData := Proxy.GetData(p.cache)
	return freshData
}

// NewProxy returns new instance of proxy
func NewProxy(cache Proxy, db Proxy) Proxy {
	return &proxy{cache, db}
}
