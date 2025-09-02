package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HandlerInfo struct {
	Method  string
	Handler http.HandlerFunc
}

type WebServer struct {
	Rotas        chi.Router
	Handlers     map[string]HandlerInfo // path -> HandlerInfo (method + handler)
	WebPortStart string
}

func NewWebServer(PortalWeb string) *WebServer {
	return &WebServer{
		Rotas:        chi.NewRouter(),
		Handlers:     make(map[string]HandlerInfo),
		WebPortStart: PortalWeb,
	}
}

// Registrar Handlers no Router e path
// Basicamente o metodo vai pegar um função criada no handlers ordens e registrar no path x passado por parametro
func (o *WebServer) AdicionarHandle(RotaWeb string, handlerFunc http.HandlerFunc, method string) {
	o.Handlers[RotaWeb] = HandlerInfo{
		Method:  method,
		Handler: handlerFunc,
	}
}

// Metodo que inicia o server na porta configura quando inciado a instancia da struct WebServer
func (s *WebServer) StartWebServer() {
	s.Rotas.Use(middleware.Logger)
	for path, info := range s.Handlers {
		s.Rotas.Method(info.Method, path, info.Handler)
	}

	fmt.Println("Servidor rodando na porta ", s.WebPortStart)
	http.ListenAndServe(s.WebPortStart, s.Rotas)
}
