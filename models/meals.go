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
}

func NewMeals(name, description string, uuid uuid.UUID, price float64) *MealsDb {
	return &MealsDb{
		Name:        name,
		Price:       price,
		Uuid:        uuid,
		Description: description,
	}
}

// Proto is
func (mdb MealsDb) Proto() *proto.MealDb {

	order := &proto.MealDb{
		Uuid:        mdb.Uuid.Bytes(),
		Name:        mdb.Name,
		Description: mdb.Description,
		Price:       float32(mdb.Price),
	}
	return order
}

func MealDbFromProto(pb *proto.MealDb) *MealsDb {

	order := &MealsDb{
		Name:        pb.Name,
		Price:       float64(pb.Price),
		Description: pb.Description,
	}
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
