package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v2/x/staging/posts/keeper"
	"github.com/desmos-labs/desmos/v2/x/staging/posts/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	tests := []struct {
		name                string
		posts               []types.Post
		answers             []types.UserAnswer
		postReactions       []types.PostReaction
		registeredReactions []types.RegisteredReaction
		expStop             bool
	}{
		{
			name: "All invariants are not violated",
			posts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post without medias",
					Created:              suite.testData.post.Created,
					LastEdited:           time.Time{},
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					Attachments:          suite.testData.post.Attachments,
					Poll:                 suite.testData.post.Poll,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post without medias",
					CommentsState:        types.CommentsStateAllowed,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Created:              suite.testData.post.Created.Add(time.Hour),
					Creator:              "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			answers: []types.UserAnswer{
				types.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", suite.testData.post.Creator, []string{"1", "2"}),
			},
			postReactions: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"+1",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				),
			},
			registeredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					suite.testData.post.Creator,
					":like:",
					"+1",
					suite.testData.post.Subspace),
			},
			expStop: true,
		},
		{
			name: "ValidPosts Invariants violated",
			posts: []types.Post{
				{
					PostID:               "1234",
					Message:              "Message",
					Created:              suite.testData.post.Created,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				}},
			answers:             nil,
			postReactions:       nil,
			registeredReactions: nil,
			expStop:             true,
		},
		{
			name: "ValidCommentsDate Invariants violated",
			posts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post without medias",
					Created:              suite.testData.post.Created,
					LastEdited:           time.Time{},
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					Attachments:          suite.testData.post.Attachments,
					Poll:                 suite.testData.post.Poll,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Message",
					Created:              suite.testData.postEndPollDateExpired,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			answers:             nil,
			postReactions:       nil,
			registeredReactions: nil,
			expStop:             true,
		},
		{
			name:    "ValidPostForReactions Invariants violated",
			posts:   []types.Post{},
			answers: nil,
			postReactions: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"+1",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				),
			},
			registeredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					suite.testData.post.Creator,
					":like:",
					"+1",
					suite.testData.post.Subspace,
				),
			},
			expStop: true,
		},
		{
			name: "ValidPollForUserAnswers Invariants violated",
			posts: []types.Post{
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post without medias",
					CommentsState:        types.CommentsStateAllowed,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Created:              suite.testData.post.Created.Add(time.Hour),
					Creator:              "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			answers: []types.UserAnswer{
				types.NewUserAnswer("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd", suite.testData.post.Creator, []string{"1", "2"}),
			},
			postReactions:       nil,
			registeredReactions: nil,
			expStop:             true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetParams(suite.ctx, types.DefaultParams())

			for _, post := range test.posts {
				suite.k.SavePost(suite.ctx, post)
			}

			for _, reaction := range test.registeredReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			for _, reaction := range test.postReactions {
				err := suite.k.SavePostReaction(suite.ctx, reaction)
				suite.Require().NoError(err)
			}

			for _, answer := range test.answers {
				suite.k.SaveUserAnswer(suite.ctx, answer)
			}

			_, stop := keeper.AllInvariants(suite.k)(suite.ctx)
			suite.Require().Equal(test.expStop, stop)
		})
	}
}
