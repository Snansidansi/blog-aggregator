package handler

import (
	"context"
	"fmt"

	"github.com/snansidansi/blog-aggregator/internal/state"
)

func HandlerResetDatabase(s *state.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: reset")
	}

	err := s.Db.ResetDatabase(context.Background())
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
