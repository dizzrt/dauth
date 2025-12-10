package biz

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/dizzrt/dauth/api/gen/errdef"
	sp_api "github.com/dizzrt/dauth/api/gen/sp"
	"github.com/dizzrt/dauth/internal/domain/sp/entity"
	"github.com/dizzrt/dauth/internal/domain/sp/repo"
	"github.com/dizzrt/ellie/log"
)

var _ ServiceProviderBiz = (*serviceProviderBiz)(nil)

type ServiceProviderBiz interface {
	CreateServiceProvider(ctx context.Context, spEntity *entity.ServiceProvider, scopeIDs []uint32) (uint32, error)
	GetServiceProvider(ctx context.Context, spID uint32) (*entity.ServiceProvider, error)
	ValidateServiceProvider(ctx context.Context, spID uint32, scope string) (bool, string, error)
}

type serviceProviderBiz struct {
	spRepo                 repo.ServiceProviderRepo
	scopeRepo              repo.ScopeRepo
	spScopeAssociationRepo repo.SPScopeAssociationRepo
}

func NewServiceProviderBiz(spRepo repo.ServiceProviderRepo, scopeRepo repo.ScopeRepo, spScopeAssociationRepo repo.SPScopeAssociationRepo) ServiceProviderBiz {
	return &serviceProviderBiz{
		spRepo:                 spRepo,
		scopeRepo:              scopeRepo,
		spScopeAssociationRepo: spScopeAssociationRepo,
	}
}

func (biz *serviceProviderBiz) CreateServiceProvider(ctx context.Context, spEntity *entity.ServiceProvider, scopeIDs []uint32) (uint32, error) {
	// create service provider
	sp, err := biz.spRepo.Create(ctx, spEntity)
	if err != nil {
		log.CtxErrorf(ctx, "failed to create sp, err: %v", err)
		return 0, err
	}

	// create sp scope associations
	if len(scopeIDs) > 0 {
		if err := biz.spScopeAssociationRepo.CreateAssociations(ctx, sp.ID, scopeIDs); err != nil {
			log.CtxErrorf(ctx, "failed to create sp scope associations, err: %v", err)
			return sp.ID, err
		}
	}

	return sp.ID, nil
}

func (biz *serviceProviderBiz) GetServiceProvider(ctx context.Context, spID uint32) (*entity.ServiceProvider, error) {
	return biz.spRepo.Get(ctx, spID)
}

func (biz *serviceProviderBiz) ValidateServiceProvider(ctx context.Context, spID uint32, scope string) (bool, string, error) {
	var msg string

	// check if sp exists
	sp, err := biz.spRepo.Get(ctx, spID)
	if err != nil {
		if errors.Is(err, errdef.RecordNotFound()) {
			msg = "sp not found"
			return false, msg, errors.New(msg)
		}

		msg = fmt.Sprintf("failed to get sp by id, err: %v", err)
		log.CtxErrorf(ctx, msg, err)
		return false, msg, err
	}

	// check if sp is active
	if sp.Status != sp_api.ServiceProvider_ACTIVE || sp.DeletedAt.Valid {
		msg = "sp is not active or deleted"
		return false, msg, errors.New(msg)
	}

	scopeAssociations, err := biz.spScopeAssociationRepo.GetSPScopes(ctx, spID)
	if err != nil && !errors.Is(err, errdef.RecordNotFound()) {
		msg = fmt.Sprintf("failed to get sp scopes, err: %v", err)
		log.CtxErrorf(ctx, msg, err)
		return false, msg, err
	}

	spScopes := make(map[string]struct{}, len(scopeAssociations))
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
			spScopes[temp.Name] = struct{}{}
		}
	}

	// check if scope is in sp scopes
	targetScope := strings.SplitSeq(scope, " ")
	for s := range targetScope {
		if _, ok := spScopes[s]; !ok {
			msg = fmt.Sprintf("scope %s is not in sp scopes", s)
			return false, msg, errors.New(msg)
		}
	}

	return true, "", nil
}
