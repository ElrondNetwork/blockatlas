package elrond

import (
	"fmt"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	var latestNonce LatestNonce
	err = c.Get(&latestNonce, "block/latest-nonce", nil)
	if err != nil {
		return 0, err
	}

	return int64(latestNonce.Nonce), nil
}

func (c *Client) GetBlockByNumber(height int64) (*blockatlas.Block, error) {
	var blockRes BlockResponse

	path := fmt.Sprintf("block/meta/%d", uint64(height))
	err := c.Get(&blockRes, path, nil)
	if err != nil {
		return nil, err
	}

	block := blockRes.Block
	txs := NormalizeTxs(block.Transactions, "")

	return &blockatlas.Block{
		Number: int64(block.Nonce),
		ID:     block.Hash,
		Txs:    txs,
	}, nil
}

func (c *Client) GetTxsOfAddress(address string) (blockatlas.TxPage, error) {
	var txPage TransactionsPage
	// TODO: enable pagination of Elrond transactions in the future.
	// TODO: currently Elrond only fetches the most recent 20 transactions.
	path := fmt.Sprintf("address/%s/transactions", address)
	err := c.Get(&txPage, path, nil)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(txPage.Transactions, address)

	return txs, nil
}
