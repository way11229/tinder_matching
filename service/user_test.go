package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/way11229/tinder_matching/domain"
	"github.com/way11229/tinder_matching/mocks"
)

func TestUserService_validateUserName(t *testing.T) {
	type args struct {
		name string
	}

	invalidName := ""
	for i := 0; i < 200; i += 1 {
		invalidName += "w"
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				name: "test1",
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				name: invalidName,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{}
			if got := u.validateUserName(tt.args.name); got != tt.want {
				t.Errorf("UserService.validateUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_validateUserHeight(t *testing.T) {
	type args struct {
		height uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				height: 150,
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				height: 350,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{}
			if got := u.validateUserHeight(tt.args.height); got != tt.want {
				t.Errorf("UserService.validateUserHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_validateUserNumberOfWantedDates(t *testing.T) {
	type args struct {
		number uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				number: 50,
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				number: 200,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{}
			if got := u.validateUserNumberOfWantedDates(tt.args.number); got != tt.want {
				t.Errorf("UserService.validateUserNumberOfWantedDates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_validateCreateUserData(t *testing.T) {
	type args struct {
		input *domain.CreateUser
	}

	invalidName := ""
	for i := 0; i < 200; i += 1 {
		invalidName += "w"
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "all valid",
			args: args{
				input: &domain.CreateUser{
					Name:                "test1",
					Height:              150,
					Gender:              domain.USER_GENDER_MALE,
					RemainNumberOfDates: 50,
				},
			},
			wantErr: false,
		},
		{
			name: "name invalid",
			args: args{
				input: &domain.CreateUser{
					Name:                invalidName,
					Height:              150,
					Gender:              domain.USER_GENDER_MALE,
					RemainNumberOfDates: 50,
				},
			},
			wantErr: true,
		},
		{
			name: "height invalid",
			args: args{
				input: &domain.CreateUser{
					Name:                "test1",
					Height:              300,
					Gender:              domain.USER_GENDER_MALE,
					RemainNumberOfDates: 50,
				},
			},
			wantErr: true,
		},
		{
			name: "remain number of dates invalid",
			args: args{
				input: &domain.CreateUser{
					Name:                "test1",
					Height:              180,
					Gender:              domain.USER_GENDER_MALE,
					RemainNumberOfDates: 200,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{}
			if err := u.validateCreateUserData(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("UserService.validateCreateUserData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_getUserById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}

	ctx := context.Background()
	genId := uuid.New()

	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "male",
			args: args{
				ctx: ctx,
				id:  genId,
			},
			want: &domain.User{
				Id:                  genId,
				Name:                "test",
				Height:              150,
				Gender:              domain.USER_GENDER_MALE,
				RemainNumberOfDates: 50,
			},
			wantErr: false,
		},
		{
			name: "female",
			args: args{
				ctx: ctx,
				id:  genId,
			},
			want: &domain.User{
				Id:                  genId,
				Name:                "test",
				Height:              150,
				Gender:              domain.USER_GENDER_FEMALE,
				RemainNumberOfDates: 50,
			},
			wantErr: false,
		},
		{
			name: "record not found",
			args: args{
				ctx: ctx,
				id:  genId,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMaleMemDB := mocks.NewUsersMaleMemDB(t)
			mockFemaleMemDB := mocks.NewUsersFemaleMemDB(t)

			switch tt.name {
			case "male":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(
					&domain.User{
						Id:                  tt.want.Id,
						Name:                tt.want.Name,
						Height:              tt.want.Height,
						Gender:              tt.want.Gender,
						RemainNumberOfDates: tt.want.RemainNumberOfDates,
					}, nil)
			case "female":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(nil, domain.ErrorRecordNotFound)
				mockFemaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(
					&domain.User{
						Id:                  tt.want.Id,
						Name:                tt.want.Name,
						Height:              tt.want.Height,
						Gender:              tt.want.Gender,
						RemainNumberOfDates: tt.want.RemainNumberOfDates,
					}, nil)
			case "record not found":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(nil, domain.ErrorRecordNotFound)
				mockFemaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(nil, domain.ErrorRecordNotFound)
			}

			u := &UserService{
				usersMaleMemDB:   mockMaleMemDB,
				usersFemaleMemDB: mockFemaleMemDB,
			}
			got, err := u.getUserById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.getUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.getUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZero(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType
	}

	ctx := context.Background()
	updateUserId := uuid.New()
	deleteUserId := uuid.New()

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "male_update",
			args: args{
				ctx: ctx,
				input: &reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType{
					Users: []*domain.User{
						{
							Id:                  updateUserId,
							Name:                "update user",
							Height:              100,
							Gender:              domain.USER_GENDER_MALE,
							RemainNumberOfDates: 11,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "male_delete",
			args: args{
				ctx: ctx,
				input: &reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType{
					Users: []*domain.User{
						{
							Id:                  deleteUserId,
							Name:                "delete user",
							Height:              100,
							Gender:              domain.USER_GENDER_MALE,
							RemainNumberOfDates: 1,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "female_update",
			args: args{
				ctx: ctx,
				input: &reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType{
					Users: []*domain.User{
						{
							Id:                  updateUserId,
							Name:                "update user",
							Height:              100,
							Gender:              domain.USER_GENDER_FEMALE,
							RemainNumberOfDates: 11,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "female_delete",
			args: args{
				ctx: ctx,
				input: &reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZeroType{
					Users: []*domain.User{
						{
							Id:                  deleteUserId,
							Name:                "delete user",
							Height:              100,
							Gender:              domain.USER_GENDER_FEMALE,
							RemainNumberOfDates: 1,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMaleMemDB := mocks.NewUsersMaleMemDB(t)
			mockFemaleMemDB := mocks.NewUsersFemaleMemDB(t)

			switch tt.name {
			case "male_update":
				mockMaleMemDB.EXPECT().UpdateBatch(
					tt.args.ctx,
					[]*domain.UsersMemDbUpdate{{
						Id:                  tt.args.input.Users[0].Id,
						Name:                tt.args.input.Users[0].Name,
						Height:              tt.args.input.Users[0].Height,
						RemainNumberOfDates: tt.args.input.Users[0].RemainNumberOfDates - 1,
					}},
				).Return(nil)

				tt.args.input.UsersMemDb = mockMaleMemDB

			case "male_delete":
				mockMaleMemDB.EXPECT().DeleteByIds(
					tt.args.ctx,
					[]uuid.UUID{tt.args.input.Users[0].Id},
				).Return(nil)

				tt.args.input.UsersMemDb = mockMaleMemDB

			case "female_update":
				mockFemaleMemDB.EXPECT().UpdateBatch(
					tt.args.ctx,
					[]*domain.UsersMemDbUpdate{{
						Id:                  tt.args.input.Users[0].Id,
						Name:                tt.args.input.Users[0].Name,
						Height:              tt.args.input.Users[0].Height,
						RemainNumberOfDates: tt.args.input.Users[0].RemainNumberOfDates - 1,
					}},
				).Return(nil)

				tt.args.input.UsersMemDb = mockFemaleMemDB

			case "female_delete":
				mockFemaleMemDB.EXPECT().DeleteByIds(
					tt.args.ctx,
					[]uuid.UUID{tt.args.input.Users[0].Id},
				).Return(nil)

				tt.args.input.UsersMemDb = mockFemaleMemDB
			}

			u := &UserService{}
			if err := u.reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZero(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("UserService.reduceUserNumberOfDatesAndDeleteUserWhenBecomeToZero() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_listMatchesByUserId(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *listMatchesByUserIdType
	}

	ctx := context.Background()
	male := &domain.User{
		Id:                  uuid.New(),
		Name:                "test",
		Height:              90,
		Gender:              domain.USER_GENDER_MALE,
		RemainNumberOfDates: 100,
	}
	female := &domain.User{
		Id:                  uuid.New(),
		Name:                "test1",
		Height:              90,
		Gender:              domain.USER_GENDER_FEMALE,
		RemainNumberOfDates: 100,
	}

	maleReduceResult := *male
	maleReduceResult.RemainNumberOfDates -= 1

	femaleReduceResult := *female
	femaleReduceResult.RemainNumberOfDates -= 1

	tests := []struct {
		name    string
		args    args
		want    []*domain.User
		wantErr bool
	}{
		{
			name: "user_male_reduce",
			args: args{
				ctx: ctx,
				input: &listMatchesByUserIdType{
					UserId:              male.Id,
					Limit:               nil,
					ReduceNumberOfDates: true,
				},
			},
			want: []*domain.User{
				&femaleReduceResult,
			},
		},
		{
			name: "user_male_not_reduce",
			args: args{
				ctx: ctx,
				input: &listMatchesByUserIdType{
					UserId:              male.Id,
					Limit:               nil,
					ReduceNumberOfDates: false,
				},
			},
			want: []*domain.User{
				female,
			},
		},
		{
			name: "user_female_reduce",
			args: args{
				ctx: ctx,
				input: &listMatchesByUserIdType{
					UserId:              female.Id,
					Limit:               nil,
					ReduceNumberOfDates: true,
				},
			},
			want: []*domain.User{
				&maleReduceResult,
			},
		},
		{
			name: "user_female_not_reduce",
			args: args{
				ctx: ctx,
				input: &listMatchesByUserIdType{
					UserId:              female.Id,
					Limit:               nil,
					ReduceNumberOfDates: false,
				},
			},
			want: []*domain.User{
				male,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMaleMemDB := mocks.NewUsersMaleMemDB(t)
			mockFemaleMemDB := mocks.NewUsersFemaleMemDB(t)

			switch tt.name {
			case "user_male_reduce":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(male, nil)
				mockFemaleMemDB.EXPECT().ListByHeightUpperBoundWithoutEqual(
					tt.args.ctx,
					&domain.UsersMemDbHeightSearch{
						Limit: tt.args.input.Limit,
						Bound: male.Height,
					},
				).Return(
					[]*domain.User{female},
					nil,
				)
				mockFemaleMemDB.EXPECT().UpdateBatch(
					tt.args.ctx,
					[]*domain.UsersMemDbUpdate{{
						Id:                  female.Id,
						Name:                female.Name,
						Height:              female.Height,
						RemainNumberOfDates: female.RemainNumberOfDates - 1,
					}},
				).Return(nil)
			case "user_male_not_reduce":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(male, nil)
				mockFemaleMemDB.EXPECT().ListByHeightUpperBoundWithoutEqual(
					tt.args.ctx,
					&domain.UsersMemDbHeightSearch{
						Limit: tt.args.input.Limit,
						Bound: male.Height,
					},
				).Return(
					[]*domain.User{female},
					nil,
				)
			case "user_female_reduce":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(nil, domain.ErrorRecordNotFound)
				mockFemaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(female, nil)
				mockMaleMemDB.EXPECT().ListByHeightLowerBoundWithoutEqual(
					tt.args.ctx,
					&domain.UsersMemDbHeightSearch{
						Limit: tt.args.input.Limit,
						Bound: female.Height,
					},
				).Return(
					[]*domain.User{male},
					nil,
				)
				mockMaleMemDB.EXPECT().UpdateBatch(
					tt.args.ctx,
					[]*domain.UsersMemDbUpdate{{
						Id:                  male.Id,
						Name:                male.Name,
						Height:              male.Height,
						RemainNumberOfDates: male.RemainNumberOfDates - 1,
					}},
				).Return(nil)
			case "user_female_not_reduce":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(nil, domain.ErrorRecordNotFound)
				mockFemaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(female, nil)
				mockMaleMemDB.EXPECT().ListByHeightLowerBoundWithoutEqual(
					tt.args.ctx,
					&domain.UsersMemDbHeightSearch{
						Limit: tt.args.input.Limit,
						Bound: female.Height,
					},
				).Return(
					[]*domain.User{male},
					nil,
				)
			}

			u := &UserService{
				usersMaleMemDB:   mockMaleMemDB,
				usersFemaleMemDB: mockFemaleMemDB,
			}
			resp, err := u.listMatchesByUserId(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.listMatchesByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assert.ElementsMatch(t, resp, tt.want) {
				t.Errorf("UserService.listMatchesByUserId() = %v, want %v", resp, tt.want)
			}
		})
	}
}

func TestUserService_CreateUserAndListMatches(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *domain.CreateUser
	}

	ctx := context.Background()
	genId := uuid.New()
	male := &domain.User{
		Id:                  uuid.New(),
		Name:                "test",
		Height:              100,
		Gender:              domain.USER_GENDER_MALE,
		RemainNumberOfDates: 100,
	}
	female := &domain.User{
		Id:                  uuid.New(),
		Name:                "test1",
		Height:              90,
		Gender:              domain.USER_GENDER_FEMALE,
		RemainNumberOfDates: 100,
	}

	maleReduceResult := *male
	maleReduceResult.RemainNumberOfDates -= 1

	femaleReduceResult := *female
	femaleReduceResult.RemainNumberOfDates -= 1

	tests := []struct {
		name    string
		args    args
		want    *domain.CreateUserResp
		wantErr bool
	}{
		{
			name: "male",
			args: args{
				ctx: ctx,
				input: &domain.CreateUser{
					Name:                "newer",
					Height:              150,
					Gender:              domain.USER_GENDER_MALE,
					RemainNumberOfDates: 50,
				},
			},
			want: &domain.CreateUserResp{
				UserId: genId,
				Matches: []*domain.User{
					&femaleReduceResult,
				},
			},
		},
		{
			name: "female",
			args: args{
				ctx: ctx,
				input: &domain.CreateUser{
					Name:                "newer",
					Height:              150,
					Gender:              domain.USER_GENDER_FEMALE,
					RemainNumberOfDates: 50,
				},
			},
			want: &domain.CreateUserResp{
				UserId: genId,
				Matches: []*domain.User{
					&maleReduceResult,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMaleMemDB := mocks.NewUsersMaleMemDB(t)
			mockFemaleMemDB := mocks.NewUsersFemaleMemDB(t)

			switch tt.name {
			case "male":
				mockMaleMemDB.EXPECT().Create(
					tt.args.ctx,
					&domain.UsersMemDbCreate{
						Name:                tt.args.input.Name,
						Height:              tt.args.input.Height,
						RemainNumberOfDates: tt.args.input.RemainNumberOfDates,
					},
				).Return(
					uuid.NullUUID{
						Valid: true,
						UUID:  genId,
					}, nil,
				)
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					genId,
				).Return(&domain.User{
					Id:                  genId,
					Name:                tt.args.input.Name,
					Height:              tt.args.input.Height,
					RemainNumberOfDates: tt.args.input.RemainNumberOfDates,
					Gender:              domain.USER_GENDER_MALE,
				}, nil)
				mockFemaleMemDB.EXPECT().ListByHeightUpperBoundWithoutEqual(
					tt.args.ctx,
					&domain.UsersMemDbHeightSearch{
						Limit: nil,
						Bound: tt.args.input.Height,
					},
				).Return(
					[]*domain.User{female},
					nil,
				)
				mockFemaleMemDB.EXPECT().UpdateBatch(
					tt.args.ctx,
					[]*domain.UsersMemDbUpdate{{
						Id:                  female.Id,
						Name:                female.Name,
						Height:              female.Height,
						RemainNumberOfDates: female.RemainNumberOfDates - 1,
					}},
				).Return(nil)
			case "female":
				mockFemaleMemDB.EXPECT().Create(
					tt.args.ctx,
					&domain.UsersMemDbCreate{
						Name:                tt.args.input.Name,
						Height:              tt.args.input.Height,
						RemainNumberOfDates: tt.args.input.RemainNumberOfDates,
					},
				).Return(
					uuid.NullUUID{
						Valid: true,
						UUID:  genId,
					}, nil,
				)
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					genId,
				).Return(nil, domain.ErrorRecordNotFound)
				mockFemaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					genId,
				).Return(&domain.User{
					Id:                  genId,
					Name:                tt.args.input.Name,
					Height:              tt.args.input.Height,
					RemainNumberOfDates: tt.args.input.RemainNumberOfDates,
					Gender:              domain.USER_GENDER_FEMALE,
				}, nil)
				mockMaleMemDB.EXPECT().ListByHeightLowerBoundWithoutEqual(
					tt.args.ctx,
					&domain.UsersMemDbHeightSearch{
						Limit: nil,
						Bound: tt.args.input.Height,
					},
				).Return(
					[]*domain.User{male},
					nil,
				)
				mockMaleMemDB.EXPECT().UpdateBatch(
					tt.args.ctx,
					[]*domain.UsersMemDbUpdate{{
						Id:                  male.Id,
						Name:                male.Name,
						Height:              male.Height,
						RemainNumberOfDates: male.RemainNumberOfDates - 1,
					}},
				).Return(nil)
			}

			u := &UserService{
				usersMaleMemDB:   mockMaleMemDB,
				usersFemaleMemDB: mockFemaleMemDB,
			}
			got, err := u.CreateUserAndListMatches(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUserAndListMatches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.CreateUserAndListMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_DeleteUserById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}

	ctx := context.Background()
	male := &domain.User{
		Id:                  uuid.New(),
		Name:                "test",
		Height:              100,
		Gender:              domain.USER_GENDER_MALE,
		RemainNumberOfDates: 100,
	}
	female := &domain.User{
		Id:                  uuid.New(),
		Name:                "test1",
		Height:              90,
		Gender:              domain.USER_GENDER_FEMALE,
		RemainNumberOfDates: 100,
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "male",
			args: args{
				ctx: ctx,
				id:  male.Id,
			},
			wantErr: false,
		},
		{
			name: "female",
			args: args{
				ctx: ctx,
				id:  female.Id,
			},
			wantErr: false,
		},
		{
			name: "record not found",
			args: args{
				ctx: ctx,
				id:  uuid.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMaleMemDB := mocks.NewUsersMaleMemDB(t)
			mockFemaleMemDB := mocks.NewUsersFemaleMemDB(t)

			switch tt.name {
			case "male":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(
					&domain.User{
						Id:                  male.Id,
						Name:                male.Name,
						Height:              male.Height,
						Gender:              male.Gender,
						RemainNumberOfDates: male.RemainNumberOfDates,
					}, nil)
				mockMaleMemDB.EXPECT().DeleteById(
					tt.args.ctx,
					tt.args.id,
				).Return(nil)
			case "female":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(nil, domain.ErrorRecordNotFound)
				mockFemaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(
					&domain.User{
						Id:                  female.Id,
						Name:                female.Name,
						Height:              female.Height,
						Gender:              female.Gender,
						RemainNumberOfDates: female.RemainNumberOfDates,
					}, nil)
				mockFemaleMemDB.EXPECT().DeleteById(
					tt.args.ctx,
					tt.args.id,
				).Return(nil)
			case "record not found":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(nil, domain.ErrorRecordNotFound)
				mockFemaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.id,
				).Return(nil, domain.ErrorRecordNotFound)
			}

			u := &UserService{
				usersMaleMemDB:   mockMaleMemDB,
				usersFemaleMemDB: mockFemaleMemDB,
			}
			if err := u.DeleteUserById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserService.DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_ListMatchesByUserId(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *domain.UserListMatchesByUserId
	}

	ctx := context.Background()
	male := &domain.User{
		Id:                  uuid.New(),
		Name:                "test",
		Height:              90,
		Gender:              domain.USER_GENDER_MALE,
		RemainNumberOfDates: 100,
	}
	female := &domain.User{
		Id:                  uuid.New(),
		Name:                "test1",
		Height:              90,
		Gender:              domain.USER_GENDER_FEMALE,
		RemainNumberOfDates: 100,
	}

	tests := []struct {
		name    string
		args    args
		want    []*domain.User
		wantErr bool
	}{
		{
			name: "male",
			args: args{
				ctx: ctx,
				input: &domain.UserListMatchesByUserId{
					UserId: male.Id,
					Limit:  20,
				},
			},
			want: []*domain.User{
				female,
			},
		},
		{
			name: "female",
			args: args{
				ctx: ctx,
				input: &domain.UserListMatchesByUserId{
					UserId: female.Id,
					Limit:  20,
				},
			},
			want: []*domain.User{
				male,
			},
		},
		{
			name: "empty",
			args: args{
				ctx: ctx,
				input: &domain.UserListMatchesByUserId{
					UserId: uuid.New(),
					Limit:  20,
				},
			},
			want: []*domain.User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMaleMemDB := mocks.NewUsersMaleMemDB(t)
			mockFemaleMemDB := mocks.NewUsersFemaleMemDB(t)

			switch tt.name {
			case "male":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(male, nil)
				mockFemaleMemDB.EXPECT().ListByHeightUpperBoundWithoutEqual(
					tt.args.ctx,
					&domain.UsersMemDbHeightSearch{
						Limit: &tt.args.input.Limit,
						Bound: male.Height,
					},
				).Return(
					[]*domain.User{female},
					nil,
				)
			case "female":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(nil, domain.ErrorRecordNotFound)
				mockFemaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(female, nil)
				mockMaleMemDB.EXPECT().ListByHeightLowerBoundWithoutEqual(
					tt.args.ctx,
					&domain.UsersMemDbHeightSearch{
						Limit: &tt.args.input.Limit,
						Bound: female.Height,
					},
				).Return(
					[]*domain.User{male},
					nil,
				)
			case "empty":
				mockMaleMemDB.EXPECT().GetById(
					tt.args.ctx,
					tt.args.input.UserId,
				).Return(male, nil)
				mockFemaleMemDB.EXPECT().ListByHeightUpperBoundWithoutEqual(
					tt.args.ctx,
					&domain.UsersMemDbHeightSearch{
						Limit: &tt.args.input.Limit,
						Bound: male.Height,
					},
				).Return(
					[]*domain.User{},
					nil,
				)
			}

			u := &UserService{
				usersMaleMemDB:   mockMaleMemDB,
				usersFemaleMemDB: mockFemaleMemDB,
			}
			got, err := u.ListMatchesByUserId(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.ListMatchesByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.ListMatchesByUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}
