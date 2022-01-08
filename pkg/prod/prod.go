package prod

import (
	"context"
	"github.com/shijting/grpcpro/pgfiles"
)

type prodServer struct {
	pgfiles.UnimplementedProdServerServer
}

func NewProdServer() *prodServer {
	return &prodServer{}
}

func (p *prodServer) GetProd(ctx context.Context, req *pgfiles.ProdRequest) (*pgfiles.ProdResponse, error)  {
	return &pgfiles.ProdResponse{
		Id:   req.GetId(),
		Name: "test",
	}, nil
}
