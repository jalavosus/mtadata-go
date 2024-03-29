package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"

	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/internal/database"
	"github.com/jalavosus/mtadata/internal/database/dbconn"
	"github.com/jalavosus/mtadata/internal/logging"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structures"
)

func makeDropTypeCmd(typeName string) string {
	return fmt.Sprintf("DROP TYPE public.%[1]s;", typeName)
}

var (
	runMigrateCmd = cli.Command{
		Name:   "run-all",
		Usage:  "Run all migrations",
		Action: migrateCmdAction,
		Flags: []cli.Flag{
			&configFlag,
		},
	}
	dropAllCmd = cli.Command{
		Name:   "drop-all",
		Usage:  "Drop all tables and types",
		Action: dropAllCmdAction,
	}
)

var (
	configFlag = cli.PathFlag{
		Name:     "config",
		Aliases:  []string{"c"},
		Required: false,
	}
)

var dbModels = []database.CustomDbTyper{
	boroughs.Borough(0),
	structures.Structure(0),
	divisions.Division(0),
	routes.Route(0),
	models.GtfsLocation{},
	models.DirectionLabels{},
	models.StationInfo{},
}

var tableNames = []string{
	"stations",
	"station_complexes",
}

var typeMigrations = []any{
	models.Station{},
	models.StationComplex{},
}

func migrateRunAll(lc fx.Lifecycle) {
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		conn := dbconn.ConnectionContext(ctx)

		for _, tableName := range tableNames {
			cmd := "DROP TABLE " + tableName
			if err := conn.Exec(cmd).Error; err != nil {
				log.Println(err)
			}
		}

		for _, dbModel := range dbModels {
			if err := dropType(conn, dbModel.GormDataType()); err != nil {
				log.Println(err)
			}
		}

		for _, dbModel := range dbModels {
			cmd := dbModel.CreateDbType()
			if err := conn.Exec(cmd).Error; err != nil {
				return errors.WithMessagef(err, "error executing sql '%[1]s'", cmd)
			}
		}

		for _, tm := range typeMigrations {
			if err := conn.AutoMigrate(&tm); err != nil {
				return err
			}
		}

		return nil
	}})
}

func migrateCmdAction(c *cli.Context) error {

	app := fx.New(
		fx.Supply(configFlag.Get(c)),
		fx.Provide(logging.NewLogger),
		logging.WithLogger,
		fx.Module("config", config.Module),
		fx.Module("database", database.Module),
		fx.Invoke(migrateRunAll),
	)

	if err := app.Start(c.Context); err != nil {
		return err
	}

	for {
		select {
		case <-app.Done():
			return app.Err()
		}
	}
}

func dropAllCmdAction(c *cli.Context) error {
	ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
	defer cancel()

	conn := dbconn.ConnectionContext(ctx)

	for _, tableName := range tableNames {
		cmd := "DROP TABLE " + tableName
		if err := conn.Exec(cmd).Error; err != nil {
			return err
		}
	}

	for _, dbModel := range dbModels {
		if err := dropType(conn, dbModel.GormDataType()); err != nil {
			return err
		}
	}

	return nil
}

func dropType(conn *gorm.DB, typeName string) error {
	cmd := makeDropTypeCmd(typeName)
	if err := conn.Exec(cmd).Error; err != nil {
		return errors.WithMessagef(err, "error dropping type %[1]s", typeName)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "migrate",
		Usage: "Run database migrations",
		Commands: []*cli.Command{
			&runMigrateCmd,
			&dropAllCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
