# Blockchain Implementation in Go

This Go project provides a basic implementation of a blockchain, demonstrating key concepts such as blocks, transactions, mining, and the Merkle root.

## Overview

The blockchain consists of the following main components:

- **Block Structure:** Represents a block in the blockchain, containing a version, previous hash, difficulty level, timestamp, Merkle root, nonce, the number of transactions, and a list of transactions.

- **Blockchain Structure:** Manages the chain of blocks and includes methods for creating the blockchain, adding blocks with transactions, and calculating the Merkle root.

- **Transaction Structure:** Represents a transaction with sender, receiver, and amount.

## How to Use

1. **Enter Difficulty:**
   - When running the program, the user is prompted to enter the mining difficulty for the blockchain.

2. **Generate Transactions:**
   - Transactions are generated with random sender, receiver, and amount values.

3. **Create Blockchain:**
   - The blockchain is created with the specified difficulty.

4. **Mine Blocks:**
   - Blocks are mined with transactions, and the Merkle root is calculated.

5. **Output Merkle Root:**
   - The final Merkle root of the blockchain is displayed.

6. **Note:**
   - The code includes commented-out sections that provide detailed information about each block. Uncomment and use these sections for more detailed block information.

## Example Output

```go
Enter Difficulty: 3
Merkle Root: 14bb1ce74a0f120e5a1a3e5ff8a5370323bae66eaab349e1c15e9cf442963fb6
```

## Dependencies

This project relies on the standard Go libraries and does not require any external dependencies.

## Acknowledgments

This blockchain implementation serves as an introductory example of building a blockchain in Go. Developers are encouraged to explore, modify, and extend the codebase to enhance their understanding of blockchain concepts in a practical setting.

Feel free to experiment, customize, and delve into the code to gain insights into the inner workings of a basic blockchain.
