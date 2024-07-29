package Cars

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	MercedesBrand = iota
	VolksvagenBrand
	LadaBrand
)

type Car interface {
	GetBrand() int
	GetMaterial() string
	GetColor() string
	Print()
}

func CreateCar(brand int, color, material string) (Car, error) {
	if len(color) == 0 || len(material) == 0 {
		return nil, errors.New("color and material must both be non-empty")
	}

	var car Car
	switch brand {
	case VolksvagenBrand:
		car = &Volksvagen{
			brand:    VolksvagenBrand,
			material: material,
			color:    color,
		}
		break
	case MercedesBrand:
		car = &Mercedes{
			brand:    MercedesBrand,
			material: material,
			color:    color,
		}
		break
	case LadaBrand:
		car = &Lada{
			brand:    LadaBrand,
			material: material,
			color:    color,
		}
		break
	default:
		return nil, errors.New("The brand should be one of the values: 'MercedesBrand', 'VolksvagenBrand', 'LadaBrand'")
	}

	return car, nil
}

func CreateRandomVolksvagen() *Volksvagen {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	colors := [5]string{"Red", "Green", "Blue", "Purple", "Yellow"}
	materials := [5]string{"Steel", "Composite", "Steel Alloy", "Alloy-10", "Alloy-20"}

	car := Volksvagen{}
	car.color = colors[r.Intn(len(colors))]
	car.material = materials[r.Intn(len(materials))]
	car.brand = VolksvagenBrand

	return &car
}

func CreateRandomMercedes() *Mercedes {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	colors := [5]string{"Red", "Green", "Blue", "Purple", "Yellow"}
	materials := [5]string{"Steel", "Composite", "Steel Alloy", "Alloy-10", "Alloy-20"}

	car := Mercedes{}
	car.color = colors[r.Intn(len(colors))]
	car.material = materials[r.Intn(len(materials))]
	car.brand = MercedesBrand

	return &car
}

func CreateRandomLada() *Lada {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	colors := [5]string{"Red", "Green", "Blue", "Purple", "Yellow"}
	materials := [5]string{"Steel", "Composite", "Steel Alloy", "Alloy-10", "Alloy-20"}

	car := Lada{}
	car.color = colors[r.Intn(len(colors))]
	car.material = materials[r.Intn(len(materials))]
	car.brand = LadaBrand

	return &car
}

func CreateCars(totalCarCount, VolksvagenCount, MercedesCount, LadaCount int) []Car {
	cars := make([]Car, totalCarCount)
	for i := 0; i < totalCarCount; i++ {
		time.Sleep(10 * time.Nanosecond)
		if VolksvagenCount > 0 {
			cars[i] = CreateRandomVolksvagen()
			VolksvagenCount--
		} else if MercedesCount > 0 {
			cars[i] = CreateRandomMercedes()
			MercedesCount--
		} else if LadaCount > 0 {
			cars[i] = CreateRandomLada()
			LadaCount--
		}
	}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	r.Shuffle(len(cars), func(i, j int) { cars[i], cars[j] = cars[j], cars[i] })
	return cars
}

type Volksvagen struct {
	brand    int
	material string
	color    string
}

func (v *Volksvagen) GetBrand() int {
	return v.brand
}

func (v *Volksvagen) GetMaterial() string {
	return v.material
}

func (v *Volksvagen) GetColor() string {
	return v.color
}

func (v *Volksvagen) Print() {
	fmt.Printf("Brand: Volksvagen\nMaterial: %s\nColor: %s\n", v.material, v.color)
}

type Mercedes struct {
	brand    int
	material string
	color    string
}

func (m *Mercedes) GetBrand() int {
	return m.brand
}

func (m *Mercedes) GetMaterial() string {
	return m.material
}

func (m *Mercedes) GetColor() string {
	return m.color
}

func (m *Mercedes) Print() {
	fmt.Printf("Brand: Mercedes\nMaterial: %s\nColor: %s\n", m.material, m.color)
}

type Lada struct {
	brand    int
	material string
	color    string
}

func (l *Lada) GetBrand() int {
	return l.brand
}

func (l *Lada) GetMaterial() string {
	return l.material
}

func (l *Lada) GetColor() string {
	return l.color
}

func (l *Lada) Print() {
	fmt.Printf("Brand: Lada\nMaterial: %s\nColor: %s\n", l.material, l.color)
}
