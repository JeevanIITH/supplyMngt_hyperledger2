/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const {
    Wallets
} = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const path = require('path');
const fs = require('fs');
const {
    buildCAClient,
    registerAndEnrollUser } = require('../utils/CAUtil.js');


// const  = require('../utils/CAUtil.js');
const {
    buildWallet,
    buildCCPPRODUCER
} = require('../utils/AppUtil.js');

const mspOrg1 = 'ProducerMSP';
const walletPath = path.join(__dirname, '../wallet');

async function registerProducerUser(userId, affiliation) {

    let response;
    try {
        // build an in memory object with the network configuration (also known as a connection profile)
        const ccp = buildCCPPRODUCER();

        // build an instance of the fabric ca services client based on
        // the information in the network configuration
        const caClient = buildCAClient(FabricCAServices, ccp, 'ca.producer.example.com');

        // setup the wallet to hold the credentials of the application user
        const wallet = await buildWallet(Wallets, walletPath);

        const identity = await wallet.get('User1');
        if (identity){
            console.log('An identity for the user "User 1" exist');
            return;
        }
        const certPath = path.resolve(__dirname,'../../organizations/peerOrganizations/producer.example.com/users/User1@producer.example.com/msp/signcerts/User1@producer.example.com-cert.pem');
        const keyPath = path.resolve(__dirname,'../../organizations/peerOrganizations/producer.example.com/users/User1@producer.example.com/msp/keystore/');

        const cert = fs.readFileSync(certPath,'utf8');
        const [keyFile] = fs.readdirSync(keyPath);
        const privateKey = fs.readFileSync(path.join(keyPath,keyFile),'utf8');

        const identityData = {
            credentials:{
                certificate: cert,
                privateKey: privateKey,

            },
            mspId: 'ProducerMSP',
            type: 'X.509',
        };

        await wallet.put('User1',identityData);
        console.log('Succefully imported User1 identity');

        
        // in a real application this would be done only when a new user was required to be added
        // and would be part of an administrative flow
        // await registerAndEnrollUser(caClient, wallet, mspOrg1, userId, affiliation);

        response = {
            success: true,
            message: `Successfully enrolled client user ${userId} and imported it into the wallet`
        };
    } catch (error) {
        console.error(`******** FAILED to run the application: ${error}`);
        response = {
            success: false,
            message: `${error}`
        };
    }

    return response;
}

module.exports = registerProducerUser;