package requestCTX

// Reprensenta dados fundamentais sobre o actor da request
type Metadata struct {
	RequestID, ClientIP, TenantID string
}

// Retorna uma cópia do (Context.Metadata)
// Faz uso de sync.RWMutex (RLock() e RUnlock())
func (c *Context) Metadata() Metadata {
	if c == nil {
		return Metadata{}
	}

	c.stateMu.RLock()
	defer c.stateMu.RUnlock()

	return Metadata{
		RequestID: c.metadata.RequestID,
		ClientIP:  c.metadata.ClientIP,
		TenantID:  c.metadata.TenantID,
	}

}

// Setters

// Define o valor da RequestID do (Context.Metadata)
// Faz uso de sync.RWMutex (RLock() e RUnlock())
func (c *Context) SetRequestID(id string) {
	if c == nil {
		return
	}

	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	c.metadata.RequestID = id
}

// Define o valor da ClientIP do (Context.Metadata)
// Faz uso de sync.RWMutex (RLock() e RUnlock())
func (c *Context) SetClientIP(ip string) {
	if c == nil {
		return
	}

	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	c.metadata.ClientIP = ip
}

// Define o valor da enantID do (Context.Metadata)
// Faz uso de sync.RWMutex (RLock() e RUnlock())
func (c *Context) SetTenantID(id string) {
	if c == nil {
		return
	}

	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	c.metadata.TenantID = id
}

// Options

func WithRequestID(id string) Option {
	return func(c *Context) {
		c.SetRequestID(id)
	}
}

func WithClientIP(ip string) Option {
	return func(c *Context) {
		c.SetClientIP(ip)
	}
}

func WithTenantID(id string) Option {
	return func(c *Context) {
		c.SetTenantID(id)
	}
}
