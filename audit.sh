#!/bin/bash

# ASCII Art Web Docker Audit Script
# Tests Docker implementation for ascii-art-web project

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

check_success() {
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ Success${NC}"
    else
        echo -e "${RED}❌ Failed${NC}"
        exit 1
    fi
}

echo -e "${BLUE}=== ASCII Art Web Docker Audit ===${NC}"

# Clean previous resources
echo -e "${YELLOW}Step 1: Cleaning previous resources${NC}"
docker rm -f ascii-art-container 2>/dev/null
docker rmi ascii-art-web 2>/dev/null
check_success

# Build Docker image
echo -e "${YELLOW}Step 2: Building Docker image${NC}"
docker build -t ascii-art-web .
check_success

# Run container
echo -e "${YELLOW}Step 3: Running container${NC}"
docker run -p 8080:8080 --detach --name ascii-art-container ascii-art-web
check_success

# Test web application
echo -e "${YELLOW}Step 4: Testing web application${NC}"
sleep 3  # Wait for container to start
curl -s -I http://localhost:8080 | grep "200 OK"
check_success

# Test ASCII art generation
echo -e "${YELLOW}Step 5: Testing ASCII art generation${NC}"
response=$(curl -s -X POST -d 'text=Hello&banner=standard' http://localhost:8080/ascii-art)
if [[ $response == *"Hello"* ]] && [[ $response == *"_"* ]]; then
    echo -e "${GREEN}✅ ASCII art generation working${NC}"
else
    echo -e "${RED}❌ ASCII art generation failed${NC}"
    exit 1
fi

# Test export functionality
echo -e "${YELLOW}Step 6: Testing export functionality${NC}"
curl -s -X POST -d 'text=Test&banner=standard&export=true' http://localhost:8080/ascii-art -o test-export.txt
if [ -f test-export.txt ] && [ -s test-export.txt ]; then
    echo -e "${GREEN}✅ Export functionality working${NC}"
    rm -f test-export.txt
else
    echo -e "${RED}❌ Export functionality failed${NC}"
    exit 1
fi

# Verify file structure
echo -e "${YELLOW}Step 7: Checking file structure${NC}"
docker exec ascii-art-container ls /app/ascii-art-web /app/templates/Page.html /app/banners/standard.txt >/dev/null
check_success

echo -e "\n${GREEN}=== Audit Summary ===${NC}"
echo -e "${GREEN}✓ Docker image built successfully${NC}"
echo -e "${GREEN}✓ Container running on port 8080${NC}"
echo -e "${GREEN}✓ Web application responding${NC}"
echo -e "${GREEN}✓ ASCII art generation working${NC}"
echo -e "${GREEN}✓ Export functionality working${NC}"
echo -e "${GREEN}✓ File structure correct${NC}"

echo -e "\n${YELLOW}Manual Testing:${NC}"
echo "Open http://localhost:8080 in browser and test all banner styles"

echo -e "\n${YELLOW}Cleanup:${NC}"
echo "docker stop ascii-art-container && docker rm ascii-art-container && docker rmi ascii-art-web"

echo -e "\n${GREEN}🎉 Audit completed successfully!${NC}"