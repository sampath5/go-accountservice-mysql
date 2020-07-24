Go-accountservice using mysql

Go accountservice is written using Go-kit 

create a 'customer' table in mysql with columns customerid string,email string,phone int8

It runs on port 8000 

to create an account
POST:  /account

	{
		"customerid":"",
		"email":"some@email.com",
		"phone":xxxxxxxxxx
	}

To get an account by id

GET:	/account/{customerid}

To get all the customers details

GET: /account/getAll

To update an account

PATCH:	/account/update

	{
		"customerid":"",
		"email":"some@email.com",
		"phone":xxxxxxxxxx
	}

To delete an account

DELETE: /account/{customerid}


To run the service use command: go build 

This generates a AccountService.exe file and Run it
