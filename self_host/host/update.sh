#!/usr/bin/env bash

# Function to check if Docker Compose is installed
check_docker_compose() {
  if ! command -v docker compose &> /dev/null; then
    echo "Docker Compose could not be found. Please install Docker Compose."
    exit 1
  fi
}

# Function to stop containers
stop_containers() {
  echo "Stopping containers..."
  docker compose down
}

# Function to pull and start containers
start_containers() {
  echo "Pulling latest images..."
  docker compose pull
  
  echo "Removing old images..."
  docker image prune -f
  
  echo "Starting containers..."
  docker compose up -d
}

# Function to view logs
view_logs() {
  echo "Fetching logs..."
  docker compose logs
}

# Main script execution
check_docker_compose  # Check for Docker Compose before running any commands

# Parse command-line arguments
case "$1" in
  up)
    start_containers
    ;;
  down)
    stop_containers
    ;;
  logs)
    view_logs
    ;;
  *)
    echo "Usage: $0 {up|down|logs}"
    exit 1
    ;;
esac
