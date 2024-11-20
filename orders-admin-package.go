package orders

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/garden-raccoon/orders-admin-package/models"
	proto "github.com/garden-raccoon/orders-admin-package/protocols/orders-admin-package"
)

// TODO need to set timeout via lib initialisation
// timeOut is  hardcoded GRPC requests timeout value
const timeOut = 60

// Debug on/off
var Debug = true

// IOrderAdminPkgAPI is
type IOrderAdminPkgAPI interface {
	GetOrders() ([]*models.OrderAdmin, error)
	AllOrders(pb *proto.OrdersAdmin) ([]*models.OrderAdmin, error)

	OrderByName(name string) (*models.OrderAdmin, error)

	CreateOrder(s *models.OrderAdmin) error
	// Close GRPC Api connection
	Close() error
}

// Api is profile-service GRPC Api
// structure with client Connection
type Api struct {
	addr    string
	timeout time.Duration
	*grpc.ClientConn
	proto.OrderServiceClient
}

// New create new Battles Api instance
func New(addr string) (IOrderAdminPkgAPI, error) {
	api := &Api{timeout: timeOut * time.Second}

	if err := api.initConn(addr); err != nil {
		return nil, fmt.Errorf("create Battles Api:  %w", err)
	}

	api.OrderServiceClient = proto.NewOrderServiceClient(api.ClientConn)
	return api, nil
}
func (api *Api) AllOrders(pb *proto.OrdersAdmin) ([]*models.OrderAdmin, error) {
	ppp := models.OrdersFromProto(pb)
	return ppp, nil
}
func (api *Api) GetOrders() ([]*models.OrderAdmin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	var resp *proto.OrdersAdmin
	empty := new(proto.OrderAdminEmpty)
	resp, err := api.OrderServiceClient.GetOrders(ctx, empty)
	if err != nil {
		return nil, fmt.Errorf("GetOrders api request: %w", err)
	}

	orders := models.OrdersFromProto(resp)

	if Debug {
		fmt.Println(orders)
	}
	return orders, nil
}
func (api *Api) CreateOrder(s *models.OrderAdmin) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	var order = &proto.OrderAdmin{

		Uuid:     s.Uuid.Bytes(),
		Name:     s.Name,
		Price:    float32(s.Price),
		Quantity: int32(s.Quantity),
		Day:      s.Day,
		MealType: s.MealType,
	}
	// DEBUG Info
	if Debug {
		fmt.Println("order is")
		fmt.Println(order)

		if order == nil {
			fmt.Println("order is nil")
		} else {
			fmt.Println("order is not nil")
		}
		if api.OrderServiceClient == nil {
			fmt.Println("OrderServiceClient is nil")
		} else {
			fmt.Println("OrderServiceClient is not nil")
		}
	}
	//--- **** ---///

	_, err = api.OrderServiceClient.CreateOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("create order api request: %w", err)
	}
	return nil
}

// initConn initialize connection to Grpc servers
func (api *Api) initConn(addr string) (err error) {
	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	api.ClientConn, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	return
}

// OrderByName is
func (api *Api) OrderByName(name string) (*models.OrderAdmin, error) {
	getter := &proto.OrderAdminGetter{
		Getter: &proto.OrderAdminGetter_Name{
			Name: name,
		},
	}
	return api.getOrder(getter)
}

// getOrder is
func (api *Api) getOrder(getter *proto.OrderAdminGetter) (*models.OrderAdmin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	resp, err := api.OrderServiceClient.OrderByName(ctx, getter)
	if err != nil {
		return nil, fmt.Errorf("get battle api request: %w", err)
	}

	return models.OrderAdminFromProto(resp), nil
}
