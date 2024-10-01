package ship

import (
	"github.com/bad-noodles/ss-lovelace/pkg/ship/modules"
)

type Ship struct {
	port    int
	modules []*modules.Module
}

func NewShip(initialPort int) *Ship {
	s := &Ship{
		initialPort,
		[]*modules.Module{},
	}

	return s
}

func (s *Ship) nextPort() int {
	s.port++
	return s.port
}

func (s *Ship) AddModule(handler modules.ModuleHandler) *modules.Module {
	mod := modules.NewModule(s.nextPort(), handler)
	s.modules = append(s.modules, mod)

	return mod
}

func (s *Ship) CheckHealth() (healthy bool, descs []modules.ModuleDescriptor) {
	healthy = true
	for _, mod := range s.modules {
		desc := mod.Descriptor()
		if !desc.Health {
			healthy = false
		}
		descs = append(descs, desc)
	}

	return
}
