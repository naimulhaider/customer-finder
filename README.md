Reads customer data from file and finds all the customers within a given distance.

The program parses customer data as JSON, malformed lines are logged to the console but does not stop the program. 

Output: The ID and Name of customers found within the given Search Radius.

```
Customer Definition
- ID
- Name
- Longitude
- Latitude
```

Usage:

To see usage, run: `go run ./main.go -h`

```
-CUSTOMERS_FILE_LOCATION string
     The directory of customers.txt file (default "data/customers.txt")
        
-SEARCH_RADIUS float
    The distance (km) from source within which to search for customers (default 100)
    
-SOURCE_LAT float
    The latitude of source (default 53.339428)
    
-SOURCE_LON float
    The longitude of source (default -6.257664)
```