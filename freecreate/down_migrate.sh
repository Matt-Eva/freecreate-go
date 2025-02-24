#!/bin/bash

source .env

migrate -source file://./migrations -database postgres://$PG_HOST/$PG_NAME?sslmode=disable down 1