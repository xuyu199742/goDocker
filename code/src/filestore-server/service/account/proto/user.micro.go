// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package go_micro_service_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	// 用户注册
	Signup(ctx context.Context, in *ReqSignup, opts ...client.CallOption) (*RespSignup, error)
	// 用户登录
	Signin(ctx context.Context, in *ReqSignin, opts ...client.CallOption) (*RespSignin, error)
	// 获取用户信息
	UserInfo(ctx context.Context, in *ReqUserInfo, opts ...client.CallOption) (*RespUserInfo, error)
	// 获取用户文件
	UserFiles(ctx context.Context, in *ReqUserFile, opts ...client.CallOption) (*RespUserFile, error)
	// 获取用户文件
	UserFileRename(ctx context.Context, in *ReqUserFileRename, opts ...client.CallOption) (*RespUserFileRename, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.service.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Signup(ctx context.Context, in *ReqSignup, opts ...client.CallOption) (*RespSignup, error) {
	req := c.c.NewRequest(c.name, "UserService.Signup", in)
	out := new(RespSignup)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Signin(ctx context.Context, in *ReqSignin, opts ...client.CallOption) (*RespSignin, error) {
	req := c.c.NewRequest(c.name, "UserService.Signin", in)
	out := new(RespSignin)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserInfo(ctx context.Context, in *ReqUserInfo, opts ...client.CallOption) (*RespUserInfo, error) {
	req := c.c.NewRequest(c.name, "UserService.UserInfo", in)
	out := new(RespUserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserFiles(ctx context.Context, in *ReqUserFile, opts ...client.CallOption) (*RespUserFile, error) {
	req := c.c.NewRequest(c.name, "UserService.UserFiles", in)
	out := new(RespUserFile)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserFileRename(ctx context.Context, in *ReqUserFileRename, opts ...client.CallOption) (*RespUserFileRename, error) {
	req := c.c.NewRequest(c.name, "UserService.UserFileRename", in)
	out := new(RespUserFileRename)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	// 用户注册
	Signup(context.Context, *ReqSignup, *RespSignup) error
	// 用户登录
	Signin(context.Context, *ReqSignin, *RespSignin) error
	// 获取用户信息
	UserInfo(context.Context, *ReqUserInfo, *RespUserInfo) error
	// 获取用户文件
	UserFiles(context.Context, *ReqUserFile, *RespUserFile) error
	// 获取用户文件
	UserFileRename(context.Context, *ReqUserFileRename, *RespUserFileRename) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Signup(ctx context.Context, in *ReqSignup, out *RespSignup) error
		Signin(ctx context.Context, in *ReqSignin, out *RespSignin) error
		UserInfo(ctx context.Context, in *ReqUserInfo, out *RespUserInfo) error
		UserFiles(ctx context.Context, in *ReqUserFile, out *RespUserFile) error
		UserFileRename(ctx context.Context, in *ReqUserFileRename, out *RespUserFileRename) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Signup(ctx context.Context, in *ReqSignup, out *RespSignup) error {
	return h.UserServiceHandler.Signup(ctx, in, out)
}

func (h *userServiceHandler) Signin(ctx context.Context, in *ReqSignin, out *RespSignin) error {
	return h.UserServiceHandler.Signin(ctx, in, out)
}

func (h *userServiceHandler) UserInfo(ctx context.Context, in *ReqUserInfo, out *RespUserInfo) error {
	return h.UserServiceHandler.UserInfo(ctx, in, out)
}

func (h *userServiceHandler) UserFiles(ctx context.Context, in *ReqUserFile, out *RespUserFile) error {
	return h.UserServiceHandler.UserFiles(ctx, in, out)
}

func (h *userServiceHandler) UserFileRename(ctx context.Context, in *ReqUserFileRename, out *RespUserFileRename) error {
	return h.UserServiceHandler.UserFileRename(ctx, in, out)
}
