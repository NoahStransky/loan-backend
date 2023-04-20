package model

type AccountingProvider struct {
	Id        float64 `xorm:"varchar(256) not null unique 'id' comment('accounting_provider_id')" json:"id"`
	Name      string  `xorm:"varchar(256) not null unique 'name' comment('accounting_provider_name')" json:"name"`
	IsDeleted bool    `xorm:"boolean not null default false" json:"is_deleted"`
}
