package api

import "context"

type Server struct{}

func (s *Server) mustEmbedUnimplementedServiceServer() {
	panic("implement me")
}

func (s *Server) Sum(ctx context.Context, r *Request) (*Response, error) {
	return &Response{Sum: r.A + r.B}, nil
}
