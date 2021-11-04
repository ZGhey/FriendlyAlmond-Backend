package test

import (
	pbJobModule "FriendlyAlmond_backend/proto/jobModule"
	"FriendlyAlmond_backend/service/jobModule/handler"
	"context"
	"testing"
)

func TestMostPopular(t *testing.T) {
	info := new(pbJobModule.Empty)
	p := &handler.JobModule{}
	result := new(pbJobModule.MostPopular)
	err := p.QueryMostPopular(context.Background(), info, result)
	if err != nil {
		return
	}
}

func TestTotalSales(t *testing.T) {
	info := new(pbJobModule.Empty)
	p := &handler.JobModule{}
	result := new(pbJobModule.TotalSales)
	err := p.QueryTotalSales(context.Background(), info, result)
	if err != nil {
		return
	}
}
