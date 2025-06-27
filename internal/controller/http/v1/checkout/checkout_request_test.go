package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationOrderRequest(t *testing.T) {
	var (
		request = OrderRequest{
			Items: []string{
				"43N23P",
			},
		}
	)

	t.Run("valid items", func(t *testing.T) {
		err := request.validationOrderRequest()
		assert.NoError(t, err)
	})

	t.Run("invalid items", func(t *testing.T) {
		request = OrderRequest{Items: []string{}}
		err := request.validationOrderRequest()
		assert.Error(t, err)
		assert.Equal(t, "order item can't empty", err.Error())
	})

}