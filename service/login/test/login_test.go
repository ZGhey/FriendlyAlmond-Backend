package test

import (
	"FriendlyAlmond_backend/pkg/utils"
	pbLogin "FriendlyAlmond_backend/proto/login"
	"FriendlyAlmond_backend/service/login/handler"
	"context"
	"testing"
)

func TestRegisting(t *testing.T) {
	info := new(pbLogin.UserInfo)
	info.LastName = "lastname"
	info.FirstName = "firstname"
	info.Uid = utils.NewLen(10)
	info.Phone = "04802xxx56"
	info.Address = "address"
	info.AreaCode = "+61"
	p := &handler.Login{}
	result := new(pbLogin.OperationResult)
	err := p.Register(context.Background(), info, result)
	if err != nil {
		return
	}
}

func TestUpdate(t *testing.T) {
	info := new(pbLogin.UserInfo)
	info.LastName = "lastname"
	info.FirstName = "firstname"
	info.Uid = utils.NewLen(10)
	info.Phone = "04802xxx56"
	info.Address = "address"
	info.AreaCode = "+61"
	p := &handler.Login{}
	result := new(pbLogin.OperationResult)
	err := p.Update(context.Background(), info, result)
	if err != nil {
		return
	}
}

func TestQuery(t *testing.T) {
	info := new(pbLogin.UserInfo)
	info.Uid = utils.NewLen(10)
	p := &handler.Login{}
	result := new(pbLogin.OperationResult)
	err := p.Update(context.Background(), info, result)
	if err != nil {
		return
	}
}

func TestIsSameEmail(t *testing.T) {
	info := new(pbLogin.UserInfo)
	info.Email = "test222@gmail.com"
	p := &handler.Login{}
	result := new(pbLogin.OperationResult)
	err := p.IsSameEmail(context.Background(), info, result)
	if err != nil {
		return
	}
}
