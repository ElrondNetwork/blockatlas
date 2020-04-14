package ontology

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

const (
	// TODO: Find a way to have a dynamic APR
	// The current value comes from https://cryptoslate.com/coins/ontology
	Annual = 4.45
)

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: Annual},
		MinimumAmount: "0",
		LockTime:      0,
		Type:          blockatlas.DelegationAuto,
	}
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	acc, err := p.client.GetBalances(address)
	if err != nil {
		return "0", err
	}
	balance := acc.Result.getBalance(AssetONT)
	if balance == nil {
		return "0", errors.E("Invalid asset balance", errors.Params{"asset": AssetONT})
	}
	return balance.Balance, nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	return blockatlas.ValidatorPage{}, nil
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	return blockatlas.DelegationsPage{}, nil
}
