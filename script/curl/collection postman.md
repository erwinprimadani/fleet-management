
### Colletion curl for postman

### 1. health check
curl --location 'http://localhost:3000/healthcheck'

### 2. get history vehicle
curl --location 'http://localhost:3000/vehicles/AA123QWE/history?start=1747815949&end=1747844749'

### 3. get last location vehicle
curl --location 'http://localhost:3000/vehicles/AA123QWE/location'

### 4. curl sent location to mqtt
curl --location 'localhost:3000/vehicles/sent/location' \
--header 'Content-Type: application/json' \
--data '{
    "vehicle_id":"S1234XYZ",
    "latitude": -6.2430015,
    "longitude": 106.8246234,
    "timestamp": ""
  }'