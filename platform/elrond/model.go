package elrond

import (
	"math/big"
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type NetworkStatus struct {
	NetworkStatus Status `json:"message"`
}

type Status struct {
	Status StatusDetails `json:"status"`
}

type StatusDetails struct {
	Round float64 `json:"erd_current_round"`
	Epoch float64 `json:"erd_epoch_number"`
	Nonce float64 `json:"erd_nonce"`
}

type BlockResponse struct {
	Block Block `json:"block"`
}

type Block struct {
	Nonce        uint64        `json:"nonce"`
	Hash         string        `json:"hash"`
	Transactions []Transaction `json:"transactions"`
}

type TransactionsPage struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Hash      string        `json:"hash"`
	Nonce     uint64        `json:"nonce"`
	Value     string        `json:"value"`
	Receiver  string        `json:"receiver"`
	Sender    string        `json:"sender"`
	Data      string        `json:"data"`
	Timestamp time.Duration `json:"timestamp"`
	Status    string        `json:"status"`
	GasPrice  uint64        `json:"gasPrice"`
	GasUsed   uint64        `json:"gasUsed"`
}

func (tx *Transaction) TxStatus() blockatlas.Status {
	switch tx.Status {
	case "Success":
		return blockatlas.StatusCompleted
	case "Pending":
		return blockatlas.StatusPending
	default:
		return blockatlas.StatusError
	}
}

func (tx *Transaction) Fee() blockatlas.Amount {
	gasPrice := big.NewInt(0).SetUint64(tx.GasPrice)
	gasUsed := big.NewInt(0).SetUint64(tx.GasUsed)
	fee := big.NewInt(0).Mul(gasPrice, gasUsed)

	return blockatlas.Amount(fee.String())
}

func (tx *Transaction) Direction(address string) blockatlas.Direction {
	switch {
	case tx.Sender == address && tx.Receiver == address:
		return blockatlas.DirectionSelf
	case tx.Sender == address && tx.Receiver != address:
		return blockatlas.DirectionOutgoing
	default:
		return blockatlas.DirectionIncoming
	}
}
