package site

type Store interface {
	Save(site *Site) error
	FindByID(id int64) (*Site, error)
	FindAll() ([]*Site, error)
	Delete(id int64) error
	Update(site *Site) error
}
