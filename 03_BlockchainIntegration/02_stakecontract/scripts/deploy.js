const { ethers, upgrades } = require("hardhat");

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function main() {
  const [deployer] = await ethers.getSigners();

  console.log("Starting deployment", deployer.address);

  const Staking = await ethers.getContractFactory("Staking");
  const staking = await Staking.deploy();

  console.log("Deploying staking contract");

  await staking.waitForDeployment();
  console.log("Staking contract deployed to:", staking.target);

  sleep(10000);

  // Verify the implementation contract
  await hre.run("verify:verify", {
    address: staking.target,
    constructorArguments: [],
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
