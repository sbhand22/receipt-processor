# receipt-processor
Fetch go challenge 

# Receipt Processor

## Overview

This project implements a simple server that processes receipts and calculates points based on purchases. It runs locally on port 8080.

## Getting Started

### Prerequisites

Ensure you have Go installed on your machine. You can download and install Go from [here](https://golang.org/dl/).

### Running the Server

1. Clone the repository to your local machine.
2. Navigate to the root path of the project.
3. Run the following command to start the server:

 ```bash
   go run .
 ``` 

The server will start running on port 8080.

Testing the Application
Once the server is up, you can test the functionality by sending HTTP requests to process receipts and retrieve points.

Adding a Receipt
To add a receipt, use the following curl command:

   ```bash
    curl -X POST http://localhost:8080/receipts/process \ -H "Content-Type: application/json" \ -d "{\"retailer\":\"Target\",\"purchaseDate\":\"2022-01-01\",\"purchaseTime\":\"13:01\",\"items\":[{\"shortDescription\":\"Mountain Dew 12PK\",\"price\":\"6.49\"},{\"shortDescription\":\"Emils Cheese Pizza\",\"price\":\"12.25\"},{\"shortDescription\":\"Knorr Creamy Chicken\",\"price\":\"1.26\"},{\"shortDescription\":\"Doritos Nacho Cheese\",\"price\":\"3.35\"},{\"shortDescription\":\"   Klarbrunn 12-PK 12 FL OZ  \",\"price\":\"12.00\"}],\"total\":\"35.35\"}"
   ```
This command sends a POST request to the server with a JSON payload containing the details of the receipt.

Retrieving Points
To retrieve points for a specific receipt, use the following curl command with the appropriate receipt ID:

   ```bash
    curl -X GET http://localhost:8080/receipts/c56e508a-d7f4-42ee-ad41-71cf35180d6b/points
   ```
This command sends a GET request to the server to fetch the points associated with the given receipt ID.
