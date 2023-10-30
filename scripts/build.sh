#!/usr/bin/env bash

function buildlocal()
{
    mkdir -p bin
    cd arkweb
    rm -f arkweb
    go build -o ../bin
    cd ..
}


function buildcloud()
{
    mkdir -p bin
    cd arkweb
    rm -f arkweb
    GOOS=linux GOARCH=amd64 go build -o ../bin
    cd ..
}

function runLocal()
{
    buildlocal
    func start
}

function publish() 
{
    buildcloud
    func azure functionapp publish katasecweb
}