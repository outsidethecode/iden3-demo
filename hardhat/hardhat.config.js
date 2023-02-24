require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-etherscan");

// require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");


// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});


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
      gas: 9900000,
      gasPrice: "auto"
    },
    localhost: {
      url: "http://127.0.0.1:8545/",
      accounts: [
       "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
      ],
      gas: 960000,
      gasPrice: "auto",
      blockGasLimit: 960000
   } 
  },
  etherscan: {
   apikey: "PMUDU7N7GSQ87D4ID27M2N7A7I3N3DQCR7"
  },
};
