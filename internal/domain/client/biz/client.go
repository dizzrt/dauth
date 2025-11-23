package biz

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/dizzrt/dauth/api/gen/client"
	"github.com/dizzrt/dauth/internal/domain/client/entity"
	"github.com/dizzrt/dauth/internal/domain/client/repo"
	"github.com/dizzrt/ellie/log"
	"gorm.io/gorm"
)

var _ ClientBiz = (*clientBiz)(nil)

type ClientBiz interface {
	Create(ctx context.Context, clientEntity *entity.Client, scopeIDs []uint32) (uint32, error)
	Validate(ctx context.Context, clientID uint32, scope string) (bool, string, error)
}

type clientBiz struct {
	clientRepo                 repo.ClientRepo
	scopeRepo                  repo.ScopeRepo
	clientScopeAssociationRepo repo.ClientScopeAssociationRepo
}

func NewClientBiz(clientRepo repo.ClientRepo, scopeRepo repo.ScopeRepo, clientScopeAssociationRepo repo.ClientScopeAssociationRepo) ClientBiz {
	return &clientBiz{
		clientRepo:                 clientRepo,
		scopeRepo:                  scopeRepo,
		clientScopeAssociationRepo: clientScopeAssociationRepo,
	}
}

func (biz *clientBiz) Create(ctx context.Context, clientEntity *entity.Client, scopeIDs []uint32) (uint32, error) {
	// create client
	cli, err := biz.clientRepo.Create(ctx, clientEntity)
	if err != nil {
		log.CtxErrorf(ctx, "failed to create client, err: %v", err)
		return 0, err
	}

	// create client scope associations
	if len(scopeIDs) > 0 {
		if err := biz.clientScopeAssociationRepo.CreateAssociations(ctx, cli.ID, scopeIDs); err != nil {
			log.CtxErrorf(ctx, "failed to create client scope associations, err: %v", err)
			return cli.ID, err
		}
	}

	return cli.ID, nil
}

func (biz *clientBiz) Validate(ctx context.Context, clientID uint32, scope string) (bool, string, error) {
	var msg string

	// check if client exists
	cli, err := biz.clientRepo.Get(ctx, clientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = "client not found"
			return false, msg, errors.New(msg)
		}

		msg = fmt.Sprintf("failed to get client by id, err: %v", err)
		log.CtxErrorf(ctx, msg, err)
		return false, msg, err
	}

	// check if client is active
	if cli.Status != client.Client_ACTIVE || cli.DeletedAt.Valid {
		msg = "client is not active or deleted"
		return false, msg, errors.New(msg)
	}

	scopeAssociations, err := biz.clientScopeAssociationRepo.GetClientScopes(ctx, clientID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		msg = fmt.Sprintf("failed to get client scopes, err: %v", err)
		log.CtxErrorf(ctx, msg, err)
		return false, msg, err
	}

	clientScopes := make(map[string]struct{}, len(scopeAssociations))
	if len(scopeAssociations) > 0 {
		// extract scope ids
		scopeIDs := make([]uint32, 0, len(scopeAssociations))
		for _, sa := range scopeAssociations {
			scopeIDs = append(scopeIDs, sa.ScopeID)
		}

		// get scopes by ids
		scopes, err := biz.scopeRepo.GetScopesByIDs(ctx, scopeIDs)
		if err != nil {
			msg = fmt.Sprintf("failed to get scopes by ids: %v", err)
			log.CtxErrorf(ctx, msg, err)
			return false, msg, err
		}

		// extract scope names
		for _, temp := range scopes {
			clientScopes[temp.Name] = struct{}{}
		}
	}

	// check if scope is in client scopes
	targetScope := strings.SplitSeq(scope, " ")
	for s := range targetScope {
		if _, ok := clientScopes[s]; !ok {
			msg = fmt.Sprintf("scope %s is not in client scopes", s)
			return false, msg, errors.New(msg)
		}
	}

	return true, "", nil
}
