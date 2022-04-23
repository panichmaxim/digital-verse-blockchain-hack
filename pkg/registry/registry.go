package registry

import (
	"gitlab.com/digitalverse/blockchain/pkg/blockchain"
	"gitlab.com/digitalverse/blockchain/pkg/blockchain/generic"
)

func CreateRegistry(c *blockchain.Config) map[string]blockchain.Network {
	return map[string]blockchain.Network{
		"heco": generic.NewGenericNetwork(
			"https://testnet.hecoinfo.com/tx/%s",
			c.HecoEndpointUrl,
			c.HecoDeployWalletPk,
			c.HecoNftContractAddress,
			-1,
		),
		"eth": generic.NewGenericNetwork(
			"https://rinkeby.etherscan.io/tx/%s",
			c.RinkebyEndpointUrl,
			c.RinkebyDeployWalletPk,
			c.RinkebyNftContractAddress,
			-1,
		),
		"near": generic.NewGenericNetwork(
			"https://explorer.testnet.aurora.dev/tx/%s",
			c.AuroraEndpointUrl,
			c.AuroraDeployWalletPk,
			c.AuroraNftContractAddress,
			c.AuroraShardID,
		),
		"polygon": generic.NewGenericNetwork(
			"https://mumbai.polygonscan.com/tx/%s",
			c.PolygonEndpointUrl,
			c.PolygonDeployWalletPk,
			c.PolygonNftContractAddress,
			c.PolygonShardID,
		),
		"solana": generic.NewGenericNetwork(
			"https://explorer.solana.com/?cluster=testnet/tx/%s",
			c.SolanaTestnetEndpointUrl,
			c.SolanaTestnetDeployWalletPk,
			c.SolanaTestnetNftContractAddress,
			-1,
		),
		"harmony": generic.NewGenericNetwork(
			"https://explorer.pops.one/tx/%s",
			c.HarmonyTestnetEndpointUrl,
			c.HarmonyTestnetDeployWalletPk,
			c.HarmonyTestnetNftContractAddress,
			c.HarmonyShardID,
		),
		"velas": generic.NewGenericNetwork(
			"https://explorer.testnet.velas.com/tx/%s",
			c.VelasTestnetEndpointUrl,
			c.VelasTestnetDeployWalletPk,
			c.VelasTestnetNftContractAddress,
			c.VelasShardID,
		),
		"okex": generic.NewGenericNetwork(
			"https://www.oklink.com/okexchain-test/address/%s",
			c.OKExChainTestnetEndpointUrl,
			c.OKExChainTestnetDeployWalletPk,
			c.OKExChainTestnetNftContractAddress,
			c.OKExChainTestnetShardID,
		),
		"coinex": generic.NewGenericNetwork(
			"https://testnet.coinex.net/tx/%s",
			c.CoinExTestnetEndpointUrl,
			c.CoinExTestnetDeployWalletPk,
			c.CoinExTestnetNftContractAddress,
			c.CoinExTestnetShardID,
		),
		"theta": generic.NewGenericNetwork(
			"https://testnet-explorer.thetatoken.org/tx/%s",
			c.ThetaTestnetEndpointUrl,
			c.ThetaTestnetDeployWalletPk,
			c.ThetaTestnetNftContractAddress,
			c.ThetaTestnetShardID,
		),
		"cronos": generic.NewGenericNetwork(
			"https://cronos.crypto.org/explorer/testnet3/tx/%s",
			c.CronosTestnetEndpointUrl,
			c.CronosTestnetDeployWalletPk,
			c.CronosTestnetNftContractAddress,
			c.CronosTestnetShardID,
		),
		"celo": generic.NewGenericNetwork(
			"https://alfajores-blockscout.celo-testnet.org/tx/%s",
			c.CeloTestnetEndpointUrl,
			c.CeloTestnetDeployWalletPk,
			c.CeloTestnetNftContractAddress,
			c.CeloTestnetShardID,
		),
	}
}
