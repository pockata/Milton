package storage

import (
	"context"
	"database/sql"
	models "milton/generated_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type FlowerPot struct {
	db  *sql.DB
	pot *models.FlowerPot
}

func (p *FlowerPot) ID() string {
	return p.pot.ID
}

func (p *FlowerPot) Name() string {
	return p.pot.Name
}

func (p *FlowerPot) Update(name string) error {
	ctx := context.Background()

	p.pot.Name = name

	_, err := p.pot.Update(ctx, p.db, boil.Whitelist(models.FlowerPotColumns.Name))

	if err != nil {
		return err
	}

	return nil
}
