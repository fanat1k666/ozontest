package in_memory

import (
	"context"
	"ozon/internal/adapters"
	"testing"
)

func TestIm_GetPost(t *testing.T) {

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
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				page: 0,
				size: 2,
			},
			expected: []adapters.PostResult{
				{AllowComments: false,
					AuthorId:  1,
					CreatedAt: "2024-06-03 21:16:17.857053 +0300 MSK m=+0.048998201",
					Id:        1,
					UpdatedAt: ""},
				{AllowComments: true,
					AuthorId:  2,
					CreatedAt: "2024-06-03 21:16:18.857053 +0300 MSK m=+1.048998201",
					Id:        2,
					UpdatedAt: ""},
			},
			expectedErr: false,
		},
		{
			name: "bad arguments",
			args: args{
				ctx:  context.Background(),
				page: -1,
				size: 5,
			},
			expected:    []adapters.PostResult{},
			expectedErr: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			im, _ := New()

			got := im.GetPost(tt.args.page, tt.args.size)

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

	type args struct {
		ctx    context.Context
		postId int
		page   int
		size   int
	}

	cases := []struct {
		name        string
		args        args
		expected    []adapters.PostWithCommentsResult
		expectedErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				postId: 2,
				page:   0,
				size:   2,
			},
			expected: []adapters.PostWithCommentsResult{
				{
					CommentAuthorId:   1,
					CommentCreatedAt:  "2024-06-03 21:24:47.3085974 +0300 MSK m=+0.119656201",
					CommentId:         1,
					CommentMsg:        "1",
					CommentParId:      0,
					CommentUpdatedAt:  "0001-01-01 00:00:00 +0000 UTC",
					PostAllowComments: true,
					PostAuthorId:      2,
					PostCreatedAt:     "2024-06-03 21:24:48.3085974 +0300 MSK m=+1.119656201",
					PostId:            2,
					PostUpdatedAt:     ""},
				{
					CommentAuthorId:   2,
					CommentCreatedAt:  "2024-06-03 21:44:35.6966789 +0300 MSK m=+1.072469701",
					CommentId:         2,
					CommentMsg:        "2",
					CommentParId:      1,
					CommentUpdatedAt:  "0001-01-01 00:00:00 +0000 UTC",
					PostAllowComments: true,
					PostAuthorId:      2,
					PostCreatedAt:     "2024-06-03 21:44:35.6966789 +0300 MSK m=+1.072469701",
					PostId:            2,
					PostUpdatedAt:     ""},
			},
			expectedErr: false,
		},
		{
			name: "bad arguments",
			args: args{
				ctx:    context.Background(),
				postId: 2,
				page:   -1,
				size:   5,
			},
			expected:    []adapters.PostWithCommentsResult{},
			expectedErr: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			im, _ := New()

			got := im.GetPostWithComments(tt.args.postId, tt.args.page, tt.args.size)

			for i := range got {
				if got[i].CommentAuthorId != tt.expected[i].CommentAuthorId ||
					got[i].CommentId != tt.expected[i].CommentId ||
					got[i].CommentMsg != tt.expected[i].CommentMsg ||
					got[i].CommentParId != tt.expected[i].CommentParId ||
					got[i].PostAllowComments != tt.expected[i].PostAllowComments ||
					got[i].PostAuthorId != tt.expected[i].PostAuthorId ||
					got[i].PostId != tt.expected[i].PostId {
					t.Errorf("got = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}
