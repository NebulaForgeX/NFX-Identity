package domain_verifications

import (
	"time"

	"github.com/google/uuid"
)

type VerificationMethod string

const (
	VerificationMethodDNS  VerificationMethod = "DNS"
	VerificationMethodTXT  VerificationMethod = "TXT"
	VerificationMethodHTML VerificationMethod = "HTML"
	VerificationMethodFILE VerificationMethod = "FILE"
)

type VerificationStatus string

const (
	VerificationStatusPending VerificationStatus = "PENDING"
	VerificationStatusVerified VerificationStatus = "VERIFIED"
	VerificationStatusFailed  VerificationStatus = "FAILED"
	VerificationStatusExpired VerificationStatus = "EXPIRED"
)

type DomainVerification struct {
	state DomainVerificationState
}

type DomainVerificationState struct {
	ID                 uuid.UUID
	TenantID           uuid.UUID
	Domain             string
	VerificationMethod VerificationMethod
	VerificationToken  *string
	Status             VerificationStatus
	VerifiedAt         *time.Time
	ExpiresAt          *time.Time
	CreatedAt          time.Time
	CreatedBy          *uuid.UUID
	Metadata           map[string]interface{}
}

func (dv *DomainVerification) ID() uuid.UUID                   { return dv.state.ID }
func (dv *DomainVerification) TenantID() uuid.UUID              { return dv.state.TenantID }
func (dv *DomainVerification) Domain() string                  { return dv.state.Domain }
func (dv *DomainVerification) VerificationMethod() VerificationMethod { return dv.state.VerificationMethod }
func (dv *DomainVerification) VerificationToken() *string      { return dv.state.VerificationToken }
func (dv *DomainVerification) Status() VerificationStatus       { return dv.state.Status }
func (dv *DomainVerification) VerifiedAt() *time.Time          { return dv.state.VerifiedAt }
func (dv *DomainVerification) ExpiresAt() *time.Time           { return dv.state.ExpiresAt }
func (dv *DomainVerification) CreatedAt() time.Time             { return dv.state.CreatedAt }
func (dv *DomainVerification) CreatedBy() *uuid.UUID            { return dv.state.CreatedBy }
func (dv *DomainVerification) Metadata() map[string]interface{} { return dv.state.Metadata }
