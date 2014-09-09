package circuitbreaker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newCircuit() *Circuit {
	return NewCircuit(3, 1*time.Second)
}

func TestOpenWhenItHasTheNumberOfConfiguredErrors(t *testing.T) {
	circuit := newCircuit()
	circuit.Error()
	circuit.Error()
	circuit.Error()
	assert.True(t, circuit.IsOpen())
}

func TestClosedAfterAnOKWhenItWasOpen(t *testing.T) {
	circuit := newCircuit()
	circuit.Error()
	circuit.Error()
	circuit.Error()
	waitUntilReset()
	circuit.Ok()
	assert.True(t, circuit.IsClosed())
}

func TestClosedWhenItIsOpenAndResetTimeHasExpired(t *testing.T) {
	circuit := newCircuit()
	circuit.Error()
	circuit.Error()
	circuit.Error()
	waitUntilReset()
	circuit.Ok()
	assert.True(t, circuit.IsClosed())
}

func TestClosedWhenTheCircuitHasBeenResetAsTimeExpired(t *testing.T) {
	circuit := newCircuit()
	circuit.Error()
	circuit.Error()
	waitUntilReset()
	circuit.Error()
	assert.True(t, circuit.IsClosed())
}

func waitUntilReset() {
	time.Sleep(1 * time.Second)
}