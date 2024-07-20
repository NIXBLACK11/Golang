package main

type startable interface {
	Start()
}

func startEngine(cars ...startable) {
	for _, car := range cars {
		car.Start()
	}
}

func main() {
	// we can use enhanced as well as normal transmission
	myConvertable := Convertable{Engine{}, EnhancedTransmission{}, SteeringWheel{}}
	myTruck := Truck{Engine{}, Transmission{}, SteeringWheel{}}

	startEngine(myConvertable, myTruck)

	myTruck.ShiftDown()
	myConvertable.ShiftUp()

	// myTruck.Start()
	// myConvertable.Start()

	// myConvertable.ConvertDown()
	// myConvertable.ConvertTop()

	// myTruck.FourWheelDrive()
}