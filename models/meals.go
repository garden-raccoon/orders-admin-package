package models

import (
	"github.com/gofrs/uuid"

	proto "github.com/garden-raccoon/orders-admin-package/protocols/orders-admin-package"
)

type MealsDb struct {
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Uuid        uuid.UUID `json:"uuid"`
	Description string    `json:"description"`
	Dummy       *Dummy    `json:"dummy"`
}
type Dummy struct {
	Dummies []string
}

func NewMeals(name, description string, uuid uuid.UUID, price float64, dummy Dummy) *MealsDb {
	meals := &MealsDb{
		Name:        name,
		Price:       price,
		Uuid:        uuid,
		Description: description,
		Dummy:       &dummy,
	}
	return meals
}

// Proto is
func (mdb MealsDb) Proto() *proto.MealDb {

	order := &proto.MealDb{
		Uuid:        mdb.Uuid.Bytes(),
		Name:        mdb.Name,
		Description: mdb.Description,
		Price:       float32(mdb.Price),
	}
	dummy2 := &Dummy{}
	mdb.Dummy = dummy2
	dummy := &proto.Dummy{Dummies: mdb.Dummy.Dummies}
	order.Dummy = dummy
	return order
}

func MealDbFromProto(pb *proto.MealDb) *MealsDb {

	order := &MealsDb{
		Name:        pb.Name,
		Price:       float64(pb.Price),
		Description: pb.Description,
		Uuid:        uuid.FromBytesOrNil(pb.Uuid),
	}
	dummy2 := &proto.Dummy{}
	pb.Dummy = dummy2
	dummy := &Dummy{Dummies: pb.Dummy.Dummies}
	order.Dummy = dummy
	return order
}

func MealsToProto(meals []*MealsDb) *proto.MealsDb {
	pb := &proto.MealsDb{}

	for _, b := range meals {
		pb.MealsDb = append(pb.MealsDb, b.Proto())
	}
	return pb
}

func MealsFromProto(pb *proto.MealsDb) (meals []*MealsDb) {
	for _, b := range pb.MealsDb {
		meals = append(meals, MealDbFromProto(b))
	}
	return
}
