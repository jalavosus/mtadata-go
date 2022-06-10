package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
)

const (
	stationInfoGormDataTypePg  string = "station_info"
	stationInfosGormDataTypePg string = "station_info[]"
)

type StationInfo struct {
	StationId  string `json:"station_id" yaml:"station_id"`
	GtfsStopId string `json:"gtfs_stop_id" yaml:"gtfs_stop_id"`
}

func (StationInfo) GormDataType() string {
	return stationInfoGormDataTypePg
}

func (StationInfo) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return stationInfoGormDataTypePg
	default:
		return stationInfoGormDataTypePg
	}
}

func (StationInfo) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS (
	station_id TEXT,
	gtfs_stop_id TEXT
);`, stationInfoGormDataTypePg)
}

func (d *StationInfo) Scan(value any) error {
	val := value.(string)
	val = utils.TrimParens(val)

	split := utils.TrimWhitespaceSlice(strings.Split(val, ","))

	d.StationId = split[0]
	d.GtfsStopId = split[1]

	return nil
}

func (d StationInfo) Value() (driver.Value, error) {
	return fmt.Sprintf(`("%[1]s","%[2]s")`, d.StationId, d.GtfsStopId), nil
}

type (
	StationInfos    []StationInfo
	RawStationInfos []StationInfo
)

func (s StationInfos) Value() (driver.Value, error) {
	var sa = make(pq.StringArray, len(s))

	for i, si := range s {
		val, _ := si.Value()
		sa[i] = val.(string)
	}

	return sa.Value()
}

func (StationInfos) GormDataType() string {
	return stationInfosGormDataTypePg
}

func (StationInfos) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return stationInfosGormDataTypePg
	default:
		return stationInfosGormDataTypePg
	}
}

var (
	stationInfosRegex = regexp.MustCompile(`([(].{1,4},.{1,4}[)])`)
)

func (s *StationInfos) Scan(value any) error {
	var infos StationInfos

	switch val := value.(type) {
	case StationInfos:
		*s = val
	case []byte:
		if err := json.Unmarshal(val, &infos); err != nil {
			return err
		}

		*s = infos
	case string:
		infoStrings := stationInfosRegex.FindAllString(val, -1)
		for _, s := range infoStrings {
			stationInfo := new(StationInfo)
			if err := stationInfo.Scan(s); err != nil {
				return err
			}

			if stationInfo != nil && stationInfo.GtfsStopId != "" {
				infos = append(infos, *stationInfo)
			}
		}

		*s = infos
	default:
		fmt.Printf("%T\n", value)
	}

	return nil
}
