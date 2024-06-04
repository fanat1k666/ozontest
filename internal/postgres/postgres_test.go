package postgres

import (
	"context"
	"go.uber.org/mock/gomock"
	mock_adapters "ozon/internal/adapters/mocks"

	"ozon/internal/adapters"
	"testing"
)

func TestIm_GetPost(t *testing.T) {

	type fields struct {
		db *mock_adapters.MockDb
	}

	type args struct {
		ctx  context.Context
		page int
		size int
	}

	cases := []struct {
		name        string
		args        args
		expected    []adapters.PostResult
		expectedErr bool
		prepare     func(args, fields)
	}{
		{
			name: "error",
			args: args{
				ctx:  context.Background(),
				page: -1,
				size: -1,
			},
			prepare: func(args args, fields fields) {
				//fields.db.EXPECT().Query(`SELECT * FROM "post"
				//ORDER BY id
				//LIMIT $1 OFFSET $2`, -1, -1).Return(nil, fmt.Errorf("db error"))
			},
			expectedErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				db: mock_adapters.NewMockDb(ctrl),
			}
			tt.prepare(tt.args, f)
			n := New(f.db)
			got := n.GetPost(tt.args.page, tt.args.size)
			for i := range got {
				if got[i].AllowComments != tt.expected[i].AllowComments ||
					got[i].Id != tt.expected[i].Id ||
					got[i].AuthorId != tt.expected[i].AuthorId {
					t.Errorf("got = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}

func TestIm_GetPostWithComments(t *testing.T) {

	type fields struct {
		db *mock_adapters.MockDb
	}

	type args struct {
		ctx    context.Context
		postId int
		page   int
		size   int
	}

	cases := []struct {
		name        string
		args        args
		expected    []adapters.PostResult
		expectedErr bool
		prepare     func(args, fields)
	}{
		{
			name: "error",
			args: args{
				ctx:    context.Background(),
				postId: 2,
				page:   -1,
				size:   -1,
			},
			prepare: func(args args, fields fields) {
				//fields.db.EXPECT().Query(`SELECT * FROM "post"
				//ORDER BY id
				//LIMIT $1 OFFSET $2`, -1, -1).Return(nil, fmt.Errorf("db error"))
			},
			expectedErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				db: mock_adapters.NewMockDb(ctrl),
			}
			tt.prepare(tt.args, f)
			n := New(f.db)
			got := n.GetPost(tt.args.page, tt.args.size)
			for i := range got {
				if got[i].AllowComments != tt.expected[i].AllowComments ||
					got[i].Id != tt.expected[i].Id ||
					got[i].AuthorId != tt.expected[i].AuthorId {
					t.Errorf("got = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}
