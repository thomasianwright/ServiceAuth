package models

type Scope struct {
	Id           int64          `json:"id"`
	Name         string         `json:"name"`
	Applications []*Application `json:"applications" gorm:"many2many:application_scopes;"`
}
