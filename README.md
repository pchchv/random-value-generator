# HTTP service for generating random values.

## HTTP Methods
```
/generate — Generation of a random value and its identifier
options: 
    type — Returned random value type (
        num — number, 
        str — alphabetic string,
        alp — alphanumeric string)
    length — Length of the return value
    
example: http://localhost:8080/generate?type=alp&length=20
```
```
/retrieve — Getting the value from the id that was returned in the generate method
```
### Params for ```.env``` file
```
MONGO=mongodb://localhost:27017
DATABASE=random_values
COLLECTION=values
```