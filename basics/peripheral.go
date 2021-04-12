package basics

type (
	// Peripheral represents a basic infrastructure which can
	// be initialized and destroyed.
	//
	// For a Peripheral, the host should add it into a list
	// and destroy them while host is shutting down.
	Peripheral interface {
		// Close provides a closer to cleanup the peripheral gracefully
		Close()
	}
)

// Basic is a base type to simplify your codes since you're using Peripheral type.
type Basic struct {
	peripherals []Peripheral
}

// AddPeripheral adds a Peripheral object into Basic holder/host.
func (s *Basic) AddPeripheral(peripherals ...Peripheral) {
	s.peripherals = append(s.peripherals, peripherals...)
}

// Close provides a closer to cleanup the peripheral gracefully
func (s *Basic) Close() {
	for _, p := range s.peripherals {
		if p != nil {
			p.Close()
		}
	}
	s.peripherals = nil
}
