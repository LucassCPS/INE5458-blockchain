#!/bin/bash

echo "Obtendo o executavel do MiniFabric"
mkdir -p ~/mywork && cd ~/mywork && curl -o minifab -sL https://tinyurl.com/yxa2q6yr && chmod +x minifab

echo "Iniciando o MiniFabric"
./minifab up

# Solicita ao usuário o diretório da pasta source
read -p "Digite o caminho completo da pasta source: " caminho_pasta_origem

# Verifica se o diretório existe
if [ ! -d "$caminho_pasta_origem" ]; then
    echo "Erro: O diretório '$caminho_pasta_origem' não existe."
    exit 1
fi

# Caminho de destino no MiniFabric
caminho_pasta_destino="$HOME/mywork/vars/chaincode/"

echo "Copiando a pasta 'app' e seus subdiretórios para o MiniFabric"
sudo cp -r "$caminho_pasta_origem" "$caminho_pasta_destino"

echo "Executando o comando para instalar e iniciar a chaincode no MiniFabric"
./minifab ccup -n app -l go -d false -v 2.0

