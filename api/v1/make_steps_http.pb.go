// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             (unknown)
// source: api/v1/make_steps.proto

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

const OperationMakeStepsServiceCreateMakeSteps = "/v1.MakeStepsService/CreateMakeSteps"
const OperationMakeStepsServiceDelMakeSteps = "/v1.MakeStepsService/DelMakeSteps"
const OperationMakeStepsServiceDelMakeStepsByDishId = "/v1.MakeStepsService/DelMakeStepsByDishId"
const OperationMakeStepsServiceGetDishesMakeSteps = "/v1.MakeStepsService/GetDishesMakeSteps"
const OperationMakeStepsServiceGetMakeSteps = "/v1.MakeStepsService/GetMakeSteps"
const OperationMakeStepsServiceUpdateMakeSteps = "/v1.MakeStepsService/UpdateMakeSteps"

type MakeStepsServiceHTTPServer interface {
	CreateMakeSteps(context.Context, *CreateMakeStepsRequest) (*CreateMakeStepsResponse, error)
	DelMakeSteps(context.Context, *DelMakeStepsRequest) (*DelMakeStepsResponse, error)
	DelMakeStepsByDishId(context.Context, *DelMakeStepsByDishIdRequest) (*DelMakeStepsByDishIdResponse, error)
	GetDishesMakeSteps(context.Context, *DishesMakeStepsRequest) (*DishesMakeStepsResponse, error)
	GetMakeSteps(context.Context, *GetMakeStepsRequest) (*GetMakeStepsResponse, error)
	UpdateMakeSteps(context.Context, *UpdateMakeStepsRequest) (*UpdateMakeStepsResponse, error)
}

func RegisterMakeStepsServiceHTTPServer(s *http.Server, srv MakeStepsServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/make-step/create", _MakeStepsService_CreateMakeSteps0_HTTP_Handler(srv))
	r.GET("/api/v1/make-step/get/{id}", _MakeStepsService_GetMakeSteps0_HTTP_Handler(srv))
	r.DELETE("/api/v1/make-step/delete/{id}", _MakeStepsService_DelMakeSteps0_HTTP_Handler(srv))
	r.DELETE("/api/v1/make-step/delete/dish/{dish_id}", _MakeStepsService_DelMakeStepsByDishId0_HTTP_Handler(srv))
	r.PUT("/api/v1/make-step/update", _MakeStepsService_UpdateMakeSteps0_HTTP_Handler(srv))
	r.GET("/api/v1/make-step/dish", _MakeStepsService_GetDishesMakeSteps0_HTTP_Handler(srv))
}

func _MakeStepsService_CreateMakeSteps0_HTTP_Handler(srv MakeStepsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMakeStepsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMakeStepsServiceCreateMakeSteps)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMakeSteps(ctx, req.(*CreateMakeStepsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateMakeStepsResponse)
		return ctx.Result(200, reply)
	}
}

func _MakeStepsService_GetMakeSteps0_HTTP_Handler(srv MakeStepsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMakeStepsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMakeStepsServiceGetMakeSteps)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMakeSteps(ctx, req.(*GetMakeStepsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMakeStepsResponse)
		return ctx.Result(200, reply)
	}
}

func _MakeStepsService_DelMakeSteps0_HTTP_Handler(srv MakeStepsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelMakeStepsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMakeStepsServiceDelMakeSteps)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelMakeSteps(ctx, req.(*DelMakeStepsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelMakeStepsResponse)
		return ctx.Result(200, reply)
	}
}

func _MakeStepsService_DelMakeStepsByDishId0_HTTP_Handler(srv MakeStepsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelMakeStepsByDishIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMakeStepsServiceDelMakeStepsByDishId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelMakeStepsByDishId(ctx, req.(*DelMakeStepsByDishIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelMakeStepsByDishIdResponse)
		return ctx.Result(200, reply)
	}
}

func _MakeStepsService_UpdateMakeSteps0_HTTP_Handler(srv MakeStepsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMakeStepsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMakeStepsServiceUpdateMakeSteps)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMakeSteps(ctx, req.(*UpdateMakeStepsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateMakeStepsResponse)
		return ctx.Result(200, reply)
	}
}

func _MakeStepsService_GetDishesMakeSteps0_HTTP_Handler(srv MakeStepsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DishesMakeStepsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMakeStepsServiceGetDishesMakeSteps)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDishesMakeSteps(ctx, req.(*DishesMakeStepsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DishesMakeStepsResponse)
		return ctx.Result(200, reply)
	}
}

type MakeStepsServiceHTTPClient interface {
	CreateMakeSteps(ctx context.Context, req *CreateMakeStepsRequest, opts ...http.CallOption) (rsp *CreateMakeStepsResponse, err error)
	DelMakeSteps(ctx context.Context, req *DelMakeStepsRequest, opts ...http.CallOption) (rsp *DelMakeStepsResponse, err error)
	DelMakeStepsByDishId(ctx context.Context, req *DelMakeStepsByDishIdRequest, opts ...http.CallOption) (rsp *DelMakeStepsByDishIdResponse, err error)
	GetDishesMakeSteps(ctx context.Context, req *DishesMakeStepsRequest, opts ...http.CallOption) (rsp *DishesMakeStepsResponse, err error)
	GetMakeSteps(ctx context.Context, req *GetMakeStepsRequest, opts ...http.CallOption) (rsp *GetMakeStepsResponse, err error)
	UpdateMakeSteps(ctx context.Context, req *UpdateMakeStepsRequest, opts ...http.CallOption) (rsp *UpdateMakeStepsResponse, err error)
}

type MakeStepsServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewMakeStepsServiceHTTPClient(client *http.Client) MakeStepsServiceHTTPClient {
	return &MakeStepsServiceHTTPClientImpl{client}
}

func (c *MakeStepsServiceHTTPClientImpl) CreateMakeSteps(ctx context.Context, in *CreateMakeStepsRequest, opts ...http.CallOption) (*CreateMakeStepsResponse, error) {
	var out CreateMakeStepsResponse
	pattern := "/api/v1/make-step/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMakeStepsServiceCreateMakeSteps))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MakeStepsServiceHTTPClientImpl) DelMakeSteps(ctx context.Context, in *DelMakeStepsRequest, opts ...http.CallOption) (*DelMakeStepsResponse, error) {
	var out DelMakeStepsResponse
	pattern := "/api/v1/make-step/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMakeStepsServiceDelMakeSteps))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MakeStepsServiceHTTPClientImpl) DelMakeStepsByDishId(ctx context.Context, in *DelMakeStepsByDishIdRequest, opts ...http.CallOption) (*DelMakeStepsByDishIdResponse, error) {
	var out DelMakeStepsByDishIdResponse
	pattern := "/api/v1/make-step/delete/dish/{dish_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMakeStepsServiceDelMakeStepsByDishId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MakeStepsServiceHTTPClientImpl) GetDishesMakeSteps(ctx context.Context, in *DishesMakeStepsRequest, opts ...http.CallOption) (*DishesMakeStepsResponse, error) {
	var out DishesMakeStepsResponse
	pattern := "/api/v1/make-step/dish"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMakeStepsServiceGetDishesMakeSteps))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MakeStepsServiceHTTPClientImpl) GetMakeSteps(ctx context.Context, in *GetMakeStepsRequest, opts ...http.CallOption) (*GetMakeStepsResponse, error) {
	var out GetMakeStepsResponse
	pattern := "/api/v1/make-step/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMakeStepsServiceGetMakeSteps))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MakeStepsServiceHTTPClientImpl) UpdateMakeSteps(ctx context.Context, in *UpdateMakeStepsRequest, opts ...http.CallOption) (*UpdateMakeStepsResponse, error) {
	var out UpdateMakeStepsResponse
	pattern := "/api/v1/make-step/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMakeStepsServiceUpdateMakeSteps))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}