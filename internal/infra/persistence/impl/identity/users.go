package identity

import (
	"context"

	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"github.com/dizzrt/dauth/internal/infra/persistence/core"
	"github.com/dizzrt/dauth/internal/infra/persistence/core/gen/dao"
	"github.com/dizzrt/dauth/internal/infra/utils"
	"github.com/dizzrt/ellie/log"
	"gorm.io/gorm"
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
	m := fromIdentityUserEntity(user)
	err := impl.WithContext(ctx).Create(m)
	if err != nil {
		log.CtxErrorf(ctx, "create user failed, err: %v", err)
		return utils.WrapError(err)
	}

	user.UID = uint32(m.ID)
	return nil
}

func (impl *UserRepoImpl) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	m, err := impl.WithContext(ctx).Where(dao.IdentityUser.ID.Eq(int32(uid))).First()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.CtxErrorf(ctx, "get user by id failed, err: %v", err)
		}

		return nil, utils.WrapError(err)
	}

	user := toIdentityUserEntity(m)
	return user, nil
}
