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

to get an account by id

GET:	/account/{customerid}

To update an account
PATCH:	/account/update

	{
		"customerid":"",
		"email":"some@email.com",
		"phone":xxxxxxxxxx
	}

To delete an account

DELETE: /account/{customerid}
