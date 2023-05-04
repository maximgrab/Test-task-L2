type Car struct {
	Brand string
	Model string
	Year  string
}

type CarBuilder interface {
	SetBrand(string) carBuilder
	SetModel(string) carBuilder
	SetYear(string) carBuilder
	build() Car
}
type carBuilder struct {
	car Car
}

func (cb *carBuilder) SetBrand(brand string) carBuilder {
	cb.car.Brand = brand
}

func (cb *carBuilder) SetModel(model string) carBuilder {
	cb.car.Model = model
}

func (cb *carBuilder) SetYear(year string) carBuilder {
	cb.car.Year = year
}
