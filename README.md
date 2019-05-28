# txn

Thor command
```
thor -c migrations/config.json reverse_db txn
```

```
jq -ncM 'while(true; .+1) | {method: "POST", url: "http://localhost:7000/transfer/internal", body: {to_account_number: . , from_account_number: (.+1) , amount:0.1} | @base64}'  | vegeta attack -lazy --format=json -name=100qps -duration=60s -rate=100 > results.100.new.qps.bin
```

```
vegeta plot results.100.new.qps.bin results.300.new.qps.bin results.600.new.qps.bin results.900.new.qps.bin results.1200.new.qps.bin > plot.combine.100-1200.html
```