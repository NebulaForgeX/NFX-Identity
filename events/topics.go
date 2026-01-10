package events

import "nfxid/pkgs/kafkax/eventbus"

const (
	// =============== Access ===============
	TKAccess    eventbus.TopicKey = "access"
	TKAccessDLQ eventbus.TopicKey = "access_poison"

	// =============== Clients ===============
	TKClients    eventbus.TopicKey = "clients"
	TKClientsDLQ eventbus.TopicKey = "clients_poison"

	// =============== Directory ===============
	TKDirectory    eventbus.TopicKey = "directory"
	TKDirectoryDLQ eventbus.TopicKey = "directory_poison"

	// =============== Audit ===============
	TKAudit    eventbus.TopicKey = "audit"
	TKAuditDLQ eventbus.TopicKey = "audit_poison"

	// =============== Auth ===============
	TKAuth    eventbus.TopicKey = "auth"
	TKAuthDLQ eventbus.TopicKey = "auth_poison"

	// =============== Tenants ===============
	TKTenants    eventbus.TopicKey = "tenants"
	TKTenantsDLQ eventbus.TopicKey = "tenants_poison"

	// =============== Image ===============
	TKImage    eventbus.TopicKey = "image"
	TKImageDLQ eventbus.TopicKey = "image_poison"

	// =============== System ===============
	TKSystem    eventbus.TopicKey = "system"
	TKSystemDLQ eventbus.TopicKey = "system_poison"
)

// AccessTopic 提供 Access 模块的 TopicKey
// 嵌入此结构体后，事件自动实现 TopicKey() 方法
type AccessTopic struct{}

func (AccessTopic) TopicKey() eventbus.TopicKey { return TKAccess }

// ClientsTopic 提供 Clients 模块的 TopicKey
type ClientsTopic struct{}

func (ClientsTopic) TopicKey() eventbus.TopicKey { return TKClients }

// DirectoryTopic 提供 Directory 模块的 TopicKey
type DirectoryTopic struct{}

func (DirectoryTopic) TopicKey() eventbus.TopicKey { return TKDirectory }

// AuditTopic 提供 Audit 模块的 TopicKey
type AuditTopic struct{}

func (AuditTopic) TopicKey() eventbus.TopicKey { return TKAudit }

// AuthTopic 提供 Auth 模块的 TopicKey
type AuthTopic struct{}

func (AuthTopic) TopicKey() eventbus.TopicKey { return TKAuth }

// TenantsTopic 提供 Tenants 模块的 TopicKey
type TenantsTopic struct{}

func (TenantsTopic) TopicKey() eventbus.TopicKey { return TKTenants }

// ImageTopic 提供 Image 模块的 TopicKey
type ImageTopic struct{}

func (ImageTopic) TopicKey() eventbus.TopicKey { return TKImage }

// SystemTopic 提供 System 模块的 TopicKey
type SystemTopic struct{}

func (SystemTopic) TopicKey() eventbus.TopicKey { return TKSystem }
