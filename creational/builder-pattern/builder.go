package main

// Car represents the car object.
type Car struct {
	color         string
	engineType    string
	hasSunroof    bool
	hasNavigation bool
}

// CarBuilder is the interface for building a car with custom features.
type CarBuilder interface {
	SetColor(color string) CarBuilder
	SetEngineType(engineType string) CarBuilder
	SetSunroof(hasSunroof bool) CarBuilder
	SetNavigation(hasNavigation bool) CarBuilder
	Build() *Car
}

// carBuilder is the object that build the car.
type carBuilder struct {
	car *Car
}

// Director is the object that manage the car builder system.
type Director struct {
	builder CarBuilder
}

// NewCarBuilder creates a carBuilder instance.
func NewCarBuilder() CarBuilder {
	return &carBuilder{
		car: &Car{}, // Initialize the car attribute
	}
}

// SetColor sets the car color.
func (cb *carBuilder) SetColor(color string) CarBuilder {
	cb.car.color = color
	return cb
}

// SetEngineType configures the car engine type.
func (cb *carBuilder) SetEngineType(engineType string) CarBuilder {
	cb.car.engineType = engineType
	return cb
}

// SetSunroof configures the car with sunroof.
func (cb *carBuilder) SetSunroof(hasSunroof bool) CarBuilder {
	cb.car.hasSunroof = hasSunroof
	return cb
}

// SetNavigation configures the car with a navigation system.
func (cb *carBuilder) SetNavigation(hasNavigation bool) CarBuilder {
	cb.car.hasNavigation = hasNavigation
	return cb
}

// Build is the method for building the car.
func (cb *carBuilder) Build() *Car {
	return cb.car
}

// ConstructCar constructs the car.
func (d *Director) ConstructCar(color, engineType string, hasSunroof, hasNavigation bool) *Car {
	d.builder.SetColor(color).
		SetEngineType(engineType).
		SetSunroof(hasSunroof).
		SetNavigation(hasNavigation)
	return d.builder.Build()
}

func main() {
	// Create a new car builder.
	builder := NewCarBuilder()
	// Create a car with the director.
	director := &Director{builder: builder}
	_ = director.ConstructCar("blue", "electric", true, true) // this is myCar
	// Use the car object with the chosen configuration.
	// ...
}
