# System Architecture

## 1. Participants

The consortium blockchain will connect the following participants:

*   **Exporter:** The seller of goods.
*   **Importer:** The buyer of goods.
*   **Exporter's Bank:** The bank that provides financing to the exporter.
*   **Importer's Bank:** The bank that issues the Letter of Credit (LC) on behalf of the importer.
*   **Central Bank:** The regulatory authority that oversees the financial system.
*   **Customs Authority:** The government agency responsible for regulating the flow of goods into and out of the country.

## 2. Hyperledger Fabric Network Components

The Hyperledger Fabric network will consist of the following components:

*   **Peers:** Each participant will have its own peer node, which will maintain a copy of the ledger and execute smart contracts.
*   **Orderers:** A cluster of orderer nodes will be responsible for ordering transactions and creating blocks.
*   **Certificate Authorities (CAs):** Each participant will have its own CA to issue digital certificates to its users.

## 3. Communication Channels

The network will use private channels to ensure that sensitive data is only shared with the relevant participants. For example, a private channel will be created for each trade transaction, and only the exporter, importer, and their respective banks will have access to it.

## 4. Identity Management

The network will use a Public Key Infrastructure (PKI) to manage the identities of the participants. Each user will have a unique digital certificate that will be used to authenticate them and authorize their transactions.

## 5. Off-Chain Storage

Large documents, such as invoices and bills of lading, will be stored off-chain in an IPFS or AWS S3 bucket. The hash of the document will be stored on the blockchain to ensure its integrity.
