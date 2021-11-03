package test

import (
	pbConfig "FriendlyAlmond_backend/proto/configuration"
	"FriendlyAlmond_backend/service/configuration/handler"
	"context"
	"testing"
)

func TestQueryBoat(t *testing.T) {
	info := new(pbConfig.Empty)
	p := &handler.Config{}
	result := new(pbConfig.ListBoat)
	err := p.QueryBoat(context.Background(), info, result)
	if err != nil {
		return
	}
}

func TestQueryComponent(t *testing.T) {
	info := new(pbConfig.Category)
	info.Type = "Motor"
	p := &handler.Config{}
	result := new(pbConfig.ListComponent)
	err := p.QueryComponent(context.Background(), info, result)
	if err != nil {
		return
	}
}

func TestQueryComponentById(t *testing.T) {
	info := new(pbConfig.Section)
	info.Id = 1
	p := &handler.Config{}
	result := new(pbConfig.Section)
	for i := 0; i < 10; i++ {
		err := p.QuerySectionById(context.Background(), info, result)
		if err != nil {
			return
		}
	}
}
