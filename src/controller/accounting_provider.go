package controller

import (
	"bernie.top/demyst/loan-backend/src/model"
	"xorm.io/xorm"
)

type AccountingProvider struct {
	engine *xorm.Engine
}

func NewAccountingProvider(engine *xorm.Engine) *AccountingProvider {
	return &AccountingProvider{
		engine: engine,
	}
}

func (c *AccountingProvider) GetProviders() ([]model.AccountingProvider, error) {
	var providers []model.AccountingProvider
	err := c.engine.Find(&providers)
	if err != nil {
		return nil, err
	}
	return providers, err
}
