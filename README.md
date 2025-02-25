# ANT Circulating Supply   
API endpoint to get the circulating supply of the ANT token.   

## Getting started:

#### Clone the repo:
```bash
git clone https://github.com/lajosdeme/ant-circulating-supply.git
```

#### Build with docker:
```bash
cd ant-circulating-supply

# This will use public Arbitrum One RPC by default
mv .env.example .env

# This will build with Docker and start the container on port 3000 of the host
make run

# If need to run on a custom port pass in port number:
make run PORT=8080
```

#### Test:
```bash
curl -X 'GET' 'http://localhost:3000/'
# {"circulating_supply":107335280.64517497}
```
