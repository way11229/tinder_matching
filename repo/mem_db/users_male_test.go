package memdb

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/way11229/tinder_matching/domain"
)

func TestUsersMaleMemDB_Create_GetById_DeleteById(t *testing.T) {
	type fields struct {
		memdb *MemDB
	}
	type args struct {
		ctx   context.Context
		input *domain.UsersMemDbCreate
	}

	db := NewMemDB()
	ctx := context.Background()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "create and get by id and delete by id success",
			fields: fields{
				memdb: db,
			},
			args: args{
				ctx: ctx,
				input: &domain.UsersMemDbCreate{
					Name:                "test1",
					Height:              180,
					RemainNumberOfDates: 50,
				},
			},
			want: &domain.User{
				Name:                "test1",
				Height:              180,
				Gender:              domain.USER_GENDER_MALE,
				RemainNumberOfDates: 50,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersMaleMemDB{
				memdb: tt.fields.memdb,
			}

			resp, err := u.Create(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.want.Id = resp.UUID

			got, err := u.GetById(tt.args.ctx, resp.UUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersMaleMemDB.GetById() = %v, want %v", got, tt.want)
				return
			}

			if err := u.DeleteById(tt.args.ctx, resp.UUID); (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.DeleteById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUsersMaleMemDB_Create_UpdateBatch_GetById_DeleteById(t *testing.T) {
	type fields struct {
		memdb *MemDB
	}
	type args struct {
		ctx         context.Context
		createInput *domain.UsersMemDbCreate
		updateInput *domain.UsersMemDbUpdate
	}

	db := NewMemDB()
	ctx := context.Background()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "create and update and get by id and delete by id success",
			fields: fields{
				memdb: db,
			},
			args: args{
				ctx: ctx,
				createInput: &domain.UsersMemDbCreate{
					Name:                "test1",
					Height:              180,
					RemainNumberOfDates: 50,
				},
				updateInput: &domain.UsersMemDbUpdate{
					Name:                "test2",
					Height:              190,
					RemainNumberOfDates: 25,
				},
			},
			want: &domain.User{
				Name:                "test2",
				Height:              190,
				Gender:              domain.USER_GENDER_MALE,
				RemainNumberOfDates: 25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersMaleMemDB{
				memdb: tt.fields.memdb,
			}

			resp, err := u.Create(tt.args.ctx, tt.args.createInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.want.Id = resp.UUID
			tt.args.updateInput.Id = resp.UUID

			if err := u.UpdateBatch(tt.args.ctx, []*domain.UsersMemDbUpdate{tt.args.updateInput}); (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.UpdateBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := u.GetById(tt.args.ctx, resp.UUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersMaleMemDB.GetById() = %v, want %v", got, tt.want)
				return
			}

			if err := u.DeleteById(tt.args.ctx, resp.UUID); (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.DeleteById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUsersMaleMemDB_Create_DeleteByIds(t *testing.T) {
	type fields struct {
		memdb *MemDB
	}
	type args struct {
		ctx   context.Context
		input *domain.UsersMemDbCreate
	}

	db := NewMemDB()
	ctx := context.Background()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create and delete by ids success",
			fields: fields{
				memdb: db,
			},
			args: args{
				ctx: ctx,
				input: &domain.UsersMemDbCreate{
					Name:                "test1",
					Height:              180,
					RemainNumberOfDates: 50,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersMaleMemDB{
				memdb: tt.fields.memdb,
			}

			resp, err := u.Create(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := u.DeleteByIds(tt.args.ctx, []uuid.UUID{resp.UUID}); (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.DeleteByIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUsersMaleMemDB_Create_ListByHeightUpperBoundWithoutEqual_DeleteByIds(t *testing.T) {
	type fields struct {
		memdb *MemDB
	}
	type args struct {
		ctx         context.Context
		createBatch []*domain.UsersMemDbCreate
		search      *domain.UsersMemDbHeightSearch
	}

	db := NewMemDB()
	ctx := context.Background()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.User
		wantErr bool
	}{
		{
			name: "create and list by height upper bound without equal and delete by ids success",
			fields: fields{
				memdb: db,
			},
			args: args{
				ctx: ctx,
				createBatch: []*domain.UsersMemDbCreate{
					{
						Name:                "test1",
						Height:              170,
						RemainNumberOfDates: 50,
					},
					{
						Name:                "test2",
						Height:              180,
						RemainNumberOfDates: 50,
					},
					{
						Name:                "test3",
						Height:              190,
						RemainNumberOfDates: 50,
					},
				},
				search: &domain.UsersMemDbHeightSearch{
					Limit: nil,
					Bound: 190,
				},
			},
			want: []*domain.User{
				{
					Name:                "test2",
					Height:              180,
					Gender:              domain.USER_GENDER_MALE,
					RemainNumberOfDates: 50,
				},
				{
					Name:                "test1",
					Height:              170,
					Gender:              domain.USER_GENDER_MALE,
					RemainNumberOfDates: 50,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersMaleMemDB{
				memdb: tt.fields.memdb,
			}

			deleteIds := []uuid.UUID{}
			for _, e := range tt.args.createBatch {
				resp, err := u.Create(tt.args.ctx, e)
				if (err != nil) != tt.wantErr {
					t.Errorf("UsersMaleMemDB.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				deleteIds = append(deleteIds, resp.UUID)
				for _, j := range tt.want {
					if e.Name == j.Name {
						j.Id = resp.UUID
					}
				}
			}

			resp, err := u.ListByHeightUpperBoundWithoutEqual(tt.args.ctx, tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.ListByHeightUpperBoundWithoutEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for idx, val := range resp {
				if !reflect.DeepEqual(val, tt.want[idx]) {
					t.Errorf("UsersMaleMemDB.ListByHeightUpperBoundWithoutEqual() = %v, want %v", val, tt.want[idx])
					return
				}
			}

			if err := u.DeleteByIds(tt.args.ctx, deleteIds); (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.DeleteByIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUsersMaleMemDB_Create_ListByHeightLowerBoundWithoutEqual_DeleteByIds(t *testing.T) {
	type fields struct {
		memdb *MemDB
	}
	type args struct {
		ctx         context.Context
		createBatch []*domain.UsersMemDbCreate
		search      *domain.UsersMemDbHeightSearch
	}

	db := NewMemDB()
	ctx := context.Background()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.User
		wantErr bool
	}{
		{
			name: "create and list by height lower bound without equal and delete by ids success",
			fields: fields{
				memdb: db,
			},
			args: args{
				ctx: ctx,
				createBatch: []*domain.UsersMemDbCreate{
					{
						Name:                "test1",
						Height:              170,
						RemainNumberOfDates: 50,
					},
					{
						Name:                "test2",
						Height:              180,
						RemainNumberOfDates: 50,
					},
					{
						Name:                "test3",
						Height:              190,
						RemainNumberOfDates: 50,
					},
				},
				search: &domain.UsersMemDbHeightSearch{
					Limit: nil,
					Bound: 170,
				},
			},
			want: []*domain.User{
				{
					Name:                "test2",
					Height:              180,
					Gender:              domain.USER_GENDER_MALE,
					RemainNumberOfDates: 50,
				},
				{
					Name:                "test3",
					Height:              190,
					Gender:              domain.USER_GENDER_MALE,
					RemainNumberOfDates: 50,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersMaleMemDB{
				memdb: tt.fields.memdb,
			}

			deleteIds := []uuid.UUID{}
			for _, e := range tt.args.createBatch {
				resp, err := u.Create(tt.args.ctx, e)
				if (err != nil) != tt.wantErr {
					t.Errorf("UsersMaleMemDB.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				deleteIds = append(deleteIds, resp.UUID)
				for _, j := range tt.want {
					if e.Name == j.Name {
						j.Id = resp.UUID
					}
				}
			}

			resp, err := u.ListByHeightLowerBoundWithoutEqual(tt.args.ctx, tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.ListByHeightLowerBoundWithoutEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for idx, val := range resp {
				if !reflect.DeepEqual(val, tt.want[idx]) {
					t.Errorf("UsersMaleMemDB.ListByHeightLowerBoundWithoutEqual() = %v, want %v", val, tt.want[idx])
					return
				}
			}

			if err := u.DeleteByIds(tt.args.ctx, deleteIds); (err != nil) != tt.wantErr {
				t.Errorf("UsersMaleMemDB.DeleteByIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
