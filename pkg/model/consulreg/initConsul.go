package consulreg

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/debug/trace"
	"github.com/asim/go-micro/v3/registry"
	"strings"
	"sync"
)

var (
	MicroSer micro.Service
	once     sync.Once
)

//func init() {
//	consulReg := consul.NewRegistry(
//		registry.Addrs(utils2.GetConfigStr("micro.addr")))
//	MicroSer = micro.NewService(
//		micro.Name(utils2.GetConfigStr("micro.name")),
//		micro.Registry(consulReg),
//		micro.Tracer(trace.DefaultTracer),
//	)
//}

func InitMicro(addrs, name string) error {
	if addrs == "" || name == "" {
		return fmt.Errorf("addr(%s) or name(%s) can not be empty", addrs, name)
	}
	addr := strings.Split(addrs, ",")
	once.Do(func() {
		consulReg := consul.NewRegistry(
			registry.Addrs(addr...),
		)
		MicroSer = micro.NewService(
			micro.Name(name),
			micro.Registry(consulReg),
			micro.Tracer(trace.DefaultTracer),
		)
	})
	return nil
}
