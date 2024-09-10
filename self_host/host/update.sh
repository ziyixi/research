#!/usr/bin/env bash

# Function to check if Docker Compose is installed
check_docker_compose() {
  if ! command -v docker-compose &> /dev/null; then
    echo "Docker Compose could not be found. Please install Docker Compose."
    exit 1
  fi
}

# Function to perform the update steps
update_containers() {
  echo "start update.sh"

  echo "Stopping containers..."
  docker compose down
  
  echo "Pulling latest images..."
  docker compose pull
  
  echo "Removing old images..."
  docker image prune -f
  
  echo "Starting containers..."
  docker compose up -d
  
  echo "end update.sh"
}

# Main script execution
check_docker_compose  # Check for Docker Compose before running any commands
update_containers  # Call the function that contains the main update steps
