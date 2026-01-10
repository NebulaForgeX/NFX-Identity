package hash_chain_checkpoints

import (
	hashChainCheckpointDomain "nfxid/modules/audit/domain/hash_chain_checkpoints"
)

type Service struct {
	hashChainCheckpointRepo *hashChainCheckpointDomain.Repo
}

func NewService(
	hashChainCheckpointRepo *hashChainCheckpointDomain.Repo,
) *Service {
	return &Service{
		hashChainCheckpointRepo: hashChainCheckpointRepo,
	}
}
