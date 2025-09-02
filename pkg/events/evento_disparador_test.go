package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EventoTest struct {
	nome  string
	value any
}

func (et *EventoTest) GetNome() string {
	return et.nome
}

func (et *EventoTest) GetValues() any {
	return et.value
}

func (e *EventoTest) SetValues(value interface{}) {
	e.value = value
}

func (e *EventoTest) GetDateTime() time.Time {
	retorno := time.Now()

	return retorno
}

type TestEventoHandler struct {
	ID int
}

func (h *TestEventoHandler) Handle(event EventoInterface, wg *sync.WaitGroup) {
}

type EventoDisparadorTestSuite struct {
	suite.Suite
	event            EventoTest
	event2           EventoTest
	handler          TestEventoHandler
	handler2         TestEventoHandler
	handler3         TestEventoHandler
	EventoDisparador *EventoDisparador
}

func (suite *EventoDisparadorTestSuite) SetupTest() {
	suite.EventoDisparador = NewEventoDisparador()
	suite.handler = TestEventoHandler{
		ID: 1,
	}
	suite.handler2 = TestEventoHandler{
		ID: 2,
	}
	suite.handler3 = TestEventoHandler{
		ID: 3,
	}
	suite.event = EventoTest{nome: "test", value: "test"}
	suite.event2 = EventoTest{nome: "test2", value: "test2"}
}

func (suite *EventoDisparadorTestSuite) TestEventoDisparador_Register() {
	err := suite.EventoDisparador.RegistrarHandler(suite.event.GetNome(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventoDisparador.handlers[suite.event.GetNome()]))

	err = suite.EventoDisparador.RegistrarHandler(suite.event.GetNome(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventoDisparador.handlers[suite.event.GetNome()]))

	assert.Equal(suite.T(), &suite.handler, suite.EventoDisparador.handlers[suite.event.GetNome()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.EventoDisparador.handlers[suite.event.GetNome()][1])
}

func (suite *EventoDisparadorTestSuite) TestEventoDisparador_Registe_HandlersRepetido() {
	err := suite.EventoDisparador.RegistrarHandler(suite.event.GetNome(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventoDisparador.handlers[suite.event.GetNome()]))

	err = suite.EventoDisparador.RegistrarHandler(suite.event.GetNome(), &suite.handler)
	suite.Error(err)
	suite.Equal(1, len(suite.EventoDisparador.handlers[suite.event.GetNome()]))

	assert.Equal(suite.T(), &suite.handler, suite.EventoDisparador.handlers[suite.event.GetNome()][0])
}

func (suite *EventoDisparadorTestSuite) TestEventoDisparador_Has() {
	err := suite.EventoDisparador.RegistrarHandler(suite.event.GetNome(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventoDisparador.handlers[suite.event.GetNome()]))

	retorno := suite.EventoDisparador.HasHandlers(suite.event.GetNome(), &suite.handler)
	suite.Equal(true, retorno)

	retorno1 := suite.EventoDisparador.HasHandlers(suite.event.GetNome(), &suite.handler2)
	suite.Equal(false, retorno1)

}

func (suite *EventoDisparadorTestSuite) TestEventoDisparador_Clear() {
	// Event 1
	err := suite.EventoDisparador.RegistrarHandler(suite.event.GetNome(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventoDisparador.handlers[suite.event.GetNome()]))

	err = suite.EventoDisparador.RegistrarHandler(suite.event2.GetNome(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventoDisparador.handlers[suite.event2.GetNome()]))

	suite.EventoDisparador.Clear()
	suite.Equal(0, len(suite.EventoDisparador.handlers))
}

func (suite *EventoDisparadorTestSuite) TestEventoDisparador_Remove() {
	// Event 1
	err := suite.EventoDisparador.RegistrarHandler(suite.event.GetNome(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventoDisparador.handlers[suite.event.GetNome()]))

	err = suite.EventoDisparador.RegistrarHandler(suite.event2.GetNome(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventoDisparador.handlers[suite.event2.GetNome()]))

	suite.EventoDisparador.Remove(suite.event.GetNome(), &suite.handler)
	suite.Equal(0, len(suite.EventoDisparador.handlers[suite.event.GetNome()]))
	suite.False(suite.EventoDisparador.HasHandlers(suite.event.GetNome(), &suite.handler))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventoInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventoDisparadorTestSuite) TestEventoDisparador_Called() {
	mockHandle := &MockHandler{}
	mockHandle.On("Handle", &suite.event)

	mockHandle2 := &MockHandler{}
	mockHandle2.On("Handle", &suite.event)

	//"Registrar" o manipulador de eventos
	suite.EventoDisparador.RegistrarHandler(suite.event.GetNome(), mockHandle)
	suite.EventoDisparador.RegistrarHandler(suite.event.GetNome(), mockHandle2)

	// Disparar o evento
	suite.EventoDisparador.Disparador(&suite.event)

	// Verificar se o manipulador foi chamado
	mockHandle.AssertExpectations(suite.T())
	mockHandle2.AssertExpectations(suite.T())

	// Verificar se o número de chamadas é o esperado
	mockHandle.AssertNumberOfCalls(suite.T(), "Handle", 1)
	mockHandle2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventoDisparadorTestSuite))
}
