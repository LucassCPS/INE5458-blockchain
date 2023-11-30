# INE5458-Blockchain MiniFabric Product Description Project

This project demonstrates a simple blockchain application using Hyperledger Fabric's MiniFabric tool. The goal of the application is to simulate a product technical description insertion system on a blockchain.

## Prerequisites

Before running the project, ensure that both [Docker](https://docs.docker.com/engine/install/) and [Go](https://go.dev/doc/install) are installed on your machine. 

## Steps to Run the Project Manually

- Get MiniFab's script:

```
mkdir -p ~/mywork && cd ~/mywork && curl -o minifab -sL https://tinyurl.com/yxa2q6yr && chmod +x minifab
```

- Get our chaincode:

```
git clone https://github.com/LucassCPS/INE5458-blockchain.git
```

- Get the MiniFab's network running:

```
./minifab up
```

- Copy and paste the "app" folder to "~/mywork/vars/chaincode" like the following command:

```
sudo cp -r [app_folder_path] ~/mywork/vars/chaincode/
```

- Install the chaincode:

```
./minifab ccup -n app -l go -d false -v 2.0
```

## How to use the application

- Test the chaincode by invoking the implemented functions. You can use MiniFab commands like:

```
./minifab invoke -p '"AddProduct", "manufacturerName", "modelName", "modelId", "anyExtraInformatin"'
```

```
./minifab invoke -p '"GetProduct", "modelName", "modelId"'
```

```
./minifab invoke -p '"GetAllProducts"'
```

```
./minifab invoke -p '"GetProductsByManufacturer", "manufacturerName"'
```

## Clean up the network

- Clean up the MiniFabric network when you are done:

```
./minifabric cleanup
```
