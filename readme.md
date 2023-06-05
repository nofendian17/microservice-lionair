# Microservice JT (API)

This repo work in progress

## Installation
Clone repo 
```bash
git clone git@github.com:nofendian17/microservice-lionair.git
cd microservice-lionair
cp configs/config.example.json configs/config.json
go mod download
go mod tidy
```


Run
```bash
go run main.go
```

## Functional List
- Flight Schedule
- Pricing (TBC)
- Addons (TBC)
- Booking (TBC)
- Issued Ticket (TBC)
- Cancel Booking (TBC)

## Usage

```bash
# Health Routes
curl --location 'http://localhost:3000/'
curl --location 'http://localhost:3000/health'
curl --location 'http://localhost:3000/ping'
curl --location 'http://localhost:3000/ready'

# Get flight schedule
curl --location 'http://localhost:3000/api/v1/schedule' \
--header 'Content-Type: application/json' \
--data '{
    "conversationID": "344104c1-45e3-4ee8-9175-a5dab26b0bb8",
    "sessionID": "e43e2f67-b9fb-4e2e-80df-17e07e61390a",
    "direction": 1,
    "departureAirPort": "DPS",
    "arrivalAirPort": "AMQ",
    "departureDateTime": "2023-06-05T00:00:00",
    "arrivalDateTime": "2023-06-10T00:00:00",
    "passengerTypeQuantityADT": 1,
    "passengerTypeQuantityCNN": 1,
    "passengerTypeQuantityINF": 1
}'

# Pricing
# Addons
# Booking
# Issued
# Cancel

```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)