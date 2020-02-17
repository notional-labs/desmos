package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

// ----------------------
// --- MsgAnswerPoll
// ----------------------

var msgAnswerPollPost = types.NewMsgAnswerPoll(types.PostID(1), []types.AnswerID{1, 2}, testOwner)

func TestMsgAnswerPollPost_Route(t *testing.T) {
	actual := msgAnswerPollPost.Route()
	require.Equal(t, "posts", actual)
}

func TestMsgAnswerPollPost_Type(t *testing.T) {
	actual := msgAnswerPollPost.Type()
	require.Equal(t, "answer_poll", actual)
}

func TestMsgAnswerPollPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgAnswerPoll
		error error
	}{
		{
			name:  "Invalid post id",
			msg:   types.NewMsgAnswerPoll(types.PostID(0), []types.AnswerID{1, 2}, msgAnswerPollPost.Answerer),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id"),
		},
		{
			name:  "Invalid answerer address",
			msg:   types.NewMsgAnswerPoll(types.PostID(1), []types.AnswerID{1, 2}, nil),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid answerer address: "),
		},
		{
			name:  "Returns error when no answer is provided",
			msg:   types.NewMsgAnswerPoll(types.PostID(1), []types.AnswerID{}, msgAnswerPollPost.Answerer),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Provided answers must contains at least one answer"),
		},
		{
			name: "Valid message returns no error",
			msg:  types.NewMsgAnswerPoll(types.PostID(1), []types.AnswerID{1, 2}, msgAnswerPollPost.Answerer),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgAnswerPollPost_GetSignBytes(t *testing.T) {
	actual := msgAnswerPollPost.GetSignBytes()
	expected := `{"type":"desmos/MsgAnswerPoll","value":{"answerer":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","answers":["1","2"],"post_id":"1"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgAnswerPollPost_GetSigners(t *testing.T) {
	actual := msgAnswerPollPost.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgAnswerPollPost.Answerer, actual[0])
}
