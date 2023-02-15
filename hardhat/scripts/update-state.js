const hre = require("hardhat");

async function main() {

  // Import State contract from existing address
  const contract = await hre.ethers.getContractAt("StateV2", "0xff8d57666B95E4f9560F4DB7412605A7B1d17918");

  // Add inputs from proof
  const id = "16902857626710129093777049632637234295525180659192520755330107189706293248"
  const oldState = "8066878375842601856475243198843085975596912141564964024736389476326270909304"
  const newState = "21115237637388129015117528037637357760051922988797104487501981291013305057525"
  const isOldStateGenesis = "1"

  const a = [
    "1630421067897642627687294525861707117104539066333376262752291887227366059760",
    "2740649345442919581954819090271932172601335600630355609761718067394798804404"
  ]
  
  const b = [
    [
     "10030581716058119958124758248302905971642575891511562064813930982127193838979",
     "10797080316543379766288890244734017902078199095917137620666224767776519275169"
    ],
    [
     "19320023229863823341945472518807692384732716652585407280165417972017122537238",
     "5593935735234647868873761806670296776501834052955717427835563593147545954783"
    ]
  ]

  const c = [
    "112117455385219068098203385689784799845678331923639369702057707867982316982",
    "3035653997312817797202225901056025078071404743173839494883489552850461225375"
  ]

  console.log("Init ... ");

  // Check Identity State for your ID before state transition
  let identityState0;
  if (await contract.idExists(id)) {
    console.log("ID exists");
    identityState0 = await contract.getState(id);
  } else {
    console.log("ID does not exist");
    identityState0 = 0;
  }

  console.log("Identity State at t=0", identityState0);

  // Execute state transition
  await contract.transitState(id, oldState, newState, isOldStateGenesis, a, b, c);
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


