### How to Run The Project <a name="how-to-run-project"></a>

Open 3 terminals

Terminal 1 start mysql server with docker:
```bash
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=tada -d mysql:latest
```
Then create a Table names hackernews for our app:
```sql
docker exec -it mysql bash
mysql -u root -p
CREATE DATABASE tada;
```
Terminal 2 run the server: 
```bash
cd server
go run server/server.go
```
Now navigate to https://localhost:4000 you can see graphiql playground and query the graphql server.

Terminal 3 start the frontend:
```bash
cd frontend
npm install
npm start
```
Now navigate to https://localhost:3000 you can see react frontend.
