# 🧾 Ethereum TodoList API

This is a beginner-friendly REST API built with **Go (Golang)** and the **Gin** framework to interact with an Ethereum smart contract (`TodolistV2`). The smart contract is deployed as a **Transparent Proxy** on the **Sepolia testnet**.

The API allows you to:

- ✅ Create tasks
- 📝 Update task content
- ❌ Delete tasks
- 🔍 Get task details by ID

The project uses:
- 🛠️ [go-ethereum](https://github.com/ethereum/go-ethereum) for blockchain interaction
- ⚙️ Contract bindings generated with `abigen`
- 🌐 RPC communication with Sepolia via Alchemy
- 🔐 Private key signing with `bind.TransactOpts`

---

## 📂 Project Structure

```
todolist-api/
├── cmd/server/ # Main entrypoint
├── config/ # Loads .env variables
├── internal/
│ ├── ethereum/ # Ethereum client and contract
│ └── api/ # API handlers and router
├── contracts/ # Go bindings from the smart contract ABI
├── .env # Private key, RPC URL, contract address
```
### 🔧 Setup Instructions

- First download GETH (go-ethereum) from https://geth.ethereum.org/downloads

```
cd folder-name
```
### 🚀 Running the Project

- Initialize Go Modules

```
go mod init github.com/user-name/folder-name
go mod tidy
```

### 🚀 Install dependencies
```
go get github.com/ethereum/go-ethereum
go get github.com/joho/godotenv
```

### 🚀 Generate Go bindings from ABI
```
abigen --abi ./Todolist.json --pkg todolistv2 --out ./todolistv2.go
```

### 🚀 Run server
```
go run cmd/server/main.go
```

### 🚀 For testing use postman