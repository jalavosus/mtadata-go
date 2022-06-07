package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"

	"github.com/jalavosus/mtadata/internal/database/connection"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/borough"
	"github.com/jalavosus/mtadata/models/division"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structure"

	_ "github.com/joho/godotenv/autoload"
)

func makeDropTypeCmd(typeName string) string {
	return fmt.Sprintf("DROP TYPE public.%[1]s;", typeName)
}

func migrateCmdAction(c *cli.Context) error {
	ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
	defer cancel()

	conn := connection.Connection()
	conn = conn.WithContext(ctx)

	var (
		boroughModel         = borough.Borough("")
		structureModel       = structure.Structure("")
		divisionModel        = division.Division("")
		routeModel           = routes.Route("")
		gtfsLocationModel    = models.GtfsLocation{}
		directionLabelsModel = models.DirectionLabels{}
		// routesModel          = routes.Routes{}
	)

	if err := conn.Exec("DROP TABLE stations;").Error; err != nil {
		log.Println(err)
	}

	if err := dropType(conn, boroughModel.GormDataType()); err != nil {
		log.Println(err)
	}

	if err := dropType(conn, structureModel.GormDataType()); err != nil {
		log.Println(err)
	}

	if err := dropType(conn, divisionModel.GormDataType()); err != nil {
		log.Println(err)
	}

	if err := dropType(conn, routeModel.GormDataType()); err != nil {
		log.Println(err)
	}

	if err := dropType(conn, gtfsLocationModel.GormDataType()); err != nil {
		log.Println(err)
	}

	if err := dropType(conn, directionLabelsModel.GormDataType()); err != nil {
		log.Println(err)
	}

	cmd := boroughModel.CreateDbType()
	if err := conn.Exec(cmd).Error; err != nil {
		return errors.WithMessagef(err, "error executing sql '%[1]s'", cmd)
	}

	cmd = structureModel.CreateDbType()
	if err := conn.Exec(cmd).Error; err != nil {
		return errors.WithMessagef(err, "error executing sql '%[1]s'", cmd)
	}

	cmd = divisionModel.CreateDbType()
	if err := conn.Exec(cmd).Error; err != nil {
		return errors.WithMessagef(err, "error executing sql '%[1]s'", cmd)
	}

	cmd = routeModel.CreateDbType()
	if err := conn.Exec(cmd).Error; err != nil {
		return errors.WithMessagef(err, "error executing sql '%[1]s'", cmd)
	}

	cmd = gtfsLocationModel.CreateDbType()
	if err := conn.Exec(cmd).Error; err != nil {
		return errors.WithMessagef(err, "error executing sql '%[1]s'", cmd)
	}

	cmd = directionLabelsModel.CreateDbType()
	if err := conn.Exec(cmd).Error; err != nil {
		return errors.WithMessagef(err, "error executing sql '%[1]s'", cmd)
	}

	if err := conn.AutoMigrate(&models.Station{}); err != nil {
		return err
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
		Name:   "migrate",
		Usage:  "Run database migrations",
		Action: migrateCmdAction,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
