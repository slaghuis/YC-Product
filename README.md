# Yellow Commerce - Product Micro Service

To be used in conjunction with the Yellow Commerce API Gateway.  
Connects to the gateway
Remember you can add more than one of the same service behind the gateway. Just
give it the same type but a different name

# Remember to
Change the port of this API both in the config and another place.  Remember the
port must be unique on the server.
Define the database in the Config

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
