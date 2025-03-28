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

// Code generated by Kitex v0.12.1. DO NOT EDIT.

package paymentservice

import (
	"context"
	"errors"

	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"

	payment "github.com/west2-online/DomTok/kitex_gen/payment"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ProcessPayment": kitex.NewMethodInfo(
		processPaymentHandler,
		newPaymentServiceProcessPaymentArgs,
		newPaymentServiceProcessPaymentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"RequestPaymentToken": kitex.NewMethodInfo(
		requestPaymentTokenHandler,
		newPaymentServiceRequestPaymentTokenArgs,
		newPaymentServiceRequestPaymentTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"RequestPaymentCheckout": kitex.NewMethodInfo(
		requestPaymentCheckoutHandler,
		newPaymentServiceRequestPaymentCheckoutArgs,
		newPaymentServiceRequestPaymentCheckoutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"RefundReview": kitex.NewMethodInfo(
		refundReviewHandler,
		newPaymentServiceRefundReviewArgs,
		newPaymentServiceRefundReviewResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"RequestRefund": kitex.NewMethodInfo(
		requestRefundHandler,
		newPaymentServiceRequestRefundArgs,
		newPaymentServiceRequestRefundResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	paymentServiceServiceInfo                = NewServiceInfo()
	paymentServiceServiceInfoForClient       = NewServiceInfoForClient()
	paymentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return paymentServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PaymentService"
	handlerType := (*payment.PaymentService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "payment",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func processPaymentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment.PaymentServiceProcessPaymentArgs)
	realResult := result.(*payment.PaymentServiceProcessPaymentResult)
	success, err := handler.(payment.PaymentService).ProcessPayment(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentServiceProcessPaymentArgs() interface{} {
	return payment.NewPaymentServiceProcessPaymentArgs()
}

func newPaymentServiceProcessPaymentResult() interface{} {
	return payment.NewPaymentServiceProcessPaymentResult()
}

func requestPaymentTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment.PaymentServiceRequestPaymentTokenArgs)
	realResult := result.(*payment.PaymentServiceRequestPaymentTokenResult)
	success, err := handler.(payment.PaymentService).RequestPaymentToken(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentServiceRequestPaymentTokenArgs() interface{} {
	return payment.NewPaymentServiceRequestPaymentTokenArgs()
}

func newPaymentServiceRequestPaymentTokenResult() interface{} {
	return payment.NewPaymentServiceRequestPaymentTokenResult()
}

func requestPaymentCheckoutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment.PaymentServiceRequestPaymentCheckoutArgs)
	realResult := result.(*payment.PaymentServiceRequestPaymentCheckoutResult)
	success, err := handler.(payment.PaymentService).RequestPaymentCheckout(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentServiceRequestPaymentCheckoutArgs() interface{} {
	return payment.NewPaymentServiceRequestPaymentCheckoutArgs()
}

func newPaymentServiceRequestPaymentCheckoutResult() interface{} {
	return payment.NewPaymentServiceRequestPaymentCheckoutResult()
}

func refundReviewHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment.PaymentServiceRefundReviewArgs)
	realResult := result.(*payment.PaymentServiceRefundReviewResult)
	success, err := handler.(payment.PaymentService).RefundReview(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentServiceRefundReviewArgs() interface{} {
	return payment.NewPaymentServiceRefundReviewArgs()
}

func newPaymentServiceRefundReviewResult() interface{} {
	return payment.NewPaymentServiceRefundReviewResult()
}

func requestRefundHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment.PaymentServiceRequestRefundArgs)
	realResult := result.(*payment.PaymentServiceRequestRefundResult)
	success, err := handler.(payment.PaymentService).RequestRefund(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentServiceRequestRefundArgs() interface{} {
	return payment.NewPaymentServiceRequestRefundArgs()
}

func newPaymentServiceRequestRefundResult() interface{} {
	return payment.NewPaymentServiceRequestRefundResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ProcessPayment(ctx context.Context, request *payment.PaymentRequest) (r *payment.PaymentResponse, err error) {
	var _args payment.PaymentServiceProcessPaymentArgs
	_args.Request = request
	var _result payment.PaymentServiceProcessPaymentResult
	if err = p.c.Call(ctx, "ProcessPayment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RequestPaymentToken(ctx context.Context, request *payment.PaymentTokenRequest) (r *payment.PaymentTokenResponse, err error) {
	var _args payment.PaymentServiceRequestPaymentTokenArgs
	_args.Request = request
	var _result payment.PaymentServiceRequestPaymentTokenResult
	if err = p.c.Call(ctx, "RequestPaymentToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RequestPaymentCheckout(ctx context.Context, request *payment.PaymentCheckoutRequest) (r *payment.PaymentCheckoutResponse, err error) {
	var _args payment.PaymentServiceRequestPaymentCheckoutArgs
	_args.Request = request
	var _result payment.PaymentServiceRequestPaymentCheckoutResult
	if err = p.c.Call(ctx, "RequestPaymentCheckout", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RefundReview(ctx context.Context, request *payment.RefundReviewRequest) (r *payment.RefundReviewResponse, err error) {
	var _args payment.PaymentServiceRefundReviewArgs
	_args.Request = request
	var _result payment.PaymentServiceRefundReviewResult
	if err = p.c.Call(ctx, "RefundReview", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RequestRefund(ctx context.Context, request *payment.RefundRequest) (r *payment.RefundResponse, err error) {
	var _args payment.PaymentServiceRequestRefundArgs
	_args.Request = request
	var _result payment.PaymentServiceRequestRefundResult
	if err = p.c.Call(ctx, "RequestRefund", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
