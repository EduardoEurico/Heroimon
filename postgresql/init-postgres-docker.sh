#!/bin/bash

# Parar e remover o container pgAdmin, se ele existir
if [ $(docker ps -a -q -f name=my-Pgadmin-new) ]; then
  echo "Parando e removendo o container pgAdmin (my-Pgadmin-new)..."
  docker stop my-Pgadmin-new
  docker rm my-Pgadmin-new
fi

# Parar e remover o container Postgres, se ele existir
if [ $(docker ps -a -q -f name=myPostgres) ]; then
  echo "Parando e removendo o container Postgres (myPostgres)..."
  docker stop myPostgres
  docker rm myPostgres
fi

# Criar a rede do Docker, se ela não existir
docker network inspect my-network >/dev/null 2>&1 || {
  echo "Criando a rede do Docker (my-network)..."
  docker network create my-network
}

# Recriar o container do Postgres
echo "Criando o container do Postgres (myPostgres)..."
docker run -d --name myPostgres --network=my-network -p 5433:5432 -e POSTGRES_PASSWORD=postgres -v /home/kaynan/Documentos/desenvolvimento/go/Heroimon/postgresql/data:/var/lib/postgresql/data postgres

# Recriar o container do pgAdmin
echo "Criando o container do pgAdmin (my-Pgadmin-new)..."
docker run -d --name my-Pgadmin-new --network=my-network -p 15432:80 -e PGADMIN_DEFAULT_EMAIL=kaynan@gmail.com -e PGADMIN_DEFAULT_PASSWORD=postgres dpage/pgadmin4

echo "Containers criados com sucesso."

