type Product interface {
	GetName() string
}

type ProductA struct {
	name string
}

func (p *ProductA) GetName() string {
	return p.name
}

type ProductB struct {
	name string
}

func (p *ProductB) GetName() string {
	return p.name
}

type Creator interface {
	CreateProduct(name string) Product
}

type sCreator struct{}

func (c *sCreator) CreateProduct(name string) Product {
	if name == "A" {
		return &ProductA{name: "Product A"}
	} else if name == "B" {
		return &ProductB{name: "Product B"}
	} else {
		return nil
	}
}

func main() {
	creator := &sCreator{}
	productA := creator.CreateProduct("A")
	productB := creator.CreateProduct("B")
	fmt.Println(productA.GetName())
	fmt.Println(productB.GetName())
}
