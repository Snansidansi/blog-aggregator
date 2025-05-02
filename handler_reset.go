package main

import (
	"context"
	"fmt"
)

func handlerResetDatabase(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: reset")
	}

	err := s.db.ResetDatabase(context.Background())
	if err != nil {
		return fmt.Errorf("unable to reset the database: %v: ", err)
	}

	err = s.Config.SetUser("")
	if err != nil {
		return fmt.Errorf("unable to clear the username in config")
	}

	fmt.Println("successfuly reset the database")
	return nil
}
