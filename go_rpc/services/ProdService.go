package services

import "context"

type ProdService struct {
}

func (p *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}

func (p *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	return &ProdResponseList{Prodres: []*ProdResponse{
		&ProdResponse{ProdStock: 20},
		&ProdResponse{ProdStock: 21},
		&ProdResponse{ProdStock: 22},
		&ProdResponse{ProdStock: 23},
	}}, nil
}
