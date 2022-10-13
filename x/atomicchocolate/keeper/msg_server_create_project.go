package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/types"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var project = types.Projects{
		Creator:     msg.Creator,
		Title:       msg.Title,
		Description: msg.Description,
		Literature:  msg.Literature,
	}

	id := k.AppendProject(ctx, project)

	return &types.MsgCreateProjectResponse{Id: id}, nil
}
