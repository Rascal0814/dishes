// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             (unknown)
// source: api/v1/order.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationOrdersServiceCreateOrders = "/v1.OrdersService/CreateOrders"
const OperationOrdersServiceDelOrders = "/v1.OrdersService/DelOrders"
const OperationOrdersServiceGetOrders = "/v1.OrdersService/GetOrders"
const OperationOrdersServiceListOrders = "/v1.OrdersService/ListOrders"
const OperationOrdersServiceUpdateOrders = "/v1.OrdersService/UpdateOrders"

type OrdersServiceHTTPServer interface {
	CreateOrders(context.Context, *CreateOrdersRequest) (*CreateOrdersResponse, error)
	DelOrders(context.Context, *DelOrdersRequest) (*DelOrdersResponse, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
	UpdateOrders(context.Context, *UpdateOrdersRequest) (*UpdateOrdersResponse, error)
}

func RegisterOrdersServiceHTTPServer(s *http.Server, srv OrdersServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/order/create", _OrdersService_CreateOrders0_HTTP_Handler(srv))
	r.GET("/api/v1/order/get/{id}", _OrdersService_GetOrders0_HTTP_Handler(srv))
	r.DELETE("/api/v1/order/delete/{id}", _OrdersService_DelOrders0_HTTP_Handler(srv))
	r.PUT("/api/v1/order/update", _OrdersService_UpdateOrders0_HTTP_Handler(srv))
	r.GET("/api/v1/order/list", _OrdersService_ListOrders0_HTTP_Handler(srv))
}

func _OrdersService_CreateOrders0_HTTP_Handler(srv OrdersServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrdersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrdersServiceCreateOrders)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOrders(ctx, req.(*CreateOrdersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOrdersResponse)
		return ctx.Result(200, reply)
	}
}

func _OrdersService_GetOrders0_HTTP_Handler(srv OrdersServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOrdersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrdersServiceGetOrders)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOrders(ctx, req.(*GetOrdersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOrdersResponse)
		return ctx.Result(200, reply)
	}
}

func _OrdersService_DelOrders0_HTTP_Handler(srv OrdersServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelOrdersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrdersServiceDelOrders)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelOrders(ctx, req.(*DelOrdersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelOrdersResponse)
		return ctx.Result(200, reply)
	}
}

func _OrdersService_UpdateOrders0_HTTP_Handler(srv OrdersServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateOrdersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrdersServiceUpdateOrders)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOrders(ctx, req.(*UpdateOrdersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateOrdersResponse)
		return ctx.Result(200, reply)
	}
}

func _OrdersService_ListOrders0_HTTP_Handler(srv OrdersServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrdersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrdersServiceListOrders)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrders(ctx, req.(*ListOrdersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOrdersResponse)
		return ctx.Result(200, reply)
	}
}

type OrdersServiceHTTPClient interface {
	CreateOrders(ctx context.Context, req *CreateOrdersRequest, opts ...http.CallOption) (rsp *CreateOrdersResponse, err error)
	DelOrders(ctx context.Context, req *DelOrdersRequest, opts ...http.CallOption) (rsp *DelOrdersResponse, err error)
	GetOrders(ctx context.Context, req *GetOrdersRequest, opts ...http.CallOption) (rsp *GetOrdersResponse, err error)
	ListOrders(ctx context.Context, req *ListOrdersRequest, opts ...http.CallOption) (rsp *ListOrdersResponse, err error)
	UpdateOrders(ctx context.Context, req *UpdateOrdersRequest, opts ...http.CallOption) (rsp *UpdateOrdersResponse, err error)
}

type OrdersServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewOrdersServiceHTTPClient(client *http.Client) OrdersServiceHTTPClient {
	return &OrdersServiceHTTPClientImpl{client}
}

func (c *OrdersServiceHTTPClientImpl) CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...http.CallOption) (*CreateOrdersResponse, error) {
	var out CreateOrdersResponse
	pattern := "/api/v1/order/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrdersServiceCreateOrders))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrdersServiceHTTPClientImpl) DelOrders(ctx context.Context, in *DelOrdersRequest, opts ...http.CallOption) (*DelOrdersResponse, error) {
	var out DelOrdersResponse
	pattern := "/api/v1/order/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrdersServiceDelOrders))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrdersServiceHTTPClientImpl) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...http.CallOption) (*GetOrdersResponse, error) {
	var out GetOrdersResponse
	pattern := "/api/v1/order/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrdersServiceGetOrders))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrdersServiceHTTPClientImpl) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...http.CallOption) (*ListOrdersResponse, error) {
	var out ListOrdersResponse
	pattern := "/api/v1/order/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrdersServiceListOrders))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrdersServiceHTTPClientImpl) UpdateOrders(ctx context.Context, in *UpdateOrdersRequest, opts ...http.CallOption) (*UpdateOrdersResponse, error) {
	var out UpdateOrdersResponse
	pattern := "/api/v1/order/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrdersServiceUpdateOrders))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}