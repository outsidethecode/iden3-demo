const hre = require("hardhat");

async function main() {

  // Import State contract from existing address
  const contract = await hre.ethers.getContractAt("State", "0xdB4c661456A023a403CF83784f9d15a1D3540702");

  // Add inputs from proof
  const id = "0x000c1a0dd64bf3b534db11151890c686dd9f04e540a41dad1abf2d650b920000"
  const oldState = "0x0dd64bf3b534db11151890c686dd9f04e540a41dad1abf2d650b921f1bd70737"
  const newState = "0x300499dc9dc3c2adce2b9be172107b0d120c04ac29442714d1ca5accf668373b"
  const isOldStateGenesis = "0x0000000000000000000000000000000000000000000000000000000000000001"

  const a = ["0x1e4781c9cbdc4ce5273a9cb64fcbd7a483bce9ede6bf4c0e9c5e10712260d128", "0x1439c81adf42d295cc87410ef501982ed2da01a32d3521f35583c942e110c045"]
  const b = [["0x03d062c4fa439de164f39506d2751caef067d96ac4e60c9900d0233303c8e9b8", "0x0b1b30b055807db86af83793050180f23bf99e8298ef8dbfdde8feac10d4b35f"],["0x22fc160c1eab53b02adcbc22d994aaa0f8ba4396517e73e90829d2de8a14395f", "0x03e90e476df94422c73af69a9f168b71cf19b41db9cf0e3debcde6a335b78f52"]]
  const c = ["0x2344601c50d8660585c34e160182aa85b74c71b0d76b606d94230e4af1088fdb", "0x2a742088ce9c77a3a56991bda1b7cf21c5e9b793eab148b379e3428b4a64642a"]

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


