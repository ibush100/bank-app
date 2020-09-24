package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {

	CreateTransaction("fresh@example.com", "email@example.com", 100)
	payee, payor := FindPayeeAndPayor("fresh@example.com", "email@example.com")
	assert.Equal(t, payee.Balence, 200, "that didn't work")
	assert.Equal(t, payor.Balence, 100, "that didn't work")

}
