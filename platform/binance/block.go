package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	info, err := p.client.fetchNodeInfo()
	if err != nil {
		return 0, err
	}

	return info.SyncInfo.LatestBlockHeight, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	//blockTxs, err := p.client.GetBlockTransactions(num)
	//if err != nil {
	//	return nil, err
	//}

	txs := make(blockatlas.TxPage, 0)
	//childTxs, err := p.getTxChildChan(blockTxs)
	//if err == nil {
	//	txs = NormalizeTxs(childTxs, "", "")
	//} else {
	//	txs = NormalizeTxs(blockTxs, "", "")
	//}
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
