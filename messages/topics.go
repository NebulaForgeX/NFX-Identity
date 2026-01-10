package messages

import "nfxid/pkgs/rabbitmqx/messaging"

const (
	// =============== Access ===============
	MKAccess messaging.MessageKey = "access"

	// =============== Clients ===============
	MKClients messaging.MessageKey = "clients"

	// =============== Directory ===============
	MKDirectory messaging.MessageKey = "directory"

	// =============== Audit ===============
	MKAudit messaging.MessageKey = "audit"

	// =============== Auth ===============
	MKAuth messaging.MessageKey = "auth"

	// =============== Tenants ===============
	MKTenants messaging.MessageKey = "tenants"

	// =============== Image ===============
	MKImage messaging.MessageKey = "image"

	// =============== System ===============
	MKSystem messaging.MessageKey = "system"
)

// AccessTopic 提供 Access 模块的 MessageKey
// 嵌入此结构体后，消息自动实现 RoutingKey() 方法
type AccessTopic struct{}

func (AccessTopic) RoutingKey() messaging.MessageKey { return MKAccess }

// ClientsTopic 提供 Clients 模块的 MessageKey
type ClientsTopic struct{}

func (ClientsTopic) RoutingKey() messaging.MessageKey { return MKClients }

// DirectoryTopic 提供 Directory 模块的 MessageKey
type DirectoryTopic struct{}

func (DirectoryTopic) RoutingKey() messaging.MessageKey { return MKDirectory }

// AuditTopic 提供 Audit 模块的 MessageKey
type AuditTopic struct{}

func (AuditTopic) RoutingKey() messaging.MessageKey { return MKAudit }

// AuthTopic 提供 Auth 模块的 MessageKey
type AuthTopic struct{}

func (AuthTopic) RoutingKey() messaging.MessageKey { return MKAuth }

// TenantsTopic 提供 Tenants 模块的 MessageKey
type TenantsTopic struct{}

func (TenantsTopic) RoutingKey() messaging.MessageKey { return MKTenants }

// ImageTopic 提供 Image 模块的 MessageKey
type ImageTopic struct{}

func (ImageTopic) RoutingKey() messaging.MessageKey { return MKImage }

// SystemTopic 提供 System 模块的 MessageKey
type SystemTopic struct{}

func (SystemTopic) RoutingKey() messaging.MessageKey { return MKSystem }
