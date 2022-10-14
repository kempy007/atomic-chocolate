package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Projects(goCtx context.Context, req *types.QueryProjectsRequest) (*types.QueryProjectsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var projects []*types.Project

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	projectStore := prefix.NewStore(store, []byte(types.ProjectKey))

	pageRes, err := query.Paginate(projectStore, req.Pagination, func(key []byte, value []byte) error {
		var project types.Project
		if err := k.cdc.Unmarshal(value, &project); err != nil {
			return err
		}

		projects = append(projects, &project)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryProjectsResponse{Project: projects, Pagination: pageRes}, nil
}
