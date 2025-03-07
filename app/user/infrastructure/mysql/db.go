/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/west2-online/DomTok/app/user/domain/model"
	"github.com/west2-online/DomTok/app/user/domain/repository"
	"github.com/west2-online/DomTok/pkg/errno"
)

// userDB impl domain.UserDB defined domain
type userDB struct {
	client *gorm.DB
}

func (db *userDB) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	var ret User
	if err := db.client.WithContext(ctx).Where("id = ?", id).First(&ret).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.Errorf(errno.ErrRecordNotFound, "mysql: failed to query user: %v", err)
		}
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query user: %v", err)
	}
	return &model.User{
		Uid:      id,
		UserName: ret.Username,
		Role:     ret.Role,
	}, nil
}

func NewUserDB(client *gorm.DB) repository.UserDB {
	return &userDB{client: client}
}

func (db *userDB) CreateUser(ctx context.Context, u *model.User) (int64, error) {
	// 将 entity 转换成 mysql 这边的 model
	user := User{
		Username: u.UserName,
		Password: u.Password,
		Email:    u.Email,
	}

	if err := db.client.WithContext(ctx).Table(User{}.TableName()).Create(&user).Error; err != nil {
		return -1, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to create user: %v", err)
	}

	return user.ID, nil
}

func (db *userDB) IsUserExist(ctx context.Context, username string) (bool, error) {
	var user User
	err := db.client.WithContext(ctx).Table(User{}.TableName()).Where("username = ?", username).First(&user).Error
	if err != nil {
		// 这里虽然是数据库返回的 err 不为 nil,
		// 但这显然是业务上的错误, 而不是我们服务本身的
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		// 这里报错了就不是业务错误了, 而是服务级别的错误
		return false, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query user: %v", err)
	}
	return true, nil
}

func (db *userDB) GetUserInfo(ctx context.Context, username string) (*model.User, error) {
	var user User
	err := db.client.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.Errorf(errno.ServiceUserNotExist, "mysql: user not exist")
		}
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query user: %v", err)
	}

	resp := &model.User{
		Uid:      user.ID,
		UserName: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
	}
	return resp, nil
}

func (db *userDB) GetAddressInfo(ctx context.Context, addressID int64) (*model.Address, error) {
	var address Address
	err := db.client.WithContext(ctx).Table(Address{}.TableName()).Where("id = ?", addressID).First(&address).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.Errorf(errno.AddressNotExist, "mysql: address not exist")
		}
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query address: %v", err)
	}

	resp := &model.Address{
		AddressID: address.ID,
		Province:  address.Province,
		City:      address.City,
		Detail:    address.Detail,
	}

	return resp, nil
}

func (db *userDB) CreateAddress(ctx context.Context, address *model.Address) (int64, error) {
	addr := Address{
		Province: address.Province,
		City:     address.City,
		Detail:   address.Detail,
	}

	if err := db.client.WithContext(ctx).Table(Address{}.TableName()).Create(&addr).Error; err != nil {
		return -1, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to create address: %v", err)
	}

	return addr.ID, nil
}

func (db *userDB) UpdateUser(ctx context.Context, user *model.User) error {
	u := User{
		Username: user.UserName,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
	}
	err := db.client.WithContext(ctx).Table(u.TableName()).
		Where("id = ?", user.Uid).Updates(&u).Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to update user: %v", err)
	}
	return nil
}
