<div style="display: flex; align-items: center;">
  <h1 style="margin-right: 10px; font-size: 32px; line-height: 40px; margin: 0; ">Toolchain &nbsp;</h1>
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" height="70" alt="Go logo"  style="vertical-align: middle; margin-top: -5px;"/>
</div>


**Toolchain Go** is a powerful collection of APIs for basic operations that every developer needs in their day-to-day life. Whether you're working on a web app or API, this toolkit simplifies tasks like encoding/decoding, cryptographic operations, data compression, and file handling, all served through a simple local server. No more waiting for slow web pages to load or dealing with annoying ads—**Toolchain Go** brings all the functionality you need directly to your development environment.

## Key Features

- **Base64 Encoding & Decoding**: Convert data to/from Base64 format.
- **AES Encryption & Decryption**: Securely encrypt and decrypt data using AES.
- **Key Generation**: Generate cryptographic keys (AES, SHA256, etc.).
- **Data Compression**: Compress and decompress data to reduce size.
- **File Operations**: Efficient file reading, writing, and streaming.
- **Hashing**: Generate hashes (SHA256, MD5, etc.).
- **Data Conversion**: Convert between different formats (e.g., string to byte array).

## Installation

### Requirements

- Go 1.22+ (or newer)

### Install via `go get`

To use **Toolchain Go**, install it with the following command:

```bash
go get github.com/yourusername/toolchain-go
```

### Running locally
**Clone the Repository**
```bash
git clone https://github.com/yourusername/toolchain-go.git
```
**Change the working directory to toolchain-go**
```bash
cd toolchain-go
```
**Download parent dependencies**
```bash
go mod tidy
```
**Run the application**
```bash
go run main.go
```
