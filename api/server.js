const express = require('express');
const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');

const app = express();
app.use(express.json());

const ccpPath = path.resolve(__dirname, '..', '..', 'test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

async function getContract() {
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    const identity = await wallet.get('appUser');
    if (!identity) {
        console.log('An identity for the user "appUser" does not exist in the wallet');
        console.log('Run the registerUser.js application before retrying');
        return;
    }

    const gateway = new Gateway();
    await gateway.connect(ccp, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });

    const network = await gateway.getNetwork('mychannel');
    const contract = network.getContract('tradefinance');
    return contract;
}

app.post('/lcs', async (req, res) => {
    try {
        const contract = await getContract();
        const { id, importer, exporter, amount, expiryDate } = req.body;
        await contract.submitTransaction('IssueLC', id, importer, exporter, amount, expiryDate);
        res.json({ lcId: id });
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
});

app.post('/lcs/:lcId/documents', async (req, res) => {
    try {
        const contract = await getContract();
        const { lcId } = req.params;
        await contract.submitTransaction('VerifyDocuments', lcId);
        res.json({ lcId });
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
});

app.post('/lcs/:lcId/settlement', async (req, res) => {
    try {
        const contract = await getContract();
        const { lcId } = req.params;
        await contract.submitTransaction('SettlePayment', lcId);
        res.json({ lcId });
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
});

app.listen(3000, () => {
    console.log('API server listening on port 3000');
});
