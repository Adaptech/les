#!/bin/bash

# Pull items from backlog into TODO:

curl -X POST "http://localhost:3001/api/v1/Todo/SetStatus" -H "accept: */*" -H "Content-Type: application/json" -d "{\"status\":\"todo\",\"todoId\":\"aa152e0c138942a599ba0f2f84541f4e\"}"
curl -X POST "http://localhost:3001/api/v1/Todo/SetStatus" -H "accept: */*" -H "Content-Type: application/json" -d "{\"status\":\"todo\",\"todoId\":\"bb152e0c138942a599ba0f2f84541f4e\"}"
curl -X POST "http://localhost:3001/api/v1/Todo/SetStatus" -H "accept: */*" -H "Content-Type: application/json" -d "{\"status\":\"todo\",\"todoId\":\"cc152e0c138942a599ba0f2f84541f4e\"}"
