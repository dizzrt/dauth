package identity

import (
	"context"

	"github.com/dizzrt/dauth/api/gen/errdef"
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/dauth/internal/domain/identity/entity"
	"github.com/dizzrt/dauth/internal/domain/identity/repo"
	"github.com/dizzrt/dauth/internal/infra/persistence/core"
	"github.com/dizzrt/dauth/internal/infra/persistence/core/gen/dao"
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
	m := fromIdentityUserEntity(user)
	err := utils.WrapError(impl.WithContext(ctx).Create(m))
	if err != nil {
		if !errdef.IsDuplicatedKey(err) {
			log.CtxErrorf(ctx, "create user failed, err: %v", err)
		}

		return err
	}

	user.UID = uint32(m.ID)
	return nil
}

func (impl *UserRepoImpl) GetUserByID(ctx context.Context, uid uint32) (*entity.User, error) {
	query := impl.Query()
	m, err := impl.WithContext(ctx).Where(
		query.IdentityUser.ID.Eq(int32(uid)),
		query.IdentityUser.DeletedAt.Eq(nil),
	).First()

	if err != nil {
		err = utils.WrapError(err)
		if !errdef.IsRecordNotFound(err) {
			log.CtxErrorf(ctx, "get user by id failed, err: %v", err)
		}

		return nil, err
	}

	user := toIdentityUserEntity(m)
	return user, nil
}

func (impl *UserRepoImpl) GetUserByName(ctx context.Context, username string) (*entity.User, error) {
	query := impl.Query()
	m, err := impl.WithContext(ctx).Where(
		query.IdentityUser.Username.Eq(username),
		query.IdentityUser.DeletedAt.Eq(nil),
	).First()

	if err != nil {
		err = utils.WrapError(err)
		if !errdef.IsRecordNotFound(err) {
			log.CtxErrorf(ctx, "get user by name failed, err: %v", err)
		}

		return nil, err
	}

	user := toIdentityUserEntity(m)
	return user, nil
}

func (impl *UserRepoImpl) ListUsers(ctx context.Context, page, size int32) ([]*entity.User, int64, error) {
	offset := (page - 1) * size

	query := impl.Query()
	users, total, err := impl.WithContext(ctx).
		Where(query.IdentityUser.DeletedAt.Eq(nil)).
		Order(query.IdentityUser.ID.Asc()).
		FindByPage(int(offset), int(size))

	if err != nil {
		log.CtxErrorf(ctx, "list users failed, err: %v", err)
		return nil, 0, utils.WrapError(err)
	}

	return toIdentityUserEntities(users), total, nil
}

func (impl *UserRepoImpl) UpdateUserStatus(ctx context.Context, uid uint32, status identity.UserStatus) error {
	query := impl.Query()
	info, err := impl.WithContext(ctx).
		Where(query.IdentityUser.ID.Eq(int32(uid))).
		Update(query.IdentityUser.Status, int32(status))

	if err != nil {
		log.CtxErrorf(ctx, "update user status failed, err: %v", err)
		return utils.WrapError(err)
	}

	if info.Error != nil {
		log.CtxErrorf(ctx, "update user status failed, err: %v", info.Error)
		return utils.WrapError(info.Error)
	}

	return nil
}
