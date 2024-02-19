package cmd

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/satshub/opscript/blockchain/blockstream"
	"github.com/satshub/opscript/blockchain/node"
	"github.com/satshub/opscript/debugger"
	"github.com/satshub/opscript/gui"
	"github.com/satshub/opscript/spec"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Blockchain ...
type Blockchain interface {
	GetTransaction(txHash string) (*wire.MsgTx, error)
}

// New ...
func New(spec spec.Script) *cobra.Command {
	rootCmd := newRootCmd(spec)
	buildSpecCmd := newBuildSpecCmd()

	rootCmd.AddCommand(buildSpecCmd)

	return rootCmd
}

var useNode, useBlockstream, useHex bool
var nodeAddr, rpcUser, rpcPass string

func drawDebuggerGui(tx *wire.MsgTx, txInput int, bchain Blockchain, spec spec.Script) error {
	prevOut := tx.TxIn[txInput].PreviousOutPoint

	prevTx, err := bchain.GetTransaction(prevOut.Hash.String())
	if err != nil {
		logrus.Fatal(fmt.Errorf("get prev transaction: %+v", err))
	}

	en, err := newEngine(tx, prevTx.TxOut[prevOut.Index].PkScript, txInput)
	if err != nil {
		logrus.Fatal(fmt.Errorf("new engine: %+v", err))
	}

	d, err := debugger.New(en)
	if err != nil {
		logrus.Fatalln(err)
	}

	gui, err := gui.New(d, spec)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer gui.Stop()

	if err := gui.Start(); err != nil {
		logrus.Fatalln(err)
	}
	return nil
}

func debugOnchainTx(args []string, bchain Blockchain, spec spec.Script) error {
	txHash, txInput, err := parseArgs(args)
	if err != nil {
		return err
	}

	tx, err := bchain.GetTransaction(txHash)
	if err != nil {
		return err
	}

	return drawDebuggerGui(tx, txInput, bchain, spec)
}

func debugHexTx(args []string, bchain Blockchain, spec spec.Script) error {
	hexTx, txInput, err := parseHexTxArgs(args)
	if err != nil {
		return err
	}

	// Decode the serialized transaction hex to raw bytes.
	serializedTx, err := hex.DecodeString(hexTx)
	if err != nil {
		return err
	}

	// Deserialize the transaction and return it.
	var msgTx wire.MsgTx
	if err := msgTx.Deserialize(bytes.NewReader(serializedTx)); err != nil {
		return err
	}

	return drawDebuggerGui(&msgTx, txInput, bchain, spec)
}

func newRootCmd(spec spec.Script) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "opscript [flags] transactionHash:inputIndex",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var bchain Blockchain

			switch {
			case useBlockstream:
				bchain = blockstream.New()
				return debugOnchainTx(args, bchain, spec)

			case useNode:
				bchain = node.New(nodeAddr, rpcUser, rpcPass)
				return debugOnchainTx(args, bchain, spec)

			case useHex:
				bchain = blockstream.New()
				return debugHexTx(args, bchain, spec)
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&useNode, "node", false, "Use Bitcoin node to get transactions (requires 'txindex=1').")
	cmd.Flags().BoolVar(&useBlockstream, "blockstream", false, "Use blockstream.info API to get transactions.")
	cmd.Flags().StringVar(&nodeAddr, "node-addr", "127.0.0.1:8332", "Bitcoin node address.")
	cmd.Flags().StringVar(&rpcUser, "rpc-user", "", "Bitcoin JSON-RPC username.")
	cmd.Flags().StringVar(&rpcPass, "rpc-pass", "", "Bitcoin JSON-RPC password.")
	cmd.Flags().BoolVar(&useHex, "hex", false, "decode raw transaction from hex")

	return cmd
}

func newEngine(tx *wire.MsgTx, output []byte, inputIdx int) (*txscript.Engine, error) {
	e, err := txscript.NewEngine(
		output,
		tx,
		inputIdx,
		txscript.ScriptVerifyWitness+txscript.ScriptBip16+txscript.ScriptVerifyCleanStack+txscript.ScriptVerifyMinimalData,
		nil,
		nil,
		-1,
	)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func parseArgs(args []string) (txHash string, input int, err error) {
	if len(args) != 1 {
		return "", -1, errors.New("wrong number of arguments")
	}

	parts := strings.Split(args[0], ":")
	if len(parts) == 2 {
		input, err = strconv.Atoi(parts[1])
		if err != nil {
			return "", -1, err
		}
	}

	if hashBytes, err := hex.DecodeString(parts[0]); err != nil || len(hashBytes) != 32 {
		return "", -1, errors.New("wrong transaction hash")
	}

	txHash = parts[0]

	return
}

func parseHexTxArgs(args []string) (tx string, input int, err error) {
	if len(args) != 1 {
		return "", -1, errors.New("wrong number of arguments")
	}

	parts := strings.Split(args[0], ":")
	if len(parts) == 2 {
		input, err = strconv.Atoi(parts[1])
		if err != nil {
			return "", -1, err
		}
	}

	tx = parts[0]
	return
}
