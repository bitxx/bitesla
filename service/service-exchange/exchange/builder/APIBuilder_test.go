package builder

import (
	"github.com/jason-wj/bitesla/service/service-exchange/exchange"
	"github.com/stretchr/testify/assert"
	"testing"
)

var builder = NewAPIBuilder()

func TestAPIBuilder_Build(t *testing.T) {
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(exchange.OkcoinCn), exchange.OkcoinCn)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(exchange.HuobiPro).GetExchangeName(), exchange.HuobiPro)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(exchange.ZB).GetExchangeName(), exchange.ZB)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(exchange.BIGONE).GetExchangeName(), exchange.BIGONE)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(exchange.OKEX).GetExchangeName(), exchange.OKEX)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(exchange.POLONIEX).GetExchangeName(), exchange.POLONIEX)
	assert.Equal(t, builder.APIKey("").APISecretkey("").Build(exchange.KRAKEN).GetExchangeName(), exchange.KRAKEN)
}
