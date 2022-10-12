package restaurant

import "github.com/dzendos/dubna/internal/model/menu"

type PaymentMethod int

const (
	ViaTelephone PaymentMethod = iota + 1
	ViaMessenger
	ViaSite
)

type Restaurant struct {
	Name            string
	Reference       string
	ImageUrl        string
	TelephoneNumber string
	PaymentMethod   PaymentMethod
	Menu            *menu.Menu
}
