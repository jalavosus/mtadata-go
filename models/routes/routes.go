package routes

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
)

const (
	routesGormDataTypePg string = "route[]"
)

type Routes []Route

func (r Routes) Proto() (routes []protosv1.Route) {
	routes = make([]protosv1.Route, len(r))

	for i := range r {
		routes[i] = r[i].Proto()
	}

	return
}

func (r *Routes) UnmarshalJSON(data []byte) error {
	var routes []Route
	if err := json.Unmarshal(data, &routes); err != nil {
		return err
	}

	*r = routes

	return nil
}

func (r Routes) Value() (driver.Value, error) {
	var sa = make(pq.StringArray, len(r))

	for i, rt := range r {
		sa[i] = rt.String()
	}

	return sa.Value()
}

func (r *Routes) Scan(value any) error {
	var sa = new(pq.StringArray)

	if err := sa.Scan(value); err != nil {
		return err
	}

	routes := make(Routes, len(*sa))
	for i, rt := range *sa {
		routes[i] = utils.DbValueToEnum(rt, AllRoutes, Unknown)
	}

	*r = routes

	return nil
}

func (Routes) GormDataType() string {
	return routesGormDataTypePg
}

func (Routes) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return routesGormDataTypePg
	default:
		return routesGormDataTypePg
	}
}
