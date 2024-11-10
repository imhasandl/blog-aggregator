package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete the user: %w", err)
	}

	fmt.Println("Databse reset successfully")
	return nil
}