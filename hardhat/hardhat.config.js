require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-etherscan");

// require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");

module.exports = {
  solidity: {
   compilers: [
      {version: "0.8.15"},
      {version: "0.6.0"},
      {version: "0.6.11"}
   ]
  },
  networks: {
    mumbai: {
       url: `https://polygon-mumbai.g.alchemy.com/v2/jNN6BxHCdHmxeTFcHHz-6DG7VTqX1tPY`,
       accounts: [`7d386504436dd36c6cbebec3502f8b1a1e3ff0e22df2bc3180d427a413037c87`],
       gas: 2100000,
       gasPrice: "auto"
   } 
  },
  etherscan: {
   apikey: "PMUDU7N7GSQ87D4ID27M2N7A7I3N3DQCR7"
  }
};
