package server

import (
	"context"

	"github.com/Prokopevs/ccc/schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GRPC) AddUser(ctx context.Context, req *schema.AddUserRequest) (*schema.AddUserResponse, error) {
	err := g.service.AddUser(ctx, convertPBAddUserToCore(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.AddUserResponse{}, nil
}

func (g *GRPC) GetUser(ctx context.Context, req *schema.GetUserRequest) (*schema.GetUserResponse, error) {
	user, ok, err := g.service.GetUser(ctx, int(req.GetId()))
	if err != nil {
		if ok {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.GetUserResponse{
		User: convertCoreUserToPB(user),
	}, nil
}

func (g *GRPC) IsUserWithIdExists(ctx context.Context, req *schema.IsUserWithIdExistsRequest) (*schema.IsUserWithIdExistsResponse, error) {
	exists, err := g.service.IsUserWithIdExists(ctx, int(req.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.IsUserWithIdExistsResponse{
		Exists: exists,
	}, nil
}

func (g *GRPC) GetUserReferrals(ctx context.Context, req *schema.GetUserReferralsRequest) (*schema.GetUserReferralsResponse, error) {
	referrals, ok, err := g.service.GetUserReferrals(ctx, int(req.GetId()))
	if err != nil {
		if ok {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.GetUserReferralsResponse{
		Referrals: convertCoreReferralsToPB(referrals),
	}, nil
}