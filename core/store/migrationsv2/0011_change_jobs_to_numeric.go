package migrationsv2

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

const (
	up11 = `
ALTER TABLE job_specs ALTER COLUMN min_payment TYPE numeric(78, 0) USING min_payment::numeric;
ALTER TABLE flux_monitor_specs ALTER COLUMN min_payment TYPE numeric(78, 0) USING min_payment::numeric;
`
	down11 = `
ALTER TABLE job_specs ALTER COLUMN min_payment TYPE varchar(255) USING min_payment::varchar;
ALTER TABLE flux_monitor_specs ALTER COLUMN min_payment TYPE varchar(255) USING min_payment::varchar;
`
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "0011_change_jobs_to_numeric",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(up11).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down11).Error
		},
	})
}
