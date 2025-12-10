package eventbus

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type EventRouterConfig struct {
	CloseTimeout         time.Duration
	Logger               watermill.LoggerAdapter
	EventTypeMetadataKey string
}

func (c *EventRouterConfig) setDefaults() {
	if c.Logger == nil {
		c.Logger = watermill.NopLogger{}
	}
	if c.CloseTimeout == 0 {
		c.CloseTimeout = 10 * time.Second
	}
	if c.EventTypeMetadataKey == "" {
		c.EventTypeMetadataKey = DefaultEventTypeHeaderKey
	}
}

type EventRouter struct {
	*message.Router
	sub             *BusSubscriber
	cfg             *EventRouterConfig
	handlersByTopic map[string]map[string]message.NoPublishHandlerFunc
	muxInstalled    map[string]struct{}
}

func NewEventRouter(sub *BusSubscriber, cfg EventRouterConfig) (*EventRouter, error) {
	cfg.setDefaults()
	r, err := message.NewRouter(message.RouterConfig{CloseTimeout: cfg.CloseTimeout}, cfg.Logger)
	if err != nil {
		return nil, err
	}
	return &EventRouter{
		Router:          r,
		sub:             sub,
		cfg:             &cfg,
		handlersByTopic: make(map[string]map[string]message.NoPublishHandlerFunc),
		muxInstalled:    make(map[string]struct{}),
	}, nil
}

func RegisterHandler[T TypedEvent](
	er *EventRouter,
	handler EventHandler[T],
) {
	// Make sure the topic is bound to the event
	topicName, ok := er.sub.GetTopic(TopicKeyOf[T]())
	if !ok || topicName == "" {
		var zero T
		panic(fmt.Sprintf("eventbus: topic key not found for event type %T", zero))
	}

	eventKey := EventTypeOf[T]()

	if er.handlersByTopic[topicName] == nil {
		er.handlersByTopic[topicName] = make(map[string]message.NoPublishHandlerFunc)
	}
	if _, exists := er.handlersByTopic[topicName][eventKey]; exists {
		panic(fmt.Sprintf("eventbus: duplicate handler for %s on topic %s", eventKey, topicName))
	}

	er.handlersByTopic[topicName][eventKey] = wrapHandler(handler)
	er.ensureMux(topicName)
}

func (er *EventRouter) ensureMux(topic string) {
	if _, ok := er.muxInstalled[topic]; ok {
		return
	}
	handlerName := fmt.Sprintf("%s__domain_mux", topic)
	er.Router.AddConsumerHandler(
		handlerName,
		topic,
		er.sub,
		func(msg *message.Message) error {
			et := msg.Metadata.Get(er.cfg.EventTypeMetadataKey)
			if et == "" {
				er.Router.Logger().Error(fmt.Sprintf("missing %s header", er.cfg.EventTypeMetadataKey), nil, watermill.LogFields{
					"topic": topic, "uuid": msg.UUID,
				})
				return nil
			}
			m := er.handlersByTopic[topic]
			if m == nil {
				er.Router.Logger().Error(fmt.Sprintf("no handler map for topic %s", topic), nil, watermill.LogFields{
					"topic": topic, er.cfg.EventTypeMetadataKey: et, "uuid": msg.UUID,
				})
				return nil
			}
			h, ok := m[et]
			if !ok {
				er.Router.Logger().Info(fmt.Sprintf("skip unknown %s", er.cfg.EventTypeMetadataKey), watermill.LogFields{
					"topic": topic, er.cfg.EventTypeMetadataKey: et, "uuid": msg.UUID,
				})
				return nil
			}
			return h(msg)
		},
	)
	er.muxInstalled[topic] = struct{}{}
}

func wrapHandler[T TypedEvent](evtHandler EventHandler[T]) message.NoPublishHandlerFunc {
	return func(msg *message.Message) error {
		var evt T
		if err := json.Unmarshal(msg.Payload, &evt); err != nil {
			return fmt.Errorf("eventbus: json unmarshal failed: %w", err)
		}

		if v, ok := any(&evt).(Validatable); ok {
			if err := v.Validate(); err != nil {
				return fmt.Errorf("eventbus: validation failed: %w", err)
			}
		}

		// Pass context, for tracing / timeout control
		ctx := msg.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		return evtHandler(ctx, evt, msg)
	}
}
