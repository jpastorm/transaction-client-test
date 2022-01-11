# Transaction client API

before testing the api you have to run the sql commands from the sqlmigration folder connecting to the database with following credentials

* user: postgres
* password: password

Create Client

    http://127.0.0.1:1101/api/v1/public/client
		{
               "name":"anonn"
		}


Create Account for Client

	http://127.0.0.1:1101/api/v1/public/account
		{
              "client_id":3,
              "currency_id":2,
              "money":30
		}

Get Client
http://127.0.0.1:1101/api/v1/public/client/:idclient

Get Account
http://127.0.0.1:1101/api/v1/public/account/:idaccount

Get Transactions for account
http://127.0.0.1:1101/api/v1/public/transaction/:idaccount

Create Transaction
http://127.0.0.1:1101/api/v1/public/transaction

	{
        "type": "DEPOSIT",
        "account_holder":1,
        "money":30
	}
