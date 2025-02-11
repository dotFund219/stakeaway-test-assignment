const { expect } = require("chai");
const { ethers } = require("hardhat");
const { getBalance } = require("./common");

describe("✅ Staking Contract", function () {
  let stakingContract, owner, staker;

  beforeEach(async function () {
    [owner, staker] = await ethers.getSigners();

    const StakingContract = await ethers.getContractFactory("Staking");
    stakingContract = await StakingContract.deploy();
  });

  it("Stake Success", async function () {
    await stakingContract.connect(staker).stake({ value: 1000 });
  });
});
