#!/bin/bash

# Create a backlog:

curl -X POST "http://localhost:3001/api/v1/Todo/AddItem" -H "accept: */*" -H "Content-Type: application/json" -d "{\"description\":\"Change lightbulb\",\"dueDate\":\"2018-04-20\",\"todoId\":\"aa152e0c138942a599ba0f2f84541f4e\"}"
curl -X POST "http://localhost:3001/api/v1/Todo/AddItem" -H "accept: */*" -H "Content-Type: application/json" -d "{\"description\":\"Repair faucet\",\"dueDate\":\"2018-04-22\",\"todoId\":\"bb152e0c138942a599ba0f2f84541f4e\"}"
curl -X POST "http://localhost:3001/api/v1/Todo/AddItem" -H "accept: */*" -H "Content-Type: application/json" -d "{\"description\":\"Paint hallway\",\"dueDate\":\"2018-04-23\",\"todoId\":\"cc152e0c138942a599ba0f2f84541f4e\"}"

