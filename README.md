# CosmeticStealer

Welcome to **CosmeticStealer**! This tool is designed for educational purposes to log and analyze cosmetic data from Minecraft players' skins. It uses `gophertunnel` to intercept and log skin packet data for each player in a Minecraft server.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Overview

**CosmeticStealer** uses the [gophertunnel](https://github.com/Sandertv/gophertunnel) library to intercept Minecraft skin packets. This allows you to log and analyze skin data from players on any Minecraft server you connect to. The tool captures skin data by logging into Xbox Live and then connecting to the desired Minecraft server.



![img](https://cdn.discordapp.com/attachments/623673869177192462/1271348045442711572/75175146_geometry.png?ex=66b702a2&is=66b5b122&hm=5e8f9c6adb45c4b874edda030804d93ac97c937f9cae4aa4e4ada840d2fd8a13&)



## Prerequisites

Before you start, ensure you have the following:

- Go (Golang) installed on your system.
- [gophertunnel](https://github.com/Sandertv/gophertunnel) library installed.
- A valid Xbox Live account for authentication.
- Access to the Minecraft server you wish to log skin data from.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/ZyloxDeveloper/CosmeticStealer.git
   cd CosmeticStealer
   ```

2. **Install dependencies:**

   Make sure you have Go installed and properly configured. Then run:

   ```bash
   go mod tidy
   ```

3. **Build the project:**

   Compile the project using Go:

   ```bash
   go build -o CosmeticStealer main.go
   ```

## Usage

1. **Run the program:**

   ```bash
   ./CosmeticStealer
   ```

2. **Sign in with Xbox Live:**

   Follow the prompts to authenticate using your Xbox Live account.

3. **Enter the server details:**

   Provide the IP address and port of the Minecraft server you wish to log skin data from.

4. **Start logging:**

   Once connected, the tool will begin intercepting and logging skin data from players on the server. The logs will be saved in a format that allows you to review and analyze the captured data.

## Contributing

We welcome contributions to **CosmeticStealer**! If you have suggestions, improvements, or fixes, please submit a pull request or open an issue on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

**Disclaimer:** This tool is intended for educational purposes only. Unauthorized interception of data and accessing private information without consent is illegal and unethical. Ensure you have proper authorization before using this tool on any Minecraft server.
