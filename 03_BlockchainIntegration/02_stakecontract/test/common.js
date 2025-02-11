const { ethers } = require("hardhat");

const getBalance = async (address) => {
  return await ethers.provider.getBalance(address);
};

module.exports = {
  getBalance,
};
