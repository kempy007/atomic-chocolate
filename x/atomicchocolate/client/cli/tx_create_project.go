package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-project [title] [description] [literature]",
		Short: "Broadcast message createProject",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argDescription := args[1]
			argLiterature := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProject(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argDescription,
				argLiterature,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
