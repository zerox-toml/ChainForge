# ChainForge

## Introduction

This project is a self-built blockchain implementation designed to deepen understanding of distributed ledger technology and peer-to-peer networking. The blockchain is constructed from scratch, using sockets for node communication. While the current version operates on localhost, the architecture is designed to support live, distributed networks in the future. All solutions and designs are original, based on independent research and experimentation.

## Features

- **Custom Blockchain Core:** Implements fundamental blockchain concepts, including blocks, transactions, and basic consensus.
- **Peer-to-Peer Networking:** Uses sockets to simulate node communication and data propagation.
- **Pluggable Architecture:** Designed for future enhancements, such as cryptographic security and advanced networking protocols.
- **Simple Transaction Storage:** Currently uses a text file for transaction storage, with plans for database integration.

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/zerox-toml/ChainForge.git
   cd ChainForge
   ```
2. **Install dependencies:**
   Make sure you have Go installed (version 1.18 or higher recommended).
   You can download it from [golang.org](https://golang.org/dl/).
3. **Build and run the project:**
   ```bash
   go run main.go
   ```
   (If your entry point is a different file, replace `main.go` with the correct filename.)

## Roadmap

- [ ] Add cryptographic functions (SHA-256, ECDSA) for transaction integrity and digital signatures.
- [ ] Implement a robust communication protocol (e.g., Gossip protocol) to prevent message flooding.
- [ ] Stress test the system to identify and resolve edge cases and vulnerabilities.
- [ ] Integrate a database for persistent transaction and block storage.
- [ ] Develop a user-friendly interface for real-time monitoring and interaction.

## Contributing

Contributions are welcome! Please open issues or submit pull requests for suggestions, bug fixes, or enhancements.

## License

This project is open-source and available under the MIT License.
