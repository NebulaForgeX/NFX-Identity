package messaging

import (
	"reflect"
	"strings"
	"sync"
	"unicode"
)

// messageTypeCache 缓存类型名到 MessageType 的映射
// key: 类型完整名称 (包含包路径)
// value: 生成的 MessageType
var messageTypeCache sync.Map

// AutoMessageType 从类型名自动生成 MessageType
// 使用反射获取类型名，然后转换为 MessageType 格式
// 结果会被缓存，避免重复计算
// MessageType 只包含消息类型本身，不包含 messageKey（路由通过 RoutingKey() 方法处理）
// 例如: "GrantsInvalidateCacheMessage" -> "grants_invalidate_cache"
//
// 命名规则:
//  1. 移除 "Message" 或 "Event" 后缀（兼容两种命名）
//  2. 将驼峰命名转换为小写下划线格式
//
// 示例:
//   - GrantsInvalidateCacheMessage -> "grants_invalidate_cache"
//   - UserCreatedMessage -> "user_created"
//   - GrantsInvalidateCacheEvent -> "grants_invalidate_cache" (兼容)
func AutoMessageType[T TypedMessage]() MessageType {
	var zero T
	t := reflect.TypeOf(zero)
	if t == nil {
		return ""
	}

	// 构建缓存键: 类型完整名称
	cacheKey := t.String()

	// 先检查缓存
	if cached, ok := messageTypeCache.Load(cacheKey); ok {
		return cached.(MessageType)
	}

	// 获取类型名（去掉包路径）
	typeName := t.Name()
	if typeName == "" {
		// 如果 Name() 为空，尝试从 String() 中提取
		fullName := t.String()
		parts := strings.Split(fullName, ".")
		if len(parts) > 0 {
			typeName = parts[len(parts)-1]
		}
	}

	// 移除 "Message" 或 "Event" 后缀（兼容两种命名）
	typeName = strings.TrimSuffix(typeName, "Message")
	typeName = strings.TrimSuffix(typeName, "Event")

	// 将驼峰命名转换为小写下划线格式
	// "GrantsInvalidateCache" -> "grants_invalidate_cache"
	var result strings.Builder
	for i, r := range typeName {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteByte('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}

	messageTypeValue := MessageType(result.String())

	// 存入缓存
	messageTypeCache.Store(cacheKey, messageTypeValue)

	return messageTypeValue
}

// AutoMessage 自动生成 MessageType 的基础消息结构体。
// 嵌入此结构体后，MessageType() 返回空字符串，MessageTypeOf 会自动使用 AutoMessageType 生成。
// 如果用户自己实现了 MessageType() 方法，会优先使用用户定义的（覆盖嵌入的方法）。
//
// 使用方式:
//
//	type GrantsInvalidateCacheMessage struct {
//	    messaging.AutoMessage
//	    ID string `json:"id"`
//	}
//
//	func (GrantsInvalidateCacheMessage) RoutingKey() MessageKey {
//	    return "access"
//	}
//
//	// MessageType() 返回空字符串，MessageTypeOf 会自动使用 AutoMessageType 生成
//	// 如果需要自定义，可以自己实现 MessageType() 方法覆盖默认行为
type AutoMessage struct {
	messageKey MessageKey
}

// MessageType 返回空字符串，表示使用自动生成。
// MessageTypeOf 函数会检测到空字符串，然后使用 AutoMessageType 自动生成。
func (m AutoMessage) MessageType() MessageType {
	return "" // 返回空字符串，触发 MessageTypeOf 使用 AutoMessageType
}

// RoutingKey 返回消息键（MessageKey）。
func (m AutoMessage) RoutingKey() MessageKey {
	return m.messageKey
}

// NewAutoMessage 创建自动生成 MessageType 的基础消息。
func NewAutoMessage(messageKey MessageKey) AutoMessage {
	return AutoMessage{
		messageKey: messageKey,
	}
}
