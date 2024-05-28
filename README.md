**BiteSpeed-Backend-Task**
```
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── controllers/
│   └── contact_controller.go
├── models/
│   └── contact.go
├── services/
│   └── contact_service.go
├── database/
│   └── database.go
├── go.mod
├── go.sum
├── README.md
└── .gitignore
```
[EndPoint URL] https://bitespeed-assessment-cf7u.onrender.com/identify

**HTTP Method **
```
POST

```
**Sample JSON**
```
{
    "phoneNumber": "123456",
    "email": "lorraine@hillvalley.edu"
}
```
**local machine Setup** 
```
1. Intall(if it is not there on machine) Go 1.18 or above (I have used 1.22.3).
2. Navigate to cmd/server/
3. use public dsn to the local machine with hosted database(here Postgres DB on render.com).
[public dsn] : postgres://nirmaldb_user:TRifLayjENPWFabVpwEcLFpn80aIHQFE@dpg-cpabursf7o1s73aebhbg-a.singapore-postgres.render.com/nirmaldb
4. run the command ```go run main.go```
```
