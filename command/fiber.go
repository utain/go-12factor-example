package command

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/utain/go/example/internal/adapters/gormadapter"
	"github.com/utain/go/example/internal/adapters/gormadapter/todos"
	"github.com/utain/go/example/internal/adapters/presenter/fiberserv"
	"github.com/utain/go/example/internal/adapters/viperadapter"
	"github.com/utain/go/example/internal/adapters/zapadapter"
	"github.com/utain/go/example/internal/core"
	"github.com/utain/go/example/internal/logs"
)

func startFiber(opts ServerOpts) (err error) {
	logger := zapadapter.ZapAdapter()

	// prepare datasource connection
	db, err := gormadapter.Connect(gormadapter.GormConfig{
		Url:         viperadapter.V().Database.URL,
		MaxPoolOpen: viperadapter.V().Database.Pool.MaxOpen,
		MaxPoolIdle: viperadapter.V().Database.Pool.MaxIdle,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer gormadapter.Close(db)

	if err = gormadapter.Migration(db); err != nil {
		logger.Error("Migrate database with erro", logs.F{"error": err})
	}

	persistences := core.PersistencesContainer{
		TodoPersistencePort: todos.WithGormPersistence(db, logger),
		//.. other persistence adapters
	}

	// setup servers
	fiberOpts := fiberserv.FiberServerOpts{
		Log:      zapadapter.ZapAdapter(),
		Services: core.ServicesRegister(logger, persistences),
	}

	serv, err := fiberserv.NewFiberServer(fiberOpts)
	if err != nil {
		return err
	}
	return serv.Listen(fmt.Sprintf(":%d", opts.Port))
}

func init() {
	fiberCmd := &cobra.Command{
		Use:   "fiber",
		Short: "Run server with fiber framework",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startFiber(ServerOpts{Port: port})
		},
	}
	root.AddCommand(fiberCmd)
	fiberCmd.PersistentFlags().Uint16VarP(&port, "port", "p", 8000, "Specify server port to start with")
}
