package keeper

import (
	"context"

	"blog/x/blog/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPostByTag(goCtx context.Context, req *types.QueryListPostByTagRequest) (*types.QueryListPostByTagResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))

	var posts []types.Post
	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var post types.Post
		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return false, err
		}

		for _, tag := range post.GetTags() {
			if tag == req.Tag {
				posts = append(posts, post)
				return true, nil
			}
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListPostByTagResponse{Post: posts, Pagination: pageRes}, nil
}
