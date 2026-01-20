package rabbitmqx

import (
	"fmt"
)

func (c *Config) Validate() error {
	// 如果 URI 为空，则必须提供分开的配置字段
	if c.URI == "" {
		if c.Host == "" {
			return fmt.Errorf("rabbitmq host is empty (either uri or host must be provided)")
		}
		// 其他字段有默认值，不需要验证
	}

	// 尝试构建 URI 以验证配置
	_, err := c.BuildURI()
	if err != nil {
		return fmt.Errorf("failed to build rabbitmq uri: %w", err)
	}

	if c.ClientID == "" {
		return fmt.Errorf("rabbitmq client_id is empty")
	}

	// 验证交换机配置
	if c.Exchange.Type != "" {
		if !c.Exchange.Type.IsValid() {
			return fmt.Errorf("rabbitmq exchange.type is invalid: %s", c.Exchange.Type)
		}
	}

	// 验证 ProducerRouting 中的 Exchange 类型
	for key, routing := range c.ProducerExchanges {
		if routing.Type != "" && !routing.Type.IsValid() {
			return fmt.Errorf("rabbitmq producer_exchanges[%s].type is invalid: %s", key, routing.Type)
		}
	}

	// 验证 DeliveryMode
	if c.Producer.DeliveryMode != 0 && c.Producer.DeliveryMode != 1 && c.Producer.DeliveryMode != 2 {
		return fmt.Errorf("rabbitmq producer.delivery_mode must be 1 (non-persistent) or 2 (persistent), got: %d", c.Producer.DeliveryMode)
	}

	return nil
}
