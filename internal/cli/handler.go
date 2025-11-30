package cli

import(
	"context"
	"log"
	"github.com/google/uuid"
	"errors"
	"fmt"
	"time"
	"github.com/Tinotsu/gator/internal/database"
	"github.com/Tinotsu/gator/internal/config"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) != 3 {
		return errors.New("login handler expects username as third argument")
	}

	Context := context.Background()
	n, err := s.DB.GetUser(Context, cmd.Arguments[2])
	if err != nil {
		fmt.Printf("user with that name need to be registred: %s", n.Name)
		config.HandleError(err)
		return err
	}
	s.Config.Username = cmd.Arguments[2]
	s.Config.SetUSer(s.Config.Username)
	fmt.Printf("Username %s has been set", s.Config.Username)
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Arguments) != 3 {
		return errors.New("register handler expects username as third argument")
	}

	u := new(database.CreateUserParams)

	Context := context.Background()
	
	u.ID = uuid.New()
	u.Name = cmd.Arguments[2]
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	n, err := s.DB.GetUser(Context, cmd.Arguments[2])
	if err == nil {
		fmt.Printf("\nuser with that name already exists: %s\n", n.Name)
		config.HandleError(err)
		return err
	}
	
	_, err = s.DB.CreateUser(Context, *u)
	config.HandleError(err)
	s.Config.SetUSer(cmd.Arguments[2])

	fmt.Printf("\nUser %s registred !\n", cmd.Arguments[2])
	log.Println(s.Config.Username)
	return nil
}
