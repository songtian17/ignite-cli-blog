package keeper

import (
	"context"
	"time"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var post = types.Post{
		Creator:   msg.Creator,
		Title:     msg.Title,
		Body:      msg.Body,
		Tags:      msg.Tags,
		Timestamp: time.Now().String(),
	}
	id := k.AppendPost(
		ctx,
		post,
	)
	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
