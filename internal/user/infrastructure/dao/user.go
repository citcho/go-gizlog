package dao

import "time"

type User struct {
	ID        string    `bun:",pk,type:char(26),notnull"`
	Name      string    `bun:",notnull"`
	Email     string    `bun:",notnull,unique"`
	Password  string    `bun:",notnull"`
	CreatedAt time.Time `bun:",notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete,default:null"`
}
