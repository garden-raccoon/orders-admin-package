package models

import (
	"github.com/gofrs/uuid"

	proto "github.com/garden-raccoon/orders-admin-package/protocols/orders-admin-package"
)

// Order is
type OrderAdmin struct {
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Uuid     uuid.UUID `json:"uuid"`
	Day      string    `json:"day"`
	MealType string    `json:"mealType"`
}

func NewOrder(name, day, mealType string, uuid uuid.UUID, price float64, quantity int) *OrderAdmin {
	return &OrderAdmin{
		Name:     name,
		Price:    price,
		Uuid:     uuid,
		Day:      day,
		MealType: mealType,
	}
}

// Proto is
func (o OrderAdmin) Proto() *proto.OrderAdmin {

	order := &proto.OrderAdmin{
		Uuid:     o.Uuid.Bytes(),
		Name:     o.Name,
		MealType: o.MealType,
		Price:    float32(o.Price),
		Day:      o.Day,
	}
	return order
}

func OrderAdminFromProto(pb *proto.OrderAdmin) *OrderAdmin {

	order := &OrderAdmin{
		Name:     pb.Name,
		Price:    float64(pb.Price),
		Day:      pb.Day,
		MealType: pb.MealType,
		Uuid:     uuid.FromBytesOrNil(pb.Uuid),
	}
	return order
}

// OrdersToProto is
func OrdersToProto(orders []*OrderAdmin) *proto.OrdersAdmin {
	pb := &proto.OrdersAdmin{}

	for _, b := range orders {
		pb.OrdersAdmin = append(pb.OrdersAdmin, b.Proto())
	}
	return pb
}

// OrdersFromProto is
func OrdersFromProto(pb *proto.OrdersAdmin) (orders []*OrderAdmin) {
	for _, b := range pb.OrdersAdmin {
		orders = append(orders, OrderAdminFromProto(b))
	}
	return
}
