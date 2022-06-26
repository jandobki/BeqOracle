# BeqOracle
## how to run
Just do:
```
cd compose
docker compose build
docker compose up
```
Expect a plaintext gRPC service at `localhost:5000`. Find the proto file in `server/api`.

To execute integration tests 
```
cd compose
docker compose --profile test build 
docker compose --profile test up 
```
## shortcuts
There is a lot of simplification, unfinshed things and known issues in this solution. 
Unfortunately I haven't managed to fix all despite exceeding the suggested time limit.

### no security
To do (minimal): 
1. scripts to generate and sign certs,
2. add TLS support to server, read certs,
3. verify transport credentials or basic http authentication with a configured secret api key,
4. handle that in the integration tests.

### race condition 
There is a race condition between the select to verify the last state and insert to add the new state. 
I wanted to avoid including the validation logic on database level so some transaction context support
needs to be added to `Store`. In `PgStore` it could be implemented as a wrapper of an actual DB transaction with table lock.

### missing index
There should be an index to facilitate filtering on event.key.

### integration tests
There are only some dirty shell integration test containing the four scenarios included in the task description. 
They are not idempotent - they don't clean the DB after execution and can't be run again without cleaning the DB (`docker compose down -v`).

### no benchmarks and load tests
There certainly should be some, especially considering the additional questions.

### other
1. Logging is sparse. 
2. Context is not passed properly between `Server` object and `Store`. 
3. Error handling for non-logical errors is unstructured.

## Additional questions
### 1. How would you support multiple users?
1. An authentication mechanism should be used as mentioned above (other than the api key). Integrating an external identity provider would be the easiest option. 
2. Should answers be private to users or not? Would the keys be unique globaly or only in user scope? In a simpliest solution user id could be added to the event's key.

### 2. How would you support answers with types other than string?
Again it depends. Are many types to be supported for different keys at the same time? Can the type of the answer change for the same key?
In most cases it would involve implementing some form of variant starting from protobuf and down to the table schema, adding validation logic.
I'd probably start with:
```
message Value {
oneof test_oneof {
    string ValueString = 1;
    int32 ValueInt = 2;
    // ... etc.
  }
}
```

### 3. What are the main bottlenecks of your solution?
I haven't done any benchmarks but I'd guess that the filtering over key without the index (easily solvable)
and updating answers requires two DB queries in transaction with a table lock (depends on actual usage patterns).

### 4. How would you scale the service to cope with thousands of requests?
Would all request be that frequent or just getting answers? How many different keys would be queried?
If most of them would be `GetAnswer`s then a simple key-value read model table or an in-memory map would help depending on the size of the data. 
As the data can be easily partitioned by key, in more demanding cases a sharding load balancer over a key range could be employed and combined with caching. 
If sharding wouldn't be enough (eg. we'd like to have multiple instances of shards for high availability) then I'd think about Redis or memcached.