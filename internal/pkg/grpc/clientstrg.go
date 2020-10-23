package grpc

import (
	"context"
	api "gRPC/api/protoc"
	"gRPC/internal/pkg/data"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type StorageGrpcClient struct {
	Client api.ServiceProtobufClient
}

func NewGrpcClient() *StorageGrpcClient {
	conn, err := grpc.Dial("grpc:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connect to grpc localhost:9000")
	client := api.NewServiceProtobufClient(conn)

	return &StorageGrpcClient{
		Client: client,
	}
}

func (s *StorageGrpcClient) Add(ad *data.Ad) error {
	request := api.RequestAdd{
		Ad: &api.Ad{
			Model: ad.GetModel(),
			Brand: ad.GetBrand(),
			Color: ad.GetColor(),
			Price: int32(ad.GetPrice()),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := s.Client.Add(ctx, &request)
	if err != nil {
		return err
	}

	// cl := response.GetStatus()

	return nil
}

func (s *StorageGrpcClient) Get(id uint) (*data.Ad, error) {
	request := api.RequestID{Id: uint32(id)}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := s.Client.Get(ctx, &request)
	if err != nil {
		return nil, err
	}

	ad := response.Ads[0]

	return &data.Ad{
		ID:    uint(ad.GetId()),
		Brand: ad.GetBrand(),
		Model: ad.GetModel(),
		Color: ad.GetColor(),
		Price: int(ad.GetPrice()),
	}, nil
}

func (s *StorageGrpcClient) GetAll() ([]*data.Ad, error) {
	request := new(empty.Empty)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := s.Client.GetAll(ctx, request)
	if err != nil {
		return nil, err
	}

	ads := response.GetAds()
	if err != nil {
		return nil, err
	}
	baseAd := make([]*data.Ad, 0, len(ads))
	for _, v := range ads {
		baseAd = append(baseAd, &data.Ad{
			ID:    uint(v.GetId()),
			Brand: v.GetBrand(),
			Model: v.GetModel(),
			Color: v.GetColor(),
			Price: int(v.GetPrice()),
		})
	}

	return baseAd, nil
}

func (s *StorageGrpcClient) Update(temp *data.Ad, id uint) error {
	request := api.RequestUpdateAd{
		Id: uint32(id),
		Ad: &api.Ad{
			Brand: temp.GetBrand(),
			Model: temp.GetModel(),
			Color: temp.GetColor(),
			Price: int32(temp.GetPrice())},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := s.Client.Update(ctx, &request)
	if err != nil {
		return err
	}
	// status := response.GetStatus()

	return nil
}

func (s *StorageGrpcClient) Delete(id uint) error {
	request := api.RequestID{Id: uint32(id)}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := s.Client.Delete(ctx, &request)
	if err != nil {
		return err
	}
	// status := response.GetStatus()

	return nil
}

func (s *StorageGrpcClient) Size() (int, error) {
	request := new(empty.Empty)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := s.Client.Size(ctx, request)
	if err != nil {
		return 0, err
	}

	return int(response.GetSize()), nil
}

func (s *StorageGrpcClient) AddAccount(acc *data.Account) error {
	request := api.Account{
		Username: acc.GetUserName(),
		Password: acc.GetPassword(),
		Token:    acc.GetToken(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := s.Client.AddAccount(ctx, &request)
	if err != nil {
		return err
	}

	// status := response.GetStatus()
	return nil
}

func (s *StorageGrpcClient) GetAccounts() ([]*data.Account, error) {
	request := new(empty.Empty)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := s.Client.GetAccounts(ctx, request)
	if err != nil {
		return nil, err
	}

	accs := response.GetAcc()

	baseAcc := make([]*data.Account, 0, len(accs))

	for _, v := range accs {
		baseAcc = append(baseAcc, &data.Account{
			Username: v.GetUsername(),
			Password: v.GetPassword(),
			Token:    v.GetToken(),
		})
	}

	return baseAcc, nil
}
