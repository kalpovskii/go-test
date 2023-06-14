package socksrepository

import "context"

type Socks struct {
	ID         int64  `gorm:"id"`
	Color      string `gorm:"color"`
	CottonPart int64  `gorm:"cottonPart"`
	Quantity   int64  `gorm:"quantity"`
}

type SocksRepository interface {
	Create(ctx context.Context, s *Socks) error
	Get(ctx context.Context, s *Socks, operation string) (int64, error)
	Delete(ctx context.Context, s *Socks) error
}
