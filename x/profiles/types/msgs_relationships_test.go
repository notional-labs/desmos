package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"

	"github.com/stretchr/testify/require"
)

var msgCreateRelationship = types.NewMsgCreateRelationship(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
)

func TestMsgCreateRelationship_Route(t *testing.T) {
	require.Equal(t, "profiles", msgCreateRelationship.Route())
}

func TestMsgCreateRelationship_Type(t *testing.T) {
	require.Equal(t, "create_relationship", msgCreateRelationship.Type())
}

func TestMsgCreateRelationship_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCreateRelationship
		shouldErr bool
	}{
		{
			name: "empty sender returns error",
			msg: types.NewMsgCreateRelationship(
				"",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "empty receiver returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "equal sender and receiver returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"1234",
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgCreateRelationship,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCreateRelationship_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateRelationship","value":{"receiver":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`
	require.Equal(t, expected, string(msgCreateRelationship.GetSignBytes()))
}

func TestMsgCreateRelationship_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCreateRelationship.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreateRelationship.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgDeleteRelationships = types.NewMsgDeleteRelationship(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
)

func TestMsgDeleteRelationships_Route(t *testing.T) {
	require.Equal(t, "profiles", msgDeleteRelationships.Route())
}

func TestMsgDeleteRelationships_Type(t *testing.T) {
	require.Equal(t, "delete_relationship", msgDeleteRelationships.Type())
}

func TestMsgDeleteRelationships_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgDeleteRelationship
		shouldErr bool
	}{
		{
			name: "empty sender returns error",
			msg: types.NewMsgDeleteRelationship(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "empty receiver returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "equal sender and receiver returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"1234",
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgDeleteRelationships,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgDeleteRelationships_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgDeleteRelationship","value":{"counterparty":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(msgDeleteRelationships.GetSignBytes()))
}

func TestMsgDeleteRelationships_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgDeleteRelationships.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeleteRelationships.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgBlockUser = types.NewMsgBlockUser(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"reason",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
)

func TestMsgBlockUser_Route(t *testing.T) {
	require.Equal(t, "profiles", msgBlockUser.Route())
}

func TestMsgBlockUser_Type(t *testing.T) {
	require.Equal(t, "block_user", msgBlockUser.Type())
}

func TestMsgBlockUser_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgBlockUser
		shouldErr bool
	}{
		{
			name: "empty sender returns error",
			msg: types.NewMsgBlockUser(
				"",
				"",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "empty receiver returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "equal sender and receiver returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
				"yeah",
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgBlockUser,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgBlockUser_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgBlockUser","value":{"blocked":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","blocker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","reason":"reason","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`
	require.Equal(t, expected, string(msgBlockUser.GetSignBytes()))
}

func TestMsgBlockUser_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgBlockUser.Blocker)
	require.Equal(t, []sdk.AccAddress{addr}, msgBlockUser.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgUnblockUser = types.NewMsgUnblockUser(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
)

func TestMsgUnblockUser_Route(t *testing.T) {
	require.Equal(t, "profiles", msgUnblockUser.Route())
}

func TestMsgUnblockUser_Type(t *testing.T) {
	require.Equal(t, "unblock_user", msgUnblockUser.Type())
}

func TestMsgUnblockUser_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgUnblockUser
		shouldErr bool
	}{
		{
			name: "empty sender returns error",
			msg: types.NewMsgUnblockUser(
				"",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "empty receiver returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "equal sender and receiver returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"yeah",
			),
			shouldErr: true,
		},
		{
			name: "valid message returs no error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgUnblockUser_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgUnblockUser","value":{"blocked":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","blocker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`
	require.Equal(t, expected, string(msgUnblockUser.GetSignBytes()))
}

func TestMsgUnblockUser_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUnblockUser.Blocker)
	require.Equal(t, []sdk.AccAddress{addr}, msgUnblockUser.GetSigners())
}
