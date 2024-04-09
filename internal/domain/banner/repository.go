package banner

import "context"

type Repository interface {
	GetAll() ([]Banner, error)
	GetByID(ctx context.Context, bannerID int) (Banner, error)
	Add(banner Banner) error
	Update(banner Banner) error
	Delete(id int) error
}
