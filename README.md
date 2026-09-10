# Yellow Commerce - Product Micro Service

To be used in conjunction with the Yellow Commerce API Gateway, supported by the Identty service.  (See dependencies below).  Implements a Faceted Search to slice and dice the list of products based on what the user has selected (in the preference service).

Connects to the gateway
Remember you can add more than one of the same service behind the gateway. Just
give it the same type but a different name

# Depends on
- YC-Gateway
- YC-Identity
- YC-Preference
- YC-Pricing
- YC-Images

Yes, CJ said to just smash it all into one service, but I am a purist. The solution contains a pricing cache and an images cache.  There is an option to take the Identity cache out of this, by moving the user context to the Preference service.  

# Improvements
As discussed, move the user context to the identity service,
Move the Like feature to the Reviews service

# External resources
- github.com/k-samuel/faceted/search


# Remember to
Change the port of this API both in the config and another place.  Remember the
port must be unique on the server.
Define the database in the Config

# Caution
In services/sku.go if a price is not in the cache, and the API does not return a price
(A error condition), then a very large negative number is returned as the price.  It is
for the consumer of the service to handle this correctly.


##Postgres Configuration
```
-- 1. Create the Database
create database category;

-- 2. Grant access to the Database
grant all privileges on database category to yellow;

-- 3.Connect to the database
\c category

-- 4. Create the Schema
CREATE SCHEMA dbo;

-- 5. Grant access to the Schema
GRANT USAGE ON SCHEMA dbo TO yellow;
GRANT CREATE ON SCHEMA dbo TO yellow;

-- 6. (Optional) Set dbo as the default for the user 'yellow'
-- This ensures GORM finds 'dbo' automatically without prefixing tables
ALTER ROLE yellow SET search_path TO dbo, public;
```
