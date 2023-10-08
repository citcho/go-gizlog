package dao

import (
	"time"

	"github.com/citcho/go-gizlog/internal/user/infrastructure/dao"
)

type Report struct {
	ID            string    `bun:",pk,type:char(26)"`
	User          *dao.User `bun:",rel:belongs-to"`
	UserID        string    `bun:",type:varchar(26),notnull"`
	Content       string    `bun:",type:text,notnull"`
	ReportingTime time.Time `bun:",type:date,notnull"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `bun:",soft_delete,default:null"`
}
