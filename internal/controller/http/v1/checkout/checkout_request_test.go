package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationCheckoutRequest(t *testing.T) {
	var (
		request = CheckoutRequest{
			Items: []string{
				"43N23P",
			},
		}
	)

	t.Run("valid items", func(t *testing.T) {
		err := request.validationCheckoutRequest()
		assert.NoError(t, err)
	})

	t.Run("invalid items", func(t *testing.T) {
		request = CheckoutRequest{Items: []string{}}
		err := request.validationCheckoutRequest()
		assert.Error(t, err)
		assert.Equal(t, "order item can't empty", err.Error())
	})

}