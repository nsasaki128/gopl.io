#!/bin/bash
go run server/main.go &
go run main.go http://localhost:8000