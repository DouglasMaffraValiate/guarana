package requestCTX

import "strings"

// Representa a entidade autenticada na requisição
type Principal struct {
	PrincipalID string
	Roles       []string
}

// Seta a identidade entidade autenticada da requisição (Principal)
// Se o parâmetro for nil não faz nada
// Faz uso de sync.RWMutex  Lock() e Unlock()
func (c *Context) SetPrincipal(principal *Principal) {
	if c == nil {
		return
	}

	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	if principal == nil {
		c.principal = nil
		return
	}

	c.principal = &Principal{
		PrincipalID: principal.PrincipalID,
		Roles:       principal.Roles,
	}
}

// Retorna o endereço de memória de uma cópia do Principal
// Se o parâmetro for nil não faz nada
// Faz uso de sync.RWMutex  RLock() e RUnlock()
func (c *Context) Principal() (*Principal, bool) {
	if c == nil {
		return nil, false
	}

	c.stateMu.RLock()
	defer c.stateMu.RUnlock()

	if c.principal == nil {
		return nil, false
	}

	return &Principal{
		PrincipalID: c.principal.PrincipalID,
		Roles:       append([]string(nil), c.principal.Roles...),
	}, true
}

// Retorna se existe uma identidade autenticada (Principal)
// Faz uso de sync.RWMutex RLock() e RUnlock()
func (c *Context) IsAuthenticated() bool {
	if c == nil {
		return false
	}

	c.stateMu.RLock()
	defer c.stateMu.RUnlock()

	return c.principal != nil
}

// Retorna se o Principal da request tem a role inserida. case-sensitive
// Faz uso de sync.RWMutex RLock() e RUnlock()
func (c *Context) ThisPrincipalHasThisRole(role string) bool {
	if c == nil {
		return false
	}

	c.stateMu.RLock()
	defer c.stateMu.RUnlock()

	for _, thisRole := range c.principal.Roles {
		if strings.Compare(thisRole, role) == 0 {
			return true
		}
	}

	return false
}

// Remove a identidade autenticada. da request
// Faz uso de sync.RWMutex Lock() e Unlock()
func (c *Context) ClearPrincipal() {
	if c == nil {
		return
	}
	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	c.principal = nil
}

func WithPrincipal(p *Principal) Option {
	return func(c *Context) {
		c.SetPrincipal(p)
	}
}
