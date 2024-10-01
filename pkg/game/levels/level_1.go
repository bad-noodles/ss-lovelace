package levels

import (
	"github.com/bad-noodles/ss-lovelace/pkg/message"
	"github.com/bad-noodles/ss-lovelace/pkg/ship/modules"
	"github.com/bad-noodles/ss-lovelace/pkg/theme"
)

type Level1 struct{}

func (r Level1) ModuleHandler() modules.ModuleHandler {
	return &modules.Co2Recycler{}
}

func (r Level1) Messages() []message.Message {
	return []message.Message{
		{
			Subject: "Welcome",
			Body: `Hello, space recruit!

Welcome to the ` + theme.Bold.Render("United Space Federation") + `.

You would usually receive training, but we see that you have and engineering degree and we are short on those, so welcome to your first mission!

You have been teleported to the ` + theme.Bold.Render("SS Lovelace") + `.

Unfotunatelly the ship has encountered itself in the middle of a big magnetic storm and that erased some floppy disks that contained the software to some very important systems.

When we tried the backup floppy disks, to our surprise, they were also erased!

You will require coding skills to do this job, so I hope you paid attention to those classes.

I know this is your first misson with absolutely no training, but the crew's lifes are on your hands.

You can use the arrows or h/j/k/l to nagivate between different items and panes of the UI.

The tabs below can be acessed by typing the number displayed on them.

Good luck!
  `,
		},
		{
			Subject: "Modules",
			Body: `The ship is composed of modules, each with it's own functionality.

As soon as a new module is made available for you, it will appear on your dashboard, along with the operational status.

It is your job to make sure all modules are green and healthy!

You will be briefed about how each module works individually, but there are some overal information that is true for every module.

The modules rely on the Transmission Control Protocol (TCP) and each module will have a port for you to connect to it.

Messages will be sent back and forth between your software and the module. Messages are delimited by a '\n'.

Please note that a module will constantly send new messages, and your software needs to keep answering it to maintain it healthy. This means a TCP connection per module you have enabled.
  `,
		},
	}
}
