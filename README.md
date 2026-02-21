
This project allows a user that has used the (api2db) project to sort and filter through all of the rows by designation (National Park, National Monument, National Historic Site, etc...) and by full name.

## goBackend
The backend requires a .env file that includes the postgres database url (DB_URL) and network port to listen and serve on (PORT).

| Method | URL                     | Parameters                                                  | Note                                                                                                   |
| ------ | ----------------------- | ----------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| GET    | /park                   | ?limit=integer(Default 50)<br>sort=full_name or designation | Provides all parks Full names and Designation                                                          |
| GET    | /park/state?state_code={StateCode} | ?limit=integer(Default 50)                                  | StateCode will be the two letter representation of your desired state. <br>ID=Idaho<br>PA=Pennsylvania |
| GET (IN PROGRESS)   | /park/{ID}              | ?limit=integer(Default 50)                                  | Lookup park by TBD                                                                                     |

##### By visiting the root at "/" in your web browser you will be given an AI generated index.html that you can use to help visualize and show off the get requests listed above. This part is the ONLY part that is AI generated and was done to give me the ability to show off my project to friends and family that are not as familiar with backend development. 

## Docker Container

I have a current docker container created at jcfullmer/npssearch or you can build your own with the docker file. Please just make sure to include the DB_URL and PORT environment variables. 
DB_URL: The full PostgreSQL connection string. Example: postgres://user:pass@Nip:5432/npsDB?sslmode=disable

PORT,The port the Go server listens on (internal). Example: :8080
```
docker run -d \
  --name=NPS_Search \
  -p 8080:8080 \
  -e DB_URL="postgres://user:password@IP:5432/npsDB?sslmode=disable" \
  -e PORT=":8080" \
  jcfullmer/npssearch
```
OR create a docker-compose.yml file that contains:
```
services:
  nps-search:
    image: jcfullmer/npssearch
    container_name: NPS_Search
    ports:
      - "8080:8080"
    environment:
      - DB_URL=postgres://user:password@IP:5432/npsDB?sslmode=disable
      - PORT=:8080
    restart: unless-stopped
```
