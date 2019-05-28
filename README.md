# txn
 
txn is transaction service , this service handle of kind transaction process such as transfer and get balance

we can use this following endpoint for to those process

## Rest API endpoint

- Transfer endpoint :

```
POST

http://localhost:7000/transfer/internal

Body 

{
	"from_account_number":4,
	"to_account_number":20,
	"amount":1
}
```

- Get Balance endpoint:

```
GET

http://localhost:7000/user/balance/3

```

## Schema DB

![Alt text](schema.png?raw=true "Schema DB")


## Spin up the project

- start the postgres , you can use `docker-compose up` or pull the postgres docker image manually and adjust the credential
- start the transaction service , you can build by docker
    - `cd to directory github.com/zainul/txn`
    - `docker build -t txn/txnapi:latest -f Dockerfile .`
    - `docker run --expose=7000 -p 7000:7000 -dit --restart unless-stopped --name txnapi --net="host" txn/txnapi`

## Testing

for unit test can run by

```
go test ./...
```

for stress test can run by supporting tools `https://github.com/tsenart/vegeta` 
- please follow the installation of `vegeta`
- install `jq`
- startup the txn service
- run this script

```
jq -ncM 'while(true; .+1) | {method: "POST", url: "http://localhost:7000/transfer/internal", body: {to_account_number: . , from_account_number: (.+1) , amount:0.1} | @base64}'  | vegeta attack -lazy --format=json -name=100qps -duration=60s -rate=100 > results.100.new.qps.bin
```

to view report use this: 
```
vegeta plot results.1000.qps.bin > plot.1000.html
```

## Test Result

The result in laptop core i5 RAM 8 Gb doing well in 750 QPS, but the best result in 250 and 500 QPS.

Actually for get more than 1000 TPS in my experience do it in real server , mostly when 750 QPS and 1000 QPS the postgres not enough to handle request
need more higher device

### 750 QPS
![Alt text](750QPS.png?raw=true "750 QPS")

```
Requests      [total, rate]            4995, 623.12
Duration      [total, attack, wait]    8.017160234s, 8.016048814s, 1.11142ms
Latencies     [mean, 50, 95, 99, max]  926.198302ms, 616.145096ms, 3.065032736s, 4.412892823s, 6.130091612s
Bytes In      [total, mean]            332196, 66.51
Bytes Out     [total, mean]            233548, 46.76
Success       [ratio]                  71.51%
Status Codes  [code:count]             0:1423  200:3572  
Error Set:
Post http://localhost:7000/transfer/internal: EOF
Post http://localhost:7000/transfer/internal: dial tcp: lookup localhost: device or resource busy
```
note : service can serve the request but little bit got unhandled the request


### 500 QPS
![Alt text](500QPS.png?raw=true "500 QPS")

```
Requests      [total, rate]            45000, 500.01
Duration      [total, attack, wait]    1m31.012152767s, 1m29.998060901s, 1.014091866s
Latencies     [mean, 50, 95, 99, max]  1.014726095s, 554.824864ms, 3.397519258s, 7.761742537s, 20.848561808s
Bytes In      [total, mean]            4184814, 93.00
Bytes Out     [total, mean]            3037664, 67.50
Success       [ratio]                  100.00%
Status Codes  [code:count]             0:2  200:44998  
Error Set:
Post http://localhost:7000/transfer/internal: EOF
```

note : service can serve and handle the request but response time is too high

### 250 QPS
![Alt text](250QPS.png?raw=true "250 QPS")

```
Requests      [total, rate]            22500, 250.01
Duration      [total, attack, wait]    1m30.001242629s, 1m29.995837361s, 5.405268ms
Latencies     [mean, 50, 95, 99, max]  8.769536ms, 4.889541ms, 32.052702ms, 60.042013ms, 165.002742ms
Bytes In      [total, mean]            2092314, 92.99
Bytes Out     [total, mean]            1507664, 67.01
Success       [ratio]                  99.99%
Status Codes  [code:count]             0:2  200:22498  
Error Set:
Post http://localhost:7000/transfer/internal: EOF

```

note : service can serve and handle the request and reponse time is under 50ms in 95 percentile

