package cli

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/pokt-network/pocket-core/app"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/state"
)

func init() {
	rootCmd.AddCommand(utilCmd)
	utilCmd.AddCommand(chainsGenCmd)
	utilCmd.AddCommand(chainsDelCmd)
	utilCmd.AddCommand(decodeTxCmd)
	utilCmd.AddCommand(unsafeRollbackCmd)
	utilCmd.AddCommand(exportGenesisForReset)
	utilCmd.AddCommand(completionCmd)
}

var utilCmd = &cobra.Command{
	Use:   "util",
	Short: "utility functions",
	Long:  `The util namespace handles all utility functions`,
}

var chainsGenCmd = &cobra.Command{
	Use:   "generate-chains",
	Short: "Generates chains file",
	Long:  `Generate the chains file for network identifiers`,
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		c := app.NewHostedChains(true)
		fmt.Println(app.GlobalConfig.PocketConfig.ChainsName + " contains: \n")
		for _, chain := range c.M {
			fmt.Println(chain.ID + " @ " + chain.URL)
		}
		fmt.Println("If incorrect: please remove the chains.json with the " + chainsDelCmd.NameAndAliases() + " command")
	},
}

var chainsDelCmd = &cobra.Command{
	Use:   "delete-chains",
	Short: "Delete chains file",
	Long:  `Delete the chains file for network identifiers`,
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		app.DeleteHostedChains()
		fmt.Println("successfully deleted " + app.GlobalConfig.PocketConfig.ChainsName)
	},
}

var decodeTxCmd = &cobra.Command{
	Use:   "decode-tx <tx>",
	Short: "Decodes a given transaction encoded in Amino base64 bytes",
	Long:  `Decodes a given transaction encoded in Amino base64 bytes`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		txStr := args[0]
		stdTx := app.UnmarshalTxStr(txStr)
		fmt.Printf(
			"Type:\t\t%s\nMsg:\t\t%v\nFee:\t\t%s\nEntropy:\t%d\nMemo:\t\t%s\nSigner\t\t%s\nSig:\t\t%s\n",
			stdTx.Msg.Type(), stdTx.Msg, stdTx.Fee.String(), stdTx.Entropy, stdTx.Memo, stdTx.Msg.GetSigner().String(),
			stdTx.Signature.RawString())
	},
}

var exportGenesisForReset = &cobra.Command{
	Use:   "export-genesis-for-reset <height> <newChainID>",
	Short: "exports new genesis based on state",
	Long:  `In the event of a network reset, this will export a genesis file based on the previous state`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		height, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("error parsing height: ", err)
			return
		}
		db, err := app.OpenDB(app.GlobalConfig.TendermintConfig.RootDir)
		if err != nil {
			fmt.Println("error loading application database: ", err)
			return
		}
		loggerFile, _ := os.Open(os.DevNull)
		a := app.NewPocketCoreApp(nil, nil, nil, nil, log.NewTMLogger(loggerFile), db)
		// initialize stores
		blockStore, _, _, _, err := state.BlocksAndStateFromDB(&app.GlobalConfig.TendermintConfig, state.DefaultDBProvider)
		if err != nil {
			fmt.Println("err loading blockstore: ", err.Error())
			return
		}
		a.SetBlockstore(blockStore)
		chainID := args[1]
		j, err := a.ExportState(int64(height), chainID)
		if err != nil {
			fmt.Println("could not export genesis state: ", err.Error())
			return
		}
		fmt.Println(j)
	},
}

func init() {
	unsafeRollbackCmd.Flags().BoolVar(&blocks, "blocks", false, "rollback blocks as well as the state")
}

var (
	blocks bool
)

var unsafeRollbackCmd = &cobra.Command{
	Use:   "unsafe-rollback <height>",
	Short: "Rollbacks the blockchain, the state, and app to a previous height",
	Long:  "Rollbacks the blockchain, the state, and app to a previous height",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		height, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("error parsing height: ", err)
			return
		}
		db, err := app.OpenDB(app.GlobalConfig.TendermintConfig.RootDir)
		if err != nil {
			fmt.Println("error loading application database: ", err)
			return
		}
		loggerFile, _ := os.Open(os.DevNull)
		a := app.NewPocketBaseApp(log.NewTMLogger(loggerFile), db)
		// initialize stores
		a.MountKVStores(a.Keys)
		a.MountTransientStores(a.Tkeys)
		// rollback the txIndexer

		err = state.RollbackTxIndexer(&app.GlobalConfig.TendermintConfig, int64(height), context.Background())
		if err != nil {
			fmt.Println("error rolling back txIndexer: ", err)
			return
		}
		// rollback the app store
		err = a.Store().RollbackVersion(int64(height))
		if err != nil {
			fmt.Println("error rolling back app: ", err)
			return
		}
		if blocks {
			// rollback block store and state
			err = app.UnsafeRollbackData(&app.GlobalConfig.TendermintConfig, true, int64(height))
			if err != nil {
				fmt.Println("error rolling back block and state: ", err)
				return
			}
		}
	},
}

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

# Dependencies:
# bash 4.1+
# bash_completion from Kubectl

# add this to bash profile
$ source <(pocket util completion bash)

# To load completions for each session, execute once:
Linux:
  $ pocket util completion bash > /etc/bash_completion.d/pocket
MacOS:
  $ pocket util completion bash > /usr/local/etc/bash_completion.d/pocket

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ pocket util completion zsh > "${fpath[1]}/pocket"

# You will need to start a new shell for this setup to take effect.

Fish:

$ pocket util completion fish | source

# To load completions for each session, execute once:
$ pocket util completion fish > ~/.config/fish/completions/pocket.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			_ = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			_ = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			_ = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			_ = cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}
