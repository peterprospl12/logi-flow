#!/bin/bash

set -e
psql "$DATABASE_URL" -f scripts/1_init.sql