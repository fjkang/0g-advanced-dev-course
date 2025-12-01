import { CONTRACTS, deployInBeaconProxy } from "./utils/utils";
import { ethers } from "hardhat";

async function main() {
    // 通过run函数获取hre
    const hre = await import("hardhat");
    
    const { getNamedAccounts } = hre;
    const { deployer } = await getNamedAccounts();

    // mint a token for the deployer
    const agentNFTDeployment = await hre.deployments.get(CONTRACTS.AgentNFT.name);
    console.log("✅ AgentNFT deployed to:", agentNFTDeployment.address);
    const agentNFT = await ethers.getContractAt(CONTRACTS.AgentNFT.name, agentNFTDeployment.address);
    await agentNFT.mint([{dataDescription: "minted inft", dataHash: ethers.hashMessage("0xdb mint inft")}], deployer);
    console.log(`✅ AgentNFT minted for ", ${deployer}`);
    // const tokenId = 0;
    // const owner = await agentNFT.ownerOf(tokenId);
    // const intelligentDatas = await agentNFT.intelligentDatasOf(tokenId);
    // console.log(`✅ tokenId: ${tokenId}, owner is ${owner}`);
    // console.log(`dataDescription is ${intelligentDatas[0].dataDescription}, dataHash is ${intelligentDatas[0].dataHash}`);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});