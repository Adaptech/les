# LES

## _"Let's Event Source Together"_

* Validates the design of an event-based system specified in Event Markdown or Event Markup Language.
* Generates an API from Event Markdown or Event Markup

![LESTER Pipeline](https://github.com/Adaptech/letseventsource/blob/master/LESTER-stack-diagram.png)

## Installation

### Prerequisites

- [NodeJS 8.11.1 LTS](https://nodejs.org/en/) or better
- [docker-compose](https://docs.docker.com/compose/install/)

### Linux (Ubuntu 16.04 x86_64)

Install the 'les' validation tool:

```sudo curl -L https://github.com/Adaptech/letseventsource/blob/master/releases/les/0.10.0/les-Linux-x86_64?raw=true -o /usr/local/bin/les && sudo chmod +x /usr/local/bin/les```

Install 'les-node':

```sudo curl -L https://github.com/Adaptech/letseventsource/blob/master/releases/les-node/0.10.0/les-node-Linux-x86_64?raw=true -o /usr/local/bin/les-node && sudo chmod +x /usr/local/bin/les-node```

### Windows (x86_84) binaries

Install the 'les' validation tool:

```curl -L https://github.com/Adaptech/letseventsource/blob/master/releases/les/0.10.0/les-windows-x86_64.exe?raw=true -o les.exe```

Install 'les-node':

```curl -L https://github.com/Adaptech/letseventsource/blob/master/releases/les-node/0.10.0/les-node-windows-x86_64.exe?raw=true -o les-node.exe```

### Max OSX (x86_64) binaries

Install the 'les' validation tool:

```sudo curl -L https://github.com/Adaptech/letseventsource/blob/master/releases/les/0.10.0/les-darwin-x86_64?raw=true -o /usr/local/bin/les && sudo chmod +x /usr/local/bin/les```

Install 'les-node':

```sudo curl -L https://github.com/Adaptech/letseventsource/blob/master/releases/les-node/0.10.0/les-node-darwin-x86_64?raw=true -o /usr/local/bin/les-node && sudo chmod +x /usr/local/bin/les-node```

## Getting Started

**Step 1:**

```bash
cat <<EOT >> Eventmarkdown.emd
# Hello World
Say Hello World->
HelloWorld Said
EOT
```

**Step 2:**

```bash
les convert && les-node -b && cd api && npm install && docker-compose up -d
```

**Step 3:**

There is no step 3.

* Swagger/OpenAPI docs for the new API: http://localhost:3001/api-docs
* Source Code: ./api
* API URI: http://localhost:3001/api/v1
* Eventstore DB: http://localhost:2113 (username 'admin', password 'changeit')

**Examples: https://github.com/Adaptech/les/src/master/samples**

## IDE Integrations & Tools

* Event Markdown [vscode extension](https://github.com/markgukov/vscode-event-markdown)


## Known Issues

* Casing shouldn't matter in property/parameter names (other than the "Id" convention at the end)

```
Receive Product-> // productId, description
InventoryItem Stocked // inventoryitemId, sku, description, purchasePrice, quantityAvailable
InventoryitemLookup* // inventoryitemId, productId, description
```
with "inventoryitemId" works. 

```
Receive Product-> // productId, description
InventoryItem Stocked // inventoryItemId, sku, description, purchasePrice, quantityAvailable
InventoryitemLookup* // inventoryItemId, productId, description
```

* Race condition when doing ```cd api && npm install && docker-compose up -d```: API doesn't start because Eventstore isn't up yet. (Workaround: ```docker-compose restart api```)

