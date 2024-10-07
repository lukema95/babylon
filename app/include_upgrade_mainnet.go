//go:build mainnet

package app

import (
	"github.com/babylonlabs-io/babylon/app/upgrades"
	v1 "github.com/babylonlabs-io/babylon/app/upgrades/v1"
	"github.com/babylonlabs-io/babylon/app/upgrades/v1/mainnet"
)

// init is used to include v1 upgrade for mainnet data
func init() {
	Upgrades = []upgrades.Upgrade{v1.CreateUpgrade(v1.UpgradeDataString{
		BtcStakingParamStr:    mainnet.BtcStakingParamStr,
		FinalityParamStr:      mainnet.FinalityParamStr,
		NewBtcHeadersStr:      mainnet.NewBtcHeadersStr,
		SignedFPsStr:          mainnet.SignedFPsStr,
		TokensDistributionStr: mainnet.TokensDistributionStr,
	})}
}