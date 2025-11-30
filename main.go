package main

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/Tinotsu/gator/internal/config"
	"github.com/Tinotsu/gator/internal/cli"
	"github.com/Tinotsu/gator/internal/database"
	_ "github.com/lib/pq"
)


func main () {
	cfg := config.NewConfig()
	cfg = config.Read(cfg)
	
	s := cli.NewState()
	s.Config = cfg

	db, err := sql.Open("postgres", s.Config.DBURL)
	config.HandleError(err)

	dbQueries := database.New(db)

	s.DB = dbQueries

	cmds := cli.NewCommands()
	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)
	cmds.Register("reset", cli.Reset)
	cmds.Register("users", cli.Users)
	cmds.Register("agg", cli.RSS)
	cmds.Register("addfeed", cli.AddFeed)
	cmds.Register("feeds", cli.Feeds)
	config.HandleError(err)

	args := os.Args
	if len(args) < 2 {
		fmt.Print("two arguments are required\n")
		os.Exit(1)
	}
	if args[1] == "login" && len(args) == 2 {
		fmt.Print("set username arguments\n")
		os.Exit(1)
	}
	cmd := cli.Command{
		Name : args[1],
		Arguments : args,
	}
	err = cmds.Run(s, cmd)
	config.HandleError(err)
}
