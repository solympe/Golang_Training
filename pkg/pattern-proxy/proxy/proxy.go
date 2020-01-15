package proxy

type Proxy interface {
	GetData() string
	SendData(data string)
}

type proxy struct {
	cache    Proxy
	dataBase Proxy
}

// SendData updates data in the main database and cache
func (p *proxy) SendData(data string) {
	Proxy.SendData(p.cache, data)
	Proxy.SendData(p.dataBase, data)
}

// GetData returns 'fresh' data from cache
func (p proxy) GetData() string {
	freshData := Proxy.GetData(p.cache)
	return freshData
}

// NewProxy returns new instance of proxy
func NewProxy(cache Proxy, db Proxy) Proxy {
	return &proxy{cache, db}
}
