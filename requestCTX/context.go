// Pacote responsável por fornecer o contexto de execução de uma requisição no Guarita.
package requestCTX

import (
	"context"
	"errors"
	"net/http"
	"sync"
)

// Representa o contexto de execução de uma requisição dentro do Guaraná.
// Implementa o context.Context e disponibiliza todos os seus métodos.
// Foi pensado para acompanha uma request integralmente em seu caminho.
// Não é um acessório, é o componente FUNDAMENTAL do Guaraná.
type Context struct {
	context.Context

	request *http.Request

	stateMu   sync.RWMutex
	metadata  Metadata
	principal *Principal

	values valueStore
}

type Option func(*Context)

// Cria uma nova instacia de Context
// req não pode ser nil
func NewCTX(req *http.Request, opts ...Option) (*Context, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}

	reqCtx := req.Context()

	if reqCtx == nil {
		reqCtx = context.Background()
	}

	ctx := &Context{
		Context: reqCtx,
		request: req.Clone(reqCtx),
		values:  newValueStore(),
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(ctx)
	}

	return ctx, nil
}

// Retorna a requisição associada ao Context
func (c *Context) Request() *http.Request {
	if c.request == nil {
		return nil
	}
	return c.request
}

func (c *Context) HTTPMethod() string {
	if c == nil {
		return ""
	}
	return c.request.Method
}

func (c *Context) HTTPPath() string {
	if c == nil || c.request == nil || c.request.URL == nil {
		return ""
	}
	return c.request.URL.Path
}

func (c *Context) HTTPHeader() http.Header {
	if c == nil || c.request == nil {
		return nil
	}
	return c.request.Header
}
