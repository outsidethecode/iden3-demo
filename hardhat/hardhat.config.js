
require("@nomicfoundation/hardhat-toolbox");

module.exports = {
  solidity: "0.8.15",
  networks: {
    mumbai: {
       url: `https://polygon-mumbai.g.alchemy.com/v2/jNN6BxHCdHmxeTFcHHz-6DG7VTqX1tPY`,
       accounts: [`393483847b8a45a1a5e7f60924899c79ceb71e3e8f38acf4bcc3e51b243aa08c`],
    } 
 }
};
