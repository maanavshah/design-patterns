package main

import "fmt"

// Seat interface
type Seat interface {
	GetDetails() string
	GetPrice() int
}

// Concrete seat types
type EconomySeat struct{}
func (e *EconomySeat) GetDetails() string { return "Economy: Basic seat, no extras" }
func (e *EconomySeat) GetPrice() int { return 100 }

type PremiumSeat struct{}
func (p *PremiumSeat) GetDetails() string { return "Premium: Comfortable seat, extra legroom" }
func (p *PremiumSeat) GetPrice() int { return 200 }

type VIPSeat struct{}
func (v *VIPSeat) GetDetails() string { return "VIP: Luxury recliner, complimentary snacks" }
func (v *VIPSeat) GetPrice() int { return 500 }

// Box Office Factory
func BookSeat(seatType string) Seat {
	switch seatType {
	case "economy":
		return &EconomySeat{}
	case "premium":
		return &PremiumSeat{}
	case "vip":
		return &VIPSeat{}
	default:
		return &EconomySeat{} // default to economy
	}
}

func main() {
	// Customer booking different seat types
	bookings := []string{"economy", "premium", "vip"}
	
	for _, booking := range bookings {
		seat := BookSeat(booking)  // Factory creates right seat type
		fmt.Printf("%s - Price: ₹%d\n", seat.GetDetails(), seat.GetPrice())
	}
}
