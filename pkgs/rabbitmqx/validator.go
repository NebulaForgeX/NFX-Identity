package rabbitmqx

import (
	"fmt"
)

func (c *Config) Validate() error {
	if c.URI == "" {
		return fmt.Errorf("rabbitmq uri is empty")
	}

	if c.ClientID == "" {
		return fmt.Errorf("rabbitmq client_id is empty")
	}

	// 验证交换机配置
	if c.Exchange.Type != "" {
		validTypes := map[string]bool{
			"direct":  true,
			"topic":   true,
			"fanout":  true,
			"headers": true,
		}
		if !validTypes[c.Exchange.Type] {
			return fmt.Errorf("rabbitmq exchange.type must be one of: direct, topic, fanout, headers, got: %s", c.Exchange.Type)
		}
	}

	// 验证 DeliveryMode
	if c.Producer.DeliveryMode != 0 && c.Producer.DeliveryMode != 1 && c.Producer.DeliveryMode != 2 {
		return fmt.Errorf("rabbitmq producer.delivery_mode must be 1 (non-persistent) or 2 (persistent), got: %d", c.Producer.DeliveryMode)
	}

	return nil
}
