#!/bin/bash

# Work on some TODOs:

curl -X POST "http://localhost:3001/api/v1/Todo/SetStatus" -H "accept: */*" -H "Content-Type: application/json" -d "{\"status\":\"doing\",\"todoId\":\"aa152e0c138942a599ba0f2f84541f4e\"}"
curl -X POST "http://localhost:3001/api/v1/Todo/SetStatus" -H "accept: */*" -H "Content-Type: application/json" -d "{\"status\":\"done\",\"todoId\":\"bb152e0c138942a599ba0f2f84541f4e\"}"
