#!/bin/sh

echo 'SCENARIO 1'

echo 'Create (succeeds)'
grpcurl -d '{"key": "testkey1", "value": "testvalue" }' \
    -plaintext server:5000 model.BeqOracle/CreateAnswer > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Delete (succeeds)'
grpcurl -d '{"key": "testkey1"}' \
    -plaintext server:5000 model.BeqOracle/DeleteAnswer > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Create (succeeds)'
grpcurl -d '{"key": "testkey1", "value": "testvalue"}' \
    -plaintext server:5000 model.BeqOracle/CreateAnswer  > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Update (succeeds)'
grpcurl -d '{"key": "testkey1", "value": "testvalue"}' \
    -plaintext server:5000 model.BeqOracle/UpdateAnswer  > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Get (succeeds)'
grpcurl -d '{"key": "testkey1"}' \
    -plaintext server:5000 model.BeqOracle/GetAnswer > /dev/null 2>&1 &&  echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'SCENARIO 2'

echo 'Create (succeeds)'
grpcurl -d '{"key": "testkey2", "value": "irure Excepteur fugiat esse id"}' \
    -plaintext server:5000 model.BeqOracle/CreateAnswer  > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }


echo 'Update (succeeds)'
grpcurl -d '{"key": "testkey2", "value": "irure Excepteur fugiat esse id"}' \
    -plaintext server:5000 model.BeqOracle/UpdateAnswer  > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Delete (succeeds)'
grpcurl -d '{"key": "testkey2"}' \
    -plaintext server:5000 model.BeqOracle/DeleteAnswer > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Create (succeeds)'
grpcurl -d '{"key": "testkey2", "value": "irure Excepteur fugiat esse id"}' \
    -plaintext server:5000 model.BeqOracle/CreateAnswer  > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Update (succeeds)'
grpcurl -d '{"key": "testkey2", "value": "irure Excepteur fugiat esse id"}' \
    -plaintext server:5000 model.BeqOracle/UpdateAnswer > /dev/null 2>&1  && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'SCENARIO 3'

echo 'Create (succeeds)'
grpcurl -d '{"key": "testkey3", "value": "testvalue" }' \
    -plaintext server:5000 model.BeqOracle/CreateAnswer > /dev/null 2>&1  && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Delete (succeeds)'
grpcurl -d '{"key": "testkey3"}' \
    -plaintext server:5000 model.BeqOracle/DeleteAnswer > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Update (error)'
grpcurl -d '{"key": "testkey3", "value": "testvalue"}' \
    -plaintext server:5000 model.BeqOracle/UpdateAnswer > /dev/null 2>&1 && { echo 'FAIL'; exit 1; } || { echo 'OK'; }

echo 'SCENARIO 4'

echo 'Create (succeeds)'
grpcurl -d '{"key": "testkey4", "value": "testvalue" }' \
    -plaintext server:5000 model.BeqOracle/CreateAnswer  > /dev/null 2>&1 && echo 'OK' || { echo 'FAIL'; exit 1; }

echo 'Create (error)'
grpcurl -d '{"key": "testkey4", "value": "testvalue" }' \
    -plaintext server:5000 model.BeqOracle/CreateAnswer > /dev/null 2>&1 && { echo 'FAIL'; exit 1; } || { echo 'OK'; }
