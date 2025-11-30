package cli

import(
	"github.com/Tinotsu/gator/internal/config"
)

type Commands struct {
	Function map[string]func(*State, Command) error
}

func NewCommands () *Commands {
	c := Commands{
		make(map[string]func(*State, Command) error),
	}
	return &c
}

func (c *Commands) Run(s *State, cmd Command) error {
	err := c.Function[cmd.Name](s, cmd)
	if err != nil {
		config.HandleError(err)
		return err
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Function[name] = f
}
