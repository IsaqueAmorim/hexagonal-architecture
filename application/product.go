package application

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	Id     string
	Name   string
	Price  float64
	Status string
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == DISABLED {
		return false, nil
	}
	return true, nil
}

func (p *Product) Enable() error {
	p.Status = ENABLED
	return nil
}

func (p *Product) Disable() error {
	p.Status = DISABLED
	return nil
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
