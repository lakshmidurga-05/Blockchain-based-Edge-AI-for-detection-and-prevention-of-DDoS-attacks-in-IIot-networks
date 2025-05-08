# Blockchain-based-Edge-AI-for-detection-and-prevention-of-DDoS-attacks-in-IIot-networks
A Blockchain-enabled Edge AI System for Real-time DDoS Detection in IIoT.

## Overview

Fabric-SmartDefender is designed to enhance the security of IoT networks by:
- Managing IoT devices using Hyperledger Fabric.
- Detecting and mitigating DDoS attacks in real-time using machine learning.
- Use 20 devices for ML algorithm predictions for testing the HyperLedger fabric network and chaincode.

## Features

- **Blockchain Integration**: Uses Hyperledger Fabric to securely manage IoT device.
- **Machine Learning**: Employs an LSTM model to predict DDoS attacks based on network traffic data.

## Prerequisites

Before setting up the project, ensure you have the following installed:

1. **WSL**: Enable WSL in Windows

- Enable WSL and Virtual Machine Platform.

```bash
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```

Run the above in PowerShell as Administrator.

- Restart Your Computer
- Install a Linux Distribution
    Open Microsoft Store, search for Ubuntu. Install the latest version of Ubuntu.
- Set WSL2 as Default Version

```bash
wsl --set-default-version 2
```

Verify Installation

```bash
wsl --list --verbose
```

2. **Docker**: Required for running the Hyperledger Fabric test network.  
   [Install Docker](https://docs.docker.com/get-docker/)

3. **Go**: Required for developing and deploying chaincode.
   [Install Go](https://golang.org/doc/install)

Install on Ubuntu.
```bash
sudo apt update
sudo apt install golang-go -y
```

4. **JQ**: Install JQ on ubuntu/Linux
   [Install Go](https://golang.org/doc/install)

Install on Ubuntu.
```bash
sudo apt install jq
```

5. **NODE JS**: Required for developing and deploying Hyperledger caliper.
   [Install Go](https://golang.org/doc/install)

Install on Ubuntu.
```bash
curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash -

sudo apt install -y nodejs

node -v

npm -v
```


## Architecture

1. Download the file `diagrams/architecture.drawio` on to your machine.
2. You can edit the architecture diagram using the URL (Draw.io)[https://app.diagrams.net/] by uploading the file.

(architecture)[diagrams/architecture.drawio]

## Colab Execution

Step-by-Step: Create Kaggle API Key

1. Log into Kaggle: Go to https://www.kaggle.com and log in to your account.
2. Go to your account settings: Click on your profile icon (top right corner) → Select "Account".

    Or go directly here: https://www.kaggle.com/account

3. Scroll to the "API" section:
4. Find the "API" section near the bottom.
5. Click the "Create New API Token" button.
6. Download the API key:
    A file named kaggle.json will be downloaded.

This contains your username and API key.

Upload this kaggle.json file to download the datasets directly into your Google Colab environment.


## Setup Instructions

### HyperLedger Fabric network setup and Blockchain deployment

Follow the steps mentioned in (Fabric Setup)[fabric_setup.md]


### IIOT Device structure

| Name | Data Type | Description |
|------|-----------|-------------|
| IP | string | IP address of the device. |
| Name | string | Name of the device. |
| DeviceType | string | Type of the device like Sensor, Remote Controller. |

### Blockchain Functions

The full chaincode definition is available in the (chaincode)[chaincode/mitigation.go].

1. **createDeviceKey**: This function will generate the composite key for the ledger for performance.
2. **DeviceExists**: This function will be used to check whether device exist or not in the network. Input argument are `ip` and `state`.
3. **InitLedger**: This function will initialize to adds 20 predefined active IIOT devices to the ledger.
4. **RegisterDevice**: This function will be used to register a new active IIOT device in the network.
4. **QueryDevicesByState**: This function will be used to retrieve the devices at that point in time by state. The accepted values for state are `active`, `blocked` and `deleted`.
7. **BlockDevice**: This function will be used to block the malicious IIOT device in the network. This will block only devices which are in `active` state. Input argument is `ip`.
6. **DeleteDevice**: This function will be used to delete the malicious IIOT device from the network. This will delete only devices which are in `blocked` state. Input argument is `ip`.


### Datasets

The datasets which are used for training the ML model `LSTM` details are available in the (Datasets)[datasets.md].

### LSTM Model Training and Accuracy

The LSTM model code is available in the below files.

- (Edge IIoT Dataset)[models/edge_iiot_lstm.ipynb]
- (WUSTL 2021 Dataset)[models/wustl_lstm.ipynb]
- (IIoTID Dataset)[models/iiotid_lstm.ipynb]

### Images

- HyperLedger Fabric network setup and blockchain deployment images are available in `images/blockchain`.
- Caliper performance test result graphs are available in `images/caliper`. The raw HTML file is available in `metrics/report.html`.
- Model training details are available in `images/blockchain`.


### HyperLedger Caliper

- The test network definition file is available in the `caliper/networks/test-network.yaml`.
- The benchmarking test cases for the chaincode available in the location `caliper/benchmarks`.


## Usage

Follow below steps when you are setting up the blockchain for the first time without any softwares.

1. Install the pre-requisites as per the steps specified in `Prerequisites`.
2. Setup the HyperLedger Fabric network and chaincode deployment using the steps specified in `fabric_setup.md`.
3. Follow the steps specifed in the  `caliper_setup.md` to setup and execute the HyperLedger caliper.


## License

This project is licensed under the GNU General Public License v3.0. See the `LICENSE` file for details.
