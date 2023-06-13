package socksrepository

import "context"

type Socks struct {
	ID         int64  `gorm:"id"`
	Color      string `gorm:"color"`
	CottonPart int64  `gorm:"cottonPart"`
	Quantity   int64  `gorm:"quantity"`
}

type SocksRepository interface {
	Create(ctx context.Context, e *Socks) error
	List(ctx context.Context, color string, cottonPart int64, quantity int64) ([]*Socks, error)
	Delete(ctx context.Context, color string, cottonPart int64, quantity int64) error
}
