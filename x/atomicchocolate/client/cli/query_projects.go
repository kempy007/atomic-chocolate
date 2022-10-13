package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdProjects() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects",
		Short: "Query projects",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryProjectsRequest{}

			res, err := queryClient.Projects(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
