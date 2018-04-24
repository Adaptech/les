# LES

## "Let's Event Source"

**Event sourcing and CQRS/ES based "microservices" are increasingly seen as a nice way to build cohesive, loosely coupled systems with good transactional integrity. Most of us aren't used to thinking event centric and designing software that way, so although the resulting systems tend to be much simpler and easier to understand than traditional (e.g.) object oriented implementations, there is a bit of a learning curve.**

LES attempts to address this in three ways:

1. **Fast microservice prototyping:** Go directly from an event storming to a working event sourced API.

2. **"Architect in a box":** ```les validate``` assesses whether a prototype will result in a "good" event sourced microservice - cohesive, loosely-coupled, transactionally consistent. Then ```les-node -b``` builds a deployment-ready NodeJS API with plenty of guide fences and best practices in place as developers go forward customizing it. If you have your own coding standards or don't like NodeJS, implement your own in a language of your choice.

3. **"Citizen IT Developer".** One of the goals of the LES project is to enable "business coders", "power users" and entrepreneurs with little technical knowledge to build highly scalable event sourced microservices from scratch, basically "I've made this API for my startup - could you build me an app for that?"

LES is currently in alpha. We have started using 1. and 2. in Real Life projects. 3. (Citizen IT Developer) is experimental, with quite a few features missing.

See also: [LES FAQ](https://github.com/Adaptech/letseventsource)

![LESTER Pipeline](https://github.com/Adaptech/letseventsource/blob/master/LESTER-stack-diagram.png)

## Getting Started

### Prerequisites

* [NodeJS 8.11.1 LTS](https://nodejs.org/en/) or better
* [docker-compose](https://docs.docker.com/compose/install/)

### Installation

**Latest version from source:**

```bash
git clone https://github.com/Adaptech/les.git
make install
```

... or ... 

[Instructions for Linux, Windows Mac & Docker](INSTALL.md)


### Hello World

**Step 1:**

Create an [Event Markdown](https://webeventstorming.com) file. Event Markdown (EMD) is a simple language used to describe an [event storming](https://ziobrando.blogspot.ca/2013/11/introducing-event-storming.html):

```bash
# TODO List
Add Item -> // description, dueDate
Todo Added // description, dueDate
TODO List* // todoId, description, dueDate
```
Save it to ```Eventstorming.emd```. 

**Step 2:**

```bash
les convert && les-node -b && cd api && npm install && docker-compose up -d --force-recreate
```

Or using Docker:
```bash
docker run -v $(pwd):/les les convert && docker run -v $(pwd):/les les-node -b && cd api && npm install && docker-compose up -d
```

(If you doing this in Linux and encounter "permission denied" errors, your USER or GROUP ID need to be specified.
 Say your USER ID is 1003, then add `--user 1003` after each `docker run` in the above command.)

**Step 3:**

There is no step 3.

* Add some items to the TODO list: http://localhost:3001/api-docs (Swagger/OpenAPI)
* View the items: http://localhost:3001/api/v1/r/TODOList
* Look at the "TodoAdded" events in the Eventstore DB: http://localhost:2113 (username 'admin', password 'changeit')
* Check out the source code for the "TODO List" system: ```./api```

## What next ...

* A collection of Event Markdown (EMD) examples: https://github.com/Adaptech/les/src/master/samples**

* Learn Event Storming: http://eventstorming.com

* Learn Event Markdown: https://webeventstorming.com

* EMD Cheat Sheet: https://github.com/Adaptech/letseventsource/raw/master/EMD-Cheatsheet-0.10.0-alpha-alpha.pdf

* https://gitter.im/Adaptech/les 

## IDE Integrations & Tools

* Event Markdown [vscode extension](https://github.com/markgukov/vscode-event-markdown)

## Known UX Impacting Issues

The issues below have been known to mystify EMD users:

#### "DromedaryCase": myaggregateId GOOD, myAggregateId BAD

https://github.com/Adaptech/les/issues/9

#### Sporadic Race condition when doing ```cd api && npm install && docker-compose up -d```

API doesn't start because Eventstore isn't up yet. (Workaround: ```docker-compose restart api```)

https://github.com/Adaptech/les/issues/11

#### Need to have at least one read model parameter which is not an aggregate ID

https://github.com/Adaptech/les/issues/10

## Running The Tests

```make test-all```
