package game

import (
	"time"

	"github.com/bad-noodles/ss-lovelace/pkg/game/levels"
	"github.com/bad-noodles/ss-lovelace/pkg/message"
	"github.com/bad-noodles/ss-lovelace/pkg/ship"
	"github.com/bad-noodles/ss-lovelace/pkg/ship/modules"
)

type Game struct {
	ship           *ship.Ship
	levelIndex     int
	gameOver       bool
	ModulesChannel chan []modules.ModuleDescriptor
	MessageChannel chan message.Message
}

func NewGame() *Game {
	return &Game{
		ship.NewShip(9000),
		-1,
		false,
		make(chan []modules.ModuleDescriptor, 1),
		make(chan message.Message, len(levels.Levels[0].Messages())+1), // +1 for module message
	}
}

func (g *Game) Start() {
	g.nextLevel()

	go func() {
		for {
			healthy, mods := g.ship.CheckHealth()
			g.ModulesChannel <- mods
			if healthy {
				g.nextLevel()
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func (g *Game) nextLevel() {
	if g.gameOver {
		return
	}

	g.levelIndex++

	if g.levelIndex == len(levels.Levels) {
		g.gameOver = true
		g.MessageChannel <- message.Message{
			Subject: "Game Over",
			Body: `That is all for now!

The game is still in alpha, many more levels to come!
`,
		}
		return
	}

	l := levels.Levels[g.levelIndex]
	modHandler := l.ModuleHandler()

	for _, msg := range l.Messages() {
		g.MessageChannel <- msg
	}

	g.ship.AddModule(modHandler)

	g.MessageChannel <- modHandler.Message()
}
