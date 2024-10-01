package levels

import (
	"github.com/bad-noodles/ss-lovelace/pkg/message"
	"github.com/bad-noodles/ss-lovelace/pkg/ship/modules"
)

type Level interface {
	Messages() []message.Message
	ModuleHandler() modules.ModuleHandler
}

var Levels = []Level{
	Level1{},
}
