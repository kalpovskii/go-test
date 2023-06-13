package socksgorm

import (
	"awesomeProject/internal/repositories/socksrepository"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type SocksGorm struct {
	db *gorm.DB
}

func New(db *gorm.DB) *SocksGorm {
	return &SocksGorm{
		db: db,
	}
}

func (r *SocksGorm) Create(ctx context.Context, s *socksrepository.Socks) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *SocksGorm) Get(ctx context.Context, s *socksrepository.Socks, operation string) (int64, error) {
	validOperations := map[string]bool{
		"moreThan": true,
		"lessThan": true,
		"equal":    true,
	}
	if !validOperations[operation] {
		return 0, errors.New("incorrect operation")
	}

	query := r.db.WithContext(ctx).Model(&socksrepository.Socks{}).Where("color = ?", s.Color)
	switch operation {
	case "moreThan":
		query = query.Where("cotton_part > ?", s.CottonPart)
	case "lessThan":
		query = query.Where("cotton_part < ?", s.CottonPart)
	case "equal":
		query = query.Where("cotton_part = ?", s.CottonPart)
	}

	var totalCount int64
	err := query.Count(&totalCount).Error
	if err != nil {
		return 0, fmt.Errorf("error while count: %v", err)
	}

	return totalCount, nil
}

func (r *SocksGorm) Delete(ctx context.Context, s *socksrepository.Socks) error {
	query := r.db.WithContext(ctx).Model(&socksrepository.Socks{}).
		Where("color = ?", s.Color).
		Where("cotton_part = ?", s.CottonPart).
		UpdateColumn("quantity", gorm.Expr("quantity - ?", s.Quantity))

	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return fmt.Errorf("no socks matching the criteria")
	}

	return nil
}
