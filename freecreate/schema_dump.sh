#!/bin/bash

source .env

pg_dump -h $PG_HOST -U $PG_USER -d $PG_NAME --schema-only > ./migrations/schema.sql