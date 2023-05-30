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
curl --location 'http://localhost:8080/'
curl --location 'http://localhost:8080/health'
curl --location 'http://localhost:8080/ping'
curl --location 'http://localhost:8080/ready'

# Get flight schedule
curl --location 'http://localhost:8080/api/v1/schedule' \
--header 'Content-Type: application/json' \
--data '{
    "conversationID": "d780da1b-6f28-431c-8b53-70e82fabd09a",
    "direction": 1,
    "departureAirPort": "CGK",
    "arrivalAirPort": "UPG",
    "departureDateTime": "2023-06-05T00:00:00",
    "arrivalDateTime": "2023-06-10T00:00:00",
    "passengerTypeQuantityADT": 1,
    "passengerTypeQuantityCNN": 1,
    "passengerTypeQuantityINF": 1,
    "cabinClass": "BUSINESS"
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