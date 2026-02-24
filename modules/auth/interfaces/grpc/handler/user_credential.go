package handler

import (
	"context"

	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	userCredentialAppCommands "nfxid/modules/auth/application/user_credentials/commands"
	userCredentialDomain "nfxid/modules/auth/domain/user_credentials"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/errx"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentialHandler struct {
	usercredentialpb.UnimplementedUserCredentialServiceServer
	userCredentialAppSvc *userCredentialApp.Service
}

func NewUserCredentialHandler(userCredentialAppSvc *userCredentialApp.Service) *UserCredentialHandler {
	return &UserCredentialHandler{userCredentialAppSvc: userCredentialAppSvc}
}

func (h *UserCredentialHandler) CreateUserCredential(
	ctx context.Context,
	req *usercredentialpb.CreateUserCredentialRequest,
) (*usercredentialpb.CreateUserCredentialResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	var credentialType userCredentialDomain.CredentialType
	switch req.CredentialType {
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_PASSWORD:
		credentialType = userCredentialDomain.CredentialTypePassword
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_PASSKEY:
		credentialType = userCredentialDomain.CredentialTypePasskey
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_OAUTH_LINK:
		credentialType = userCredentialDomain.CredentialTypeOauthLink
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_SAML:
		credentialType = userCredentialDomain.CredentialTypeSaml
	case usercredentialpb.AuthCredentialType_AUTH_CREDENTIAL_TYPE_LDAP:
		credentialType = userCredentialDomain.CredentialTypeLdap
	default:
		credentialType = userCredentialDomain.CredentialTypePassword
	}
	var passwordHash *string
	var hashAlg *string
	hashParams := make(map[string]interface{})
	if credentialType == userCredentialDomain.CredentialTypePassword && req.Password != "" {
		hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errx.ErrInternal.WithCause(err)
		}
		hashedStr := string(hashedBytes)
		passwordHash = &hashedStr
		alg := "bcrypt"
		hashAlg = &alg
		hashParams["cost"] = bcrypt.DefaultCost
	}
	cmd := userCredentialAppCommands.CreateUserCredentialCmd{
		UserID:             userID,
		CredentialType:     credentialType,
		PasswordHash:       passwordHash,
		HashAlg:            hashAlg,
		HashParams:         hashParams,
		Status:             userCredentialDomain.CredentialStatusActive,
		MustChangePassword: req.MustChangePassword,
	}
	userCredentialID, err := h.userCredentialAppSvc.CreateUserCredential(ctx, cmd)
	if err != nil {
		return nil, err
	}
	userCredentialView, err := h.userCredentialAppSvc.GetUserCredential(ctx, userCredentialID)
	if err != nil {
		return nil, err
	}
	userCredential := mapper.UserCredentialROToProto(&userCredentialView)
	return &usercredentialpb.CreateUserCredentialResponse{UserCredential: userCredential}, nil
}

func (h *UserCredentialHandler) GetUserCredentialByID(
	ctx context.Context,
	req *usercredentialpb.GetUserCredentialByIDRequest,
) (*usercredentialpb.GetUserCredentialByIDResponse, error) {
	userCredentialID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	userCredentialView, err := h.userCredentialAppSvc.GetUserCredential(ctx, userCredentialID)
	if err != nil {
		return nil, err
	}
	userCredential := mapper.UserCredentialROToProto(&userCredentialView)
	return &usercredentialpb.GetUserCredentialByIDResponse{UserCredential: userCredential}, nil
}

func (h *UserCredentialHandler) GetUserCredentialByUserID(
	ctx context.Context,
	req *usercredentialpb.GetUserCredentialByUserIDRequest,
) (*usercredentialpb.GetUserCredentialByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	userCredentialView, err := h.userCredentialAppSvc.GetUserCredentialByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	userCredential := mapper.UserCredentialROToProto(&userCredentialView)
	return &usercredentialpb.GetUserCredentialByUserIDResponse{UserCredential: userCredential}, nil
}
