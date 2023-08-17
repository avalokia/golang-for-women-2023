## GORM Mini Challenge
The gorm mini challenge is to create a simple CRUD for two tables, which are products and variants.

## Functions
The CRUD functions are:
1. createProduct
1. updateProduct
1. getProductById
1. updateProductById
1. createVariant
1. getProductWithVariant
1. deleteVariantById

**The functions are initially commented in the code. Uncomment and edit part of the code to execute the function(s).**

## Environment
In my local environment, I used **PostgreSQL** for the database. Thus, the package that I imported in `db.go` is `"gorm.io/driver/postgres"`. In cases where other database is used, please change accordingly.

For connecting the database, the program is retrieving the following variables from the .env file:
* HOST=&lt;host_name&gt;
* DB_USER=&lt;database_user&gt;
* DB_PASSWORD=&lt;database_password&gt;
* DB_NAME=&lt;database_name&gt;
* DB_PORT=&lt;database_port&gt;

**The .env file is ignored from being submitted into the repository, so please create a new .env file before executing the program**
