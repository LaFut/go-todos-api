#!/bin/bash

log() {
  echo -e "${NAMI_DEBUG:+${CYAN}${MODULE} ${MAGENTA}$(date "+%T.%2N ")}${RESET}${@}" >&2
}

log "Waiting for Postgres..."
/root/wait-for-it.sh db:5432 --timeout=180 -- echo "PostgreSQL started"


log "Start server"

todos-rest
