package transaction

import "testing"

func TestCreateTransaction(t *testing.T) {

	CreateTransaction("fresh@example.com", "email@example.com", 100)

}
