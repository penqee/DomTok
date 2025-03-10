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

// Code generated by hertz generator.

package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	api "github.com/west2-online/DomTok/app/gateway/model/api/user"
	"github.com/west2-online/DomTok/app/gateway/pack"
	"github.com/west2-online/DomTok/app/gateway/rpc"
	"github.com/west2-online/DomTok/kitex_gen/user"
	"github.com/west2-online/DomTok/pkg/constants"
	"github.com/west2-online/DomTok/pkg/errno"
	"github.com/west2-online/DomTok/pkg/utils"
)

// Register .
// @router api/v1/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}
	resp, err := rpc.RegisterRPC(ctx, &user.RegisterRequest{
		Username: req.Name,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}
	pack.RespData(c, resp)
}

// Login .
// @router api/v1/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	resp, err := rpc.LoginRPC(ctx, &user.LoginRequest{
		Username: req.Name,
		Password: req.Password,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}

	accessToken, refreshToken, err := utils.CreateAllToken(resp.User.UserId)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	c.Header(constants.AccessTokenHeader, accessToken)
	c.Header(constants.RefreshTokenHeader, refreshToken)

	pack.RespData(c, resp)
}

// GetAddress .
// @router api/v1/user/address [GET]
func GetAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetAddressRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	resp := new(api.GetAddressResponse)
	resp, err = rpc.GetAddressRPC(ctx, &user.GetAddressRequest{AddressId: req.AddressId})
	if err != nil {
		pack.RespError(c, err)
		return
	}
	pack.RespData(c, resp.Address)
}

// AddAddress .
// @router api/v1/user/address [POST]
func AddAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.AddAddressRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamVerifyError.WithError(err))
		return
	}

	resp := new(api.AddAddressResponse)
	resp.AddressID, err = rpc.AddAddressRPC(ctx, &user.AddAddressRequest{
		Province: req.Province,
		City:     req.City,
		Detail:   req.Detail,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}
	pack.RespData(c, resp.AddressID)
}
