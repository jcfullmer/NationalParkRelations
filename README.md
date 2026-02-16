
This project allows a user that has used the (api2db) project to sort and filter through all of the rows by designation (National Park, National Monument, National Historic Site, etc...) and by full name.

## goBackend
The backend requires a .env file that includes the postgres database url (DB_URL).

| Method | URL                     | Parameters                                                  | Note                                                                                                   |
| ------ | ----------------------- | ----------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| GET    | /park                   | ?limit=integer(Default 50)<br>sort=full_name or designation | Provides all parks Full names and Designation                                                          |
| GET    | /park/state/{StateCode} | ?limit=integer(Default 50)                                  | StateCode will be the two letter representation of your desired state. <br>ID=Idaho<br>PA=Pennsylvania |
| GET    | /park/{ID}              | ?limit=integer(Default 50)                                  | Lookup park by TBD                                                                                     |
