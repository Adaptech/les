# LES

## _"Let's Event Source Together"_

* Validates the design of an event-based system specified in Event Markdown or Event Markup Language.
* Generates an API from Event Markdown or Event Markup

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

## Event Markup Language (EML) & Event Markdown (EMD) Versions

### Current (0.1-alpha)


**Known Issues** 

* TODO: Casing shouldn't matter in property/parameter names (other than the "Id" convention at the end)

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

with "inventoryItemId" does not; gives misleading validation errors.

* TODO: 'command -> someParameter' should fail emd validation - 'command -> // someParameter is the correct way. Alternatively: Should the EMD DSL be changed so that this (... and 'This Happened* firstProperty, secondProperty' for events... ) is valid instead of requiring the '//'? 

* TODO: "Emails* // emailId, fromUserId," leads to a validation error (because of the comma at the end), but should be OK.
* Race condition when doing ```cd api && npm install && docker-compose up -d```: API doesn't start because Eventstore isn't up yet. (Workaround: ```docker-compose restart api```)
* TODO: Find a way of finding SubscribesTo events for read models which exclusively have aggregateId fields.
* TODO: EMD: A command followed immediately by another command without an event in between should fail validation.

* TODO: Introduce numeric, date and boolean types.
* TODO: vscode EMD code colorizer

* TODO: Prompt before overwriting 'les convert' files.
* TODO: Prompt before overwriting 'les-node -b' .api dirctory.
* TODO: Introduce '-y' to always say yes when prompted about overwriting files.
* TODO: Do a non-zero return code from the les command line client when there are validation errors.


### Next (0.2-alpha)


* TODO: Preconditions for commands, e.g. ```Has TimesheetCreated``` or ```Not TimesheetDeleted``` or ```LastEventIs Not TimesheetSubmitted```
* TODO: Selecting individual event properties in Readmodel SubscribesTo
* TODO: Spreadsheet-like functionality for Readmodels, e.g. ```@SUM(TimesheetHoursLogged.hours).totalHours```

```yaml

Solution: Timesheets & Billing
EmlVersion: 0.2-alpha
Contexts:
- Name: Timesheets & Billing
  Streams:
  - Stream: User
    Commands:
    - Command:
        ID: RegisterUser
        Name: Register User
        Parameters:
        - Name: email
          Type: string
          Rules: []
        - Name: password
          Type: string
          Rules: []
        - Name: userId
          Type: string
          Rules:
          - IsRequired
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        ID: UserRegistered
        Name: User Registered
        Properties:
        - Name: email
          Type: string
          IsHashed: false
        - Name: password
          Type: string
          IsHashed: true
        - Name: userId
          Type: string
          IsHashed: false
        Type: ""
  - Stream: Timesheet
    Commands:
    - Command:
        ID: CreateTimesheet
        Name: Create Timesheet
        Parameters:
        - Name: userId
          Type: string
          Rules:
          - MustExistIn UserLookup
        - Name: description
          Type: string
          Rules: []
        - Name: timesheetId
          Type: string
          Rules:
          - IsRequired
        Preconditions:
        - Not TimesheetCreated
        - LastEventIs Not TimesheetSubmitted
        Postconditions:
        - TimesheetCreated
    - Command:
        ID: SubmitTimesheet
        Name: Submit Timesheet
        Parameters:
        - Name: submissionDate
          Type: string
          Rules: []
        - Name: userId
          Type: string
          Rules:
          - MustExistIn UserLookup
        - Name: timesheetId
          Type: string
          Rules:
          - IsRequired
        Preconditioins:
        - Has TimesheetCreated
        - Not Has TimesheetSubmitted
        Postconditions:
        - TimesheetSubmitted
    Events:
    - Event:
        ID: TimesheetCreated
        Name: Timesheet Created
        Properties:
        - Name: userId
          Type: string
          IsHashed: false
        - Name: description
          Type: string
          IsHashed: false
        - Name: timesheetId
          Type: string
          IsHashed: false
        Type: ""
    - Event:
        ID: TimesheetSubmitted
        Name: Timesheet Submitted
        Properties:
        - Name: timesheetId
          Type: string
          IsHashed: false
        - Name: submissionDate
          Type: string
          IsHashed: false
        - Name: userId
          Type: string
          IsHashed: false
        Type: ""
  - Stream: TimesheetHours
    Commands:
    - Command:
        ID: LogHours
        Name: Log Hours
        Parameters:
        - Name: timesheethoursId
          Type: string
          Rules:
          - IsRequired
        - Name: timesheetId
          Type: string
          Rules:
          - MustExistIn TimesheetLookup
        - Name: date
          Type: string
          Rules: []
        - Name: hours
          Type: string
          Rules: []
        Postconditions:
        - TimesheetHoursLogged
    Events:
    - Event:
        ID: TimesheetHoursLogged
        Name: TimesheetHours Logged
        Properties:
        - Name: timesheethoursId
          Type: string
          IsHashed: false
        - Name: timesheetId
          Type: string
          IsHashed: false
        - Name: date
          Type: string
          IsHashed: false
        - Name: hours
          Type: string
          IsHashed: false
        Type: ""
  Readmodels:
  - Readmodel:
      ID: UserLookup
      Name: UserLookup
      Key: userId
      Columns:
      - UserRegistered.*
  - Readmodel:
      ID: TimesheetLookup
      Name: TimesheetLookup
      Key: timesheetId
      SubscribesTo:
      - TimesheetCreated.*
      - TimesheetSubmitted.submissionDate
      - @SUM(TimesheetHoursLogged.hours).totalHours
      - @MIN(TimesheetHoursLogged.date).fromDate
      - @MAX(TimesheetHoursLogged.date).toDate
      - @IF(@ISBLANK(TimesheetSubmitted.submissionDate),"","submitted on {{TimesheetSubmitted.submissionDate}}")
  - Readmodel:
      ID: TimesheetHoursLookup
      Name: TimesheetHoursLookup
      Key: timesheethoursId
      Columns:
      - TimesheetHoursLogged.*
Errors: []

```
