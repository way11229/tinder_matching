package handler

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/way11229/tinder_matching/domain"
	pb "github.com/way11229/tinder_matching/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcHandler struct {
	pb.UnimplementedUserServiceServer

	userService domain.UserService
}

func NewGrpcHandler(
	userService domain.UserService,
) *GrpcHandler {
	return &GrpcHandler{
		userService: userService,
	}
}

func (g *GrpcHandler) CreateUserAndListMatches(ctx context.Context, input *pb.CreateUserAndListMatchesRequest) (*pb.CreateUserAndListMatchesResponse, error) {
	if input.GetName() == "" || input.GetHeight() == 0 || input.GetNumberOfWantedDates() == 0 {
		err := domain.ErrorMissRequiredParameters
		return nil, status.Error(g.getGrpcCodeFromError(err), err.Error())
	}

	createUser := &domain.CreateUser{
		Name:                input.GetName(),
		Height:              input.GetHeight(),
		RemainNumberOfDates: input.GetNumberOfWantedDates(),
	}

	switch input.GetGender() {
	case pb.UserGender_USER_GENDER_MALE:
		createUser.Gender = domain.USER_GENDER_MALE
	case pb.UserGender_USER_GENDER_FEMALE:
		createUser.Gender = domain.USER_GENDER_FEMALE
	default:
		err := domain.ErrorMissRequiredParameters
		return nil, status.Error(g.getGrpcCodeFromError(err), err.Error())
	}

	resp, err := g.userService.CreateUserAndListMatches(ctx, createUser)
	if err != nil {
		return nil, status.Error(g.getGrpcCodeFromError(err), err.Error())
	}

	return &pb.CreateUserAndListMatchesResponse{
		UserId:  resp.UserId.String(),
		Matches: g.remodelUsers(resp.Matches),
	}, nil
}

func (g *GrpcHandler) DeleteUserById(ctx context.Context, input *pb.DeleteUserByIdRequest) (*emptypb.Empty, error) {
	userId, err := uuid.Parse(input.GetUserId())
	if err != nil {
		log.Printf("handler grpc, deleteUserById, id parse error: %v, input: %v", err, input.GetUserId())
		respErr := domain.ErrorUserIdInvalid
		return nil, status.Error(g.getGrpcCodeFromError(respErr), respErr.Error())
	}

	if err := g.userService.DeleteUserById(ctx, userId); err != nil {
		return nil, status.Error(g.getGrpcCodeFromError(err), err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g *GrpcHandler) ListMatchesByUserId(ctx context.Context, input *pb.ListMatchesByUserIdRequest) (*pb.ListMatchesByUserIdResponse, error) {
	userId, err := uuid.Parse(input.GetUserId())
	if err != nil {
		log.Printf("handler grpc, listMatchesById, id parse error: %v, input: %v", err, input.GetUserId())
		respErr := domain.ErrorUserIdInvalid
		return nil, status.Error(g.getGrpcCodeFromError(respErr), respErr.Error())
	}

	search := &domain.UserListMatchesByUserId{
		UserId: userId,
		Limit:  int(input.GetLimit()),
	}

	resp, err := g.userService.ListMatchesByUserId(ctx, search)
	if err != nil {
		return nil, status.Error(g.getGrpcCodeFromError(err), err.Error())
	}

	return &pb.ListMatchesByUserIdResponse{
		Matches: g.remodelUsers(resp),
	}, nil
}

func (g *GrpcHandler) getGrpcCodeFromError(err error) codes.Code {
	switch err {
	case
		domain.ErrorInternalServerError,
		domain.ErrorRecordNotFound:
		return codes.Internal
	case
		domain.ErrorMissRequiredParameters,
		domain.ErrorUserIdInvalid,
		domain.ErrorUserNameInvalid,
		domain.ErrorUserHeightInvalid,
		domain.ErrorUserGenderInvalid,
		domain.ErrorUserNumberOfWantedDatesInvalid:
		return codes.InvalidArgument
	default:
		return codes.Unknown
	}
}

func (g *GrpcHandler) remodelUsers(data []*domain.User) []*pb.User {
	rtn := []*pb.User{}
	for _, e := range data {
		user := &pb.User{
			Id:                  e.Id.String(),
			Name:                e.Name,
			Height:              e.Height,
			Gender:              pb.UserGender_USER_GENDER_MALE,
			RemainNumberOfDates: e.RemainNumberOfDates,
		}

		if e.Gender == domain.USER_GENDER_FEMALE {
			user.Gender = pb.UserGender_USER_GENDER_FEMALE
		}

		rtn = append(rtn, user)
	}

	return rtn
}
