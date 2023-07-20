# Architecture Documentation

This document describes the high-level architecture and workflow of this project.

## Table of Content

1. [Introduction](#introduction)
2. [States](#states)
    1. [Startup Sequence](#startup-sequence)
    2. [Runtime Actions](#runtime-actions)


## Introduction
TBD

## States
Description of tasks depending on states

### Startup Sequence
1. Initialize Application
  - Generate `global app object` with placeholders and server path for further objects and information
  - Create and start logger instance
  - Get IP of machine where app is running on
  - Display 'PUBLIC' SSH key to user / logger (useful for git repository access configuration)

2. Initialize Webhooks
  - Read and parse all available *.yaml webhooks from ./webhooks/. These are the webhooks the server will be listening for
  - Store all available webhooks in the `global app object`

3. Initialize Repositories
  - Range over all available webhooks from the `global app object` and clone the repository of each webhook in the ./repositories/ folder

4. Initialize Processor
  - Create a queue buffer (channel) for incomming actions to be processed 

5. Initialize Webserver
  - Create handler to read static files from ./frontend/dist/ and listen to : '/'
  - Create handler to listen to: '/webhooks' 
  - Create handler to listen to: '/webhooks/refresh'
  - Create handler to listen to: '/webhooks/{repo_name}/{action}'

### Runtime Actions
--> TRIGGER -> /*
  - Read requested file from given path
  - Send file content to browser 

--> TRIGGER -> /webhooks
  - Read all available webhooks from the `global app object`
  - Pass result as JSON to browser
  - If no webhooks available: Browser will display information 
  - Create table and for each webhook: Display 'repository name' and show as 'active'

--> TRIGGER -> /webhooks/refresh
  - Read and parse all available *.yaml webhooks from ./webhooks/. These are the webhooks the server will be listening for
  - Store all available webhooks in the `global app object`
  - Read all available webhooks from the `global app object`
  - Pass result as JSON to browser
  - If no webhooks available: Browser will display information 
  - Create table and for each webhook: Display 'repository name' and show as 'active' if so

--> TRIGGER -> /webhooks/{repo_name}/toggle_active
  - Read specific webhook via repo_name and toggle the active status
  - If true: It will process incomming webhooks
  - If false: It will not process incomming webhooks
  - Write the active state to the webhook