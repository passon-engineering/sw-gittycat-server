#!/bin/bash

cd frontend
npm run build
cd ..

go build -o gittycat-server main.go   