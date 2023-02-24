const hre = require("hardhat");

async function main() {

  // Import State contract from existing address
  // Mumbai
  const contract = await hre.ethers.getContractAt("StateV2", "0xd04E12407D5B5ec66c0400Cb88347B89Ccf5e42e");

  const id = "0x000f370e06d4e2f589488a5ca49b72dddece47db5bd18c79b4802feda2420000"
  const oldState = "0x0e06d4e2f589488a5ca49b72dddece47db5bd18c79b4802feda2421e8db009ab"
  const newState = "0x276f10801ff6f5c62159ff58efcd900915756116344cc4474ad31e78107742df"
  const isOldStateGenesis = "0x0000000000000000000000000000000000000000000000000000000000000001"
  
  const a = ["0x0e535a98dcfd463c64bbde46aadb5020671db74b5f8a70633ff897928025fa25", "0x1c236fb9d89c9270a13b36ef7ba47fb2a6cdf6010a731b89f91457109688ccfc"]
  const b = [["0x07d2551699b67717902ce85993632f74852cb5569910fdba1a7584e115b6a105", "0x00500d6563d5ab43319b660621ec5a61e0e79b856b253e2eae9e11a91e7746f5"],["0x2e2f89b159132fcbcf8ed69ea5dbe1f8dba42081465b7aa32c52dc8a1fd89424", "0x2f40fafcfc060c4de230a72da177d4edcb7fbbe9e70a59796c6a53abca311892"]]
  const c = ["0x000dfa821bcaa380ed899f65f1129104350ea4b8593ba54eb52faf932a062ba2", "0x1b7c70ecce5027775e3c5850ce2db735c69d3f576c8b5694113b194f81088429"]

  console.log("Init ... ");

  // Check Identity State for your ID before state transition
  let identityState0 = 0;//await contract.getStateInfoById(id);

  console.log("Identity State at t=0", identityState0);

  // Execute state transition
  let tx = await contract.transitState(id, oldState, newState, isOldStateGenesis, a, b, c);

  // wait for transaction to be mined
  await tx.wait();

  console.log("-----------------------------");
  // Check Identity State for your ID after state transition
  identityState = await contract.getStateInfoById(id);

  console.log("Identity State at t=1", identityState);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});


