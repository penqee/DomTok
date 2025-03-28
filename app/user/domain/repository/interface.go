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

package repository

import (
	"context"

	"github.com/west2-online/DomTok/app/user/domain/model"
)

// domain中的 repository 表示 service / use case 所依赖的外部资源，比如数据库、缓存等

// UserDB 表示持久化存储接口 (或者也可以叫做 DBPort)
type UserDB interface {
	IsUserExist(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	GetUserInfo(ctx context.Context, username string) (*model.User, error)
	GetAddressInfo(ctx context.Context, addressID int64) (*model.Address, error)
	CreateAddress(ctx context.Context, address *model.Address) (int64, error)
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

type UserCache interface {
	IsExist(ctx context.Context, key string) bool
	SetUserLogOut(ctx context.Context, key string, token string) error
	DeleteUserLogOut(ctx context.Context, key string) error
	SetUserBaned(ctx context.Context, key string) error
	DeleteUserBaned(ctx context.Context, key string) error
	UserLogOutKey(uid int64) string
	UserBanedKey(uid int64) string
	GetToken(ctx context.Context, key string) (string, error)
}
