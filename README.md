# INE5458-Blockchain MiniFabric Product Description Project

This project demonstrates a simple blockchain application using Hyperledger Fabric's MiniFabric tool. The goal of the application is to simulate a product description insertion system on a blockchain.

## Prerequisites

Before running the project, ensure that Docker is installed on your machine. You can follow the instructions for installing Docker on Ubuntu [here](https://docs.docker.com/engine/install/).

## Steps to Run the Project Manually

- Clone the repository:

```git clone git@github.com:LucassCPS/INE5458-blockchain.git```

```cd INE5458-blockchain```

```chmod +x setup.sh```

```./setup.sh```

- Test the chaincode by invoking functions. You can use MiniFab commands like:

```./minifab invoke -n app -p '"AddProduct", "arg"'```

```./minifab invoke -n app -p '"GetProduct", "arg"'```

```./minifab invoke -n app -p '"GetAllProducts", "arg"'```

- Clean up the MiniFabric network when done:

```./minifabric cleanup```
