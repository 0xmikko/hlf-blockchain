# hlf-blockchain

## Установка и настройка Hyperledger Fabric

Перед установкой Hyperledger Fabric необходимо убедиться, что на платформе установлены все необходимые компоненты: 
- Git
- cURL
- Docker & Docker Compose
- Golang 1.15
— Node.js & npm
— Python 2.7

После подготовки платформы можно переходить к установке Hyperledger Fabric. 

#### Шаг 1. Загрузка образа

В продукте используется последняя версия Hyperledger Fabric (v2.2.1), поэтому для загрузки последней версии Hyperledger Fabric выполняем команду: 

```
curl -sSL http://bit.ly/2ysbOFE | bash -s
```

#### Шаг 2. Подъем тестовой сети

После загрузки образа Hyperledger Fabric запустим сеть с конфигурацией по умолчанию (2 peer-ноды и 1 ordered-нода):  

```
cd fabric-samples/test-network
./network.sh up
```

В случае успеха мы должны увидеть: 

```
Creating network "net_test" with the default driver
Creating volume "net_orderer.example.com" with default driver
Creating volume "net_peer0.org1.example.com" with default driver
Creating volume "net_peer0.org2.example.com" with default driver
Creating orderer.example.com    ... done
Creating peer0.org2.example.com ... done
Creating peer0.org1.example.com ... done
CONTAINER ID        IMAGE                               COMMAND             CREATED             STATUS                  PORTS                              NAMES
8d0c74b9d6af        hyperledger/fabric-orderer:latest   "orderer"           4 seconds ago       Up Less than a second   0.0.0.0:7050->7050/tcp             orderer.example.com
ea1cf82b5b99        hyperledger/fabric-peer:latest      "peer node start"   4 seconds ago       Up Less than a second   0.0.0.0:7051->7051/tcp             peer0.org1.example.com
cd8d9b23cb56        hyperledger/fabric-peer:latest      "peer node start"   4 seconds ago       Up 1 second             7051/tcp, 0.0.0.0:9051->9051/tcp   peer0.org2.example.com
```

"Погасить" сеть можно командой: 
```
./network.sh down
```

Далее можно переходить к реализации бизнес логики нашего проекта. 

## Установка chaincode

Каждый узел и пользователь, которые взаимодействуют с сетью, должны принадлежать к организации, которая является членом сети. Группу организаций, входящих в сеть Fabric, часто называют консорциумом. В нашей сети есть два члена консорциума (peer-ноды): Org1 и Org2. В сеть также входит одна организация-заказчик (ordered-нода), которая поддерживает сервис запросов в сети. 

Peer-ноды - фундаментальные компоненты любой сети Fabric. Они хранят реестр блокчейна и проверяют транзакции, прежде чем они будут зафиксированы в реестре, запускают смарт-контракты, содержащие бизнес-логику, которая используется для управления активами в реестре блокчейна.

В нашей сети каждая организация использует по одному партнеру: peer0.org1.example.com и peer0.org2.example.com.

После того, как мы настроили сеть, можем переходить к установке chaincode.

#### Шаг 1. Создание канала между Org1 и Org2

Каналы - это частный уровень связи между определенными участниками сети. Каналы могут использоваться только организациями, которые приглашены на канал, и невидимы для других участников сети. У каждого канала есть отдельный регистр блокчейна. Приглашенные организации «присоединяются» к своим коллегам к каналу, чтобы хранить реестр каналов и проверять транзакции на канале. Для создания канала выполняем команду: 
```
./network.sh createChannel
```

В случае успеха в консоли выведется сообщение: 
```
========= Channel successfully joined ===========
```

#### Шаг 2. Упаковка chaincode для установки

Перед установкой chaincode необходимо создать архив с кодом и зависимостями. Наш chaincode мы поместили в папку fabric-samples/asset-transfer-basic/chaincode-go, предварительно почистив все содержимое  в ней, поэтому выполняем след. действия: 
```
cd fabric-samples/asset-transfer-basic/chaincode-go
sudo GO111MODULE=on go mod vendor
cd ../../test-network
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
peer lifecycle chaincode package basic.tar.gz --path ../asset-transfer-basic/chaincode-go/ --lang golang --label basic_1.0
```

В конечном итоге на этом шаге мы должны получить файл basic.tar.gz в каталоге test-network. 

#### Шаг 3. Установка chaincode на peer-ноды 

Chaincode должен быть установлен на каждом узле, который подтвердит транзакцию. Поскольку мы собираемся настроить консенсус так, чтобы требовать подтверждения как от Org1, так и от Org2, нам нужно установить chaincode на peer-нодах, управляемых обеими организациями: peer0.org1.example.com и peer0.org2.example.com. 

Устанавливаем chaincode для Org1: 
```
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
peer lifecycle chaincode install basic.tar.gz
```
 
В случае успеха увидим: 
```
2020-12-13 10:09:57.534 CDT [cli.lifecycle.chaincode] submitInstallProposal -> INFO 001 Installed remotely: response:<status:200 payload:"\nJbasic_1.0:e2db7f693d4aa6156e652741d5606e9c5f0de9ebb88c5721cb8248c3aead8123\022\tbasic_1.0" >
2020-12-13 10:09:57.534 CDT [cli.lifecycle.chaincode] submitInstallProposal -> INFO 002 Chaincode code package identifier: basic_1.0:e2db7f693d4aa6156e652741d5606e9c5f0de9ebb88c5721cb8248c3aead8123
```

Теперь установим для Org2:
```
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
peer lifecycle chaincode install basic.tar.gz
```

#### Шаг 4. Подтверждение chaincode

После установки chaincode необходимо утвердить определение chaincode для вашей организации. Определение включает важные параметры управления chaincode, такие как имя, версия и политика поддержки цепного кода. Чтобы связать chaincode, установленный на peer-ноде, с утвержденным определением chaincode используется идентификатор пакета, который позволяет организации использовать chaincode для подтверждения транзакций. Идентификатор пакета chaincode можно получить с помощью команды:
```
peer lifecycle chaincode queryinstalled
```

Результат выполнения команды: 
```
Installed chaincodes on peer:
Package ID: basic_1.0:69de748301770f6ef64b42aa6bb6cb291df20aa39542c3ef94008615704007f3, Label: basic_1.0
```

В нашем случае идентификатор — basic_1.0:69de748301770f6ef64b42aa6bb6cb291df20aa39542c3ef94008615704007f3: 
```
export CC_PACKAGE_ID=basic_1.0:69de748301770f6ef64b42aa6bb6cb291df20aa39542c3ef94008615704007f3
```

Далее подтвердим определение chaincode для обеих организаций: 
```
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

#### Шаг 5. Передача chaincode в канал

После того, как достаточное количество организаций одобрило определение chaincode, одна организация может передать определение chaincode в канал. Если большинство участников канала одобрили определение, транзакция будет успешной, и параметры, согласованные в определении chaincode, будут реализованы в канале.

```
peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name basic --version 1.0 --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --output json
```

Команда создаст JSON, который отображает, одобрил ли участник канала параметры, указанные в команде checkcommitreadiness:

```
 {
            "Approvals": {
                    "Org1MSP": true,
                    "Org2MSP": true
            }
 }
```

Поскольку обе организации, являющиеся членами канала, утвердили одни и те же параметры, определение chaincode готово к передаче в канал:

```
peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
```

#### Шаг 6. Вызов chaincode

После передачи chaincode в канал мы можем обращаться к нему из клиентских приложений. В следующем вызове мы зададим начальный набор активов в реестре: 
```
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"InitLedger","Args":[]}'
```

## Установка приложений 

Для установки и запуска приложений (каталог app1) необходимо создать 2 католога в fabric-samples (go1 и go2), куда поместить файлы из app1:
```
mkdir go1
mkdir go2
```

Далее необходимо собрать приложение: 
```
go build -o ./build/main ./cmd/main.go
```

Сборку приложения нужно выполнить в каждой папке (go1 и go2). Для запуска приложения необходимо определить также файл .env:
```
DISCOVERY_AS_LOCALHOST=false
SALT=Nuugoal
ENV=development
PORT=8080
```

PORT в нашем случае для Org1 = 8080, для Org2 = 5000. 

Для запуска приложения выполняем команду:
```
cd build
./main
```

После запуска приложений можно протестировать вызов снаружи для каждой ноды: 
```
GET http://130.193.59.251:8080/api/receivables/
Accept: application/json

GET http://130.193.59.251:5000/api/receivables/
Accept: application/json
```

## Клиент

Для установки клиента (client1): 
```
npm install
npm start 
```

## Инфраструктура

В конечном итоге после правильного выполнения вышеописанных действий мы получаем следующую инфраструктуру:

- Сеть 1: 130.193.59.251 (Oil supplier network)
```
Org1: localhost:7051
Org2: localhost:9051
AppOrg1: 130.193.59.251:8080
AppOrg2: 130.193.59.251:5000
```
- Сеть 2: 130.193.58.229 (Factoring network)
```
Org1: localhost:7051
Org2: localhost:9051
AppOrg1: 130.193.58.229:8080
AppOrg2: 130.193.58.229:5000
```
- Фронтенд: http://178.154.235.148:3000

## Описание вызовов 

#### Получение списка записей 

```
GET http://130.193.59.251:8080/api/receivables/
Accept: application/json
```

#### Добавление записи

```
POST http://130.193.59.251:8080/api/receivables/
Content-Type: application/json

{
  "issuer": "Gazprom Test",
  "payer": "K5",
  "amount": 500
}
```

#### Передача записи из одной сети в другую
```
POST http://130.193.58.229:8080/api/receivables/sync/

{
  "id": "e8d4e61f-0ea8-48b7-a5b5-0cd4964f7b9d"
}
```

