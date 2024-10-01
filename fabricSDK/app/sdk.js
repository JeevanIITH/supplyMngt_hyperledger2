const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');

class ChaincodeSdk {
    constructor(connectionProfile, walletPath, channelName, contractName) {
        this.connectionProfile = connectionProfile;
        this.walletPath = walletPath;
        this.channelName = channelName;
        this.contractName = contractName;
    }

    async connect() {
        const wallet = await Wallets.newFileSystemWallet(this.walletPath);
        const gateway = new Gateway();

        await gateway.connect(this.connectionProfile, {
            wallet,
            identity: 'admin', // replace with your identity name
            discovery: { enabled: true, asLocalhost: true }
        });

        this.contract = gateway.getNetwork(this.channelName).getContract(this.contractName);
        this.gateway = gateway;
    }

    async disconnect() {
        await this.gateway.disconnect();
    }

    async invokeTransaction(functionName, ...args) {
        try {
            const result = await this.contract.submitTransaction(functionName, ...args);
            return JSON.parse(result.toString());
        } catch (error) {
            console.error(`Failed to submit transaction: ${error}`);
            throw new Error(`Transaction failed: ${error}`);
        }
    }

    async queryTransaction(functionName, ...args) {
        try {
            const result = await this.contract.evaluateTransaction(functionName, ...args);
            return JSON.parse(result.toString());
        } catch (error) {
            console.error(`Failed to evaluate transaction: ${error}`);
            throw new Error(`Query failed: ${error}`);
        }
    }
}

module.exports = ChaincodeSdk;