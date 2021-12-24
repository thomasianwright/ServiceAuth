package models

type Application struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	Status    string   `json:"status"`
	ApiKey    string   `json:"-"`
	PublicKey string   `json:"publicKey"`
	Enabled   bool     `json:"enabled"`
	Scopes    []*Scope `json:"scopes" gorm:"many2many:application_scopes;"`
}
