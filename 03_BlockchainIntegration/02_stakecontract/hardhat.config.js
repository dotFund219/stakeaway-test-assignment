const { PRIVATE_KEY } = require("./config");
require("@openzeppelin/hardhat-upgrades");

require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    version: "0.8.20",
    settings: {
      viaIR: true,
      optimizer: {
        enabled: true,
        details: {
          yulDetails: {
            optimizerSteps: "u",
          },
        },
      },
    },
  },

  networks: {
    sepolia: {
      url: "https://sepolia.infura.io/v3/086acf6e4f764de28c1a82d044e74737",
      chainId: 11155111,
      gasPrice: 50000000000,
      accounts: [PRIVATE_KEY],
    },
  },
  etherscan: {
    apiKey: "PYH77TDUQQ1R5WA87UI5RC9UAVHW7VMIBP",
  },
};
