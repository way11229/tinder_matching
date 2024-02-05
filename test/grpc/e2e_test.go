package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/way11229/tinder_matching/pb"
)

func TestUser_Create_DeleteById(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	createResp, err := client.CreateUserAndListMatches(ctx, &pb.CreateUserAndListMatchesRequest{
		Name:                "test",
		Height:              150,
		Gender:              pb.UserGender_USER_GENDER_MALE,
		NumberOfWantedDates: 50,
	})
	if err != nil {
		t.Errorf("CreateUserAndListMatches API error = %v", err)
		return
	}

	if len(createResp.Matches) != 0 {
		t.Errorf("matches wanted empty, but get: %v", createResp.Matches)
		return
	}

	if _, err := client.DeleteUserById(ctx, &pb.DeleteUserByIdRequest{
		UserId: createResp.UserId,
	}); err != nil {
		t.Errorf("DeleteUserById API error: %v", err)
		return
	}
}

func TestUser_CreateBatch_ListMatchesByUserId_DeleteById(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	males := []*pb.User{
		{
			Name:                "test1",
			Height:              180,
			Gender:              pb.UserGender_USER_GENDER_MALE,
			RemainNumberOfDates: 50,
		},
		{
			Name:                "test2",
			Height:              170,
			Gender:              pb.UserGender_USER_GENDER_MALE,
			RemainNumberOfDates: 50,
		},
	}

	females := []*pb.User{
		{
			Name:                "test3",
			Height:              160,
			Gender:              pb.UserGender_USER_GENDER_FEMALE,
			RemainNumberOfDates: 50,
		},
	}

	for _, e := range males {
		createResp, err := client.CreateUserAndListMatches(ctx, &pb.CreateUserAndListMatchesRequest{
			Name:                e.Name,
			Height:              e.Height,
			Gender:              e.Gender,
			NumberOfWantedDates: e.RemainNumberOfDates,
		})
		if err != nil {
			t.Errorf("CreateUserAndListMatches API error = %v", err)
			return
		}

		if len(createResp.Matches) != 0 {
			t.Errorf("matches wanted empty, but get: %v", createResp.Matches)
			return
		}

		e.Id = createResp.UserId
	}

	for _, e := range females {
		createResp, err := client.CreateUserAndListMatches(ctx, &pb.CreateUserAndListMatchesRequest{
			Name:                e.Name,
			Height:              e.Height,
			Gender:              e.Gender,
			NumberOfWantedDates: e.RemainNumberOfDates,
		})
		if err != nil {
			t.Errorf("CreateUserAndListMatches API error = %v", err)
			return
		}

		for _, male := range males {
			male.RemainNumberOfDates -= 1
		}

		if !assert.ElementsMatch(t, createResp.Matches, males) {
			t.Errorf("CreateUserAndListMatches want get matches = %v, but get = %v", males, createResp.Matches)
			return
		}

		e.Id = createResp.UserId
		e.RemainNumberOfDates -= uint32(len(males))
	}

	for _, e := range males {
		resp, err := client.ListMatchesByUserId(ctx, &pb.ListMatchesByUserIdRequest{
			UserId: e.Id,
		})
		if err != nil {
			t.Errorf("ListMatchesByUserId API error = %v", err)
			return
		}

		if !assert.ElementsMatch(t, resp.Matches, females) {
			t.Errorf("ListMatchesByUserId want get matches = %v, but get = %v", females, resp.Matches)
			return
		}
	}

	deleteUsers := append(males, females...)
	for _, e := range deleteUsers {
		if _, err := client.DeleteUserById(ctx, &pb.DeleteUserByIdRequest{
			UserId: e.Id,
		}); err != nil {
			t.Errorf("DeleteUserById API error: %v", err)
			return
		}
	}
}
