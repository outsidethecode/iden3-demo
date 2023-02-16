const hre = require("hardhat");

async function main() {

  // Import State contract from existing address
  const contract = await hre.ethers.getContractAt("State", "0xdB4c661456A023a403CF83784f9d15a1D3540702");

  // Add inputs from proof
  const id = "0x00ca0b252b3f5ac5f13c9f675c297c5f2f043335dfc4d18c1265a5de49ab0000"
  const oldState = "0x252b3f5ac5f13c9f675c297c5f2f043335dfc4d18c1265a5de49abbb3d0aaacd"
  const newState = "0x07a64c7d4c5be707b9188cca29261d3fc89dc3086af0b229ff73f7dbcc47eb02"
  const isOldStateGenesis = "0x0000000000000000000000000000000000000000000000000000000000000001"

  const a = ["0x162b69da7ee75d3f84b0ac04263c869ca022977ff7eea5524edd1f187bca6d2f", "0x00d845508f1530315c01f2eca0d46a1a063b0c85db428fcd220181f8008beb5e"]
  const b = [["0x0cb4f4b0a0be51e36dc691b8bfe6502c1465ba32bef6b2e21fd35f0ccf1ccde8", "0x2f4000e5deaa4a10d76229414be393896a4f17056cf08c8ac0d1128384fe1d10"],["0x03762181a364bbbe6719620c7e878a817510b517232313a9bb9a558dce35c949", "0x2e396d764aa22357e4ff0d094dfde884f38a8d5e6b647d3a756c288a5066f426"]]
  const c = ["0x168823b92ec61074a71fdb8db35739f1b6e1a0d8d330ef856751539e245ccdeb", "0x0026a45b87a02618917c787b9ee8ffe7cfa62234da1ca3d62525596346741b00"]

  // Check Identity State for your ID before state transition
  let identityState0 = await contract.getState(id);

  console.log("Identity State at t=0", identityState0);

  // Execute state transition
  await contract.transitState(id, oldState, newState, isOldStateGenesis, a, b, c);

  // Check Identity State for your ID after state transition
  identityState = await contract.getState(id);

  console.log("Identity State at t=1", identityState);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});


