package identity

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"github.com/dizzrt/dauth/internal/infra/repo/core"
	"github.com/dizzrt/dauth/internal/infra/repo/core/gen/dao"
	"github.com/dizzrt/dauth/internal/infra/repo/core/gen/model"
	"github.com/dizzrt/dauth/internal/infra/utils"
	"github.com/dizzrt/ellie/log"
)

var _ repo.UserRepo = (*UserRepoImpl)(nil)

type UserRepoImpl struct {
	core.RepoCore
}

func NewUserRepoImpl(repoCore core.RepoCore) repo.UserRepo {
	return &UserRepoImpl{
		RepoCore: repoCore,
	}
}

func (impl *UserRepoImpl) WithContext(ctx context.Context) dao.IIdentityUserDo {
	return impl.Query().IdentityUser.WithContext(ctx)
}

func (impl *UserRepoImpl) CreateUser(ctx context.Context, user *entity.User) error {
	m := model.IdentityUser{}
	err := impl.WithContext(ctx).Create(&m)
	if err != nil {
		log.CtxErrorf(ctx, "create user failed, err: %v", err)
		return utils.WrapError(err)
	}

	user.UID = uint32(m.ID)
	return nil
}

func (impl *UserRepoImpl) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	return nil, nil
}
