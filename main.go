package main

import (
	"fmt"
	"gatorapp/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

// command represents a CLI command
type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

// register a handler
func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

// run executes a command if it exists
func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username is required")
	}
	username := cmd.args[0]

	// set user and write to config
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}

	fmt.Printf("User set to '%s'\n", username)
	return nil
}

func main() {
	// read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	// setup state
	appState := &state{cfg: &cfg}

	// setup commands
	cmds := &commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	// parse CLI args
	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{name: cmdName, args: cmdArgs}

	// run command
	if err := cmds.run(appState, cmd); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	// // Read config
	// cfg, err := config.Read()
	// if err != nil {
	// 	log.Fatalf("failed to read config: %v", err)
	// }

	// // Set current user
	// err = cfg.SetUser("your_name") // <-- replace with your name
	// if err != nil {
	// 	log.Fatalf("failed to set user: %v", err)
	// }

	// // Read again
	// updatedCfg, err := config.Read()
	// if err != nil {
	// 	log.Fatalf("failed to read config: %v", err)
	// }

	// fmt.Println(updatedCfg.DBUrl)
	// Print config
	// fmt.Printf("Config: %+v\n", updatedCfg)
}
