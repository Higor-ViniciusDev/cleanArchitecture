package events

import (
	"errors"
	"sync"
)

type EventoDisparador struct {
	handlers map[string][]EventoHandlerInterface
}

func NewEventoDisparador() *EventoDisparador {
	return &EventoDisparador{
		handlers: make(map[string][]EventoHandlerInterface),
	}
}

func (ev *EventoDisparador) Disparador(event EventoInterface) error {
	if handlers, ok := ev.handlers[event.GetNome()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ev *EventoDisparador) RegistrarHandler(eventoNome string, handler EventoHandlerInterface) error {
	if _, ok := ev.handlers[eventoNome]; ok {
		for _, h := range ev.handlers[eventoNome] {
			if h == handler {
				return errors.New("Handler já registrado")
			}
		}
	}

	ev.handlers[eventoNome] = append(ev.handlers[eventoNome], handler)
	return nil
}

func (ev *EventoDisparador) HasHandlers(eventoNome string, handler EventoHandlerInterface) bool {
	if _, ok := ev.handlers[eventoNome]; ok {
		for _, h := range ev.handlers[eventoNome] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (ev *EventoDisparador) Clear() error {
	ev.handlers = make(map[string][]EventoHandlerInterface)

	return nil
}

func (ev *EventoDisparador) Remove(eventoNome string, handler EventoHandlerInterface) error {
	if _, ok := ev.handlers[eventoNome]; ok {
		for i, h := range ev.handlers[eventoNome] {
			if h == handler {
				ev.handlers[eventoNome] = append(ev.handlers[eventoNome][:i], ev.handlers[eventoNome][i+1:]...)
			}
		}
	}

	return nil
}
