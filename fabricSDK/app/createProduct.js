// there is createProduct function in chaincode 
// which is called by this function

'use strict';

const { 
	Gateway,
	Wallets
} = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const path = require('path');

// const ChaincodeSdk = require('./sdk');

// const connectionProfilePath = path.resolve(__dirname, '..', 'certs', 'connection-producer.json');
// const walletPath = path.join(process.cwd(), 'wallet'); // Adjust as necessary
// const ccp = require('../utils/AppUtil').buildCCPPRODUCER();
// const wallet = Wallets.newFileSystemWallet(walletPath);

const {
	buildCCPPRODUCER,
	buildWallet
} = require('../utils/AppUtil.js');

const channelName = 'channel1'; // Replace with your channel name
const contractName = 'supply1'; // Replace with your chaincode name
const walletPath = path.join(__dirname, '../wallet');




exports.createProduct = async (productid, name_1, description, manufacturingDate, batchNumber , userId) => {

    // const sdk = new ChaincodeSdk(connectionProfilePath, walletPath, channelName, contractName);

    // try {
    //     await sdk.connect();

    //     // Example of invoking a transaction
    //     const result = await sdk.invokeTransaction('createProduct', productid, name_1, description, manufacturingDate, batchNumber);
    //     console.log('Transaction Result:', result);

    //     // Example of querying a transaction
    //     // const queryResult = await sdk.queryTransaction('queryProduct', 'product1');
    //     // console.log('Query Result:', queryResult);
    // } catch (error) {
    //     console.error('Error:', error);
    // } finally {
    //     await sdk.disconnect();
    // }

    const ccp = buildCCPPRODUCER();
    const wallet = await buildWallet(Wallets, walletPath);


    try {
        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        console.log('Gateway Not connected');
        await gateway.connect(ccp, {
            wallet,
            identity: userId,
            discovery: { enabled: true, asLocalhost: true }
        });
        console.log('Gateway connected');

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(channelName);
        console.log('Contract got');
        // Get the contract from the network.
        const contract = network.getContract(contractName);
        
        // Submit the specified transaction.
        await contract.submitTransaction('CreateProduct', productid, name_1, description, manufacturingDate, batchNumber);
        console.log('Transaction has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}