# tada-auth
 
### How to Run The Project <a name="how-to-run-project"></a>

Open 3 terminals

Terminal 1 start mysql server with docker:
```bash
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=tada -d mysql:latest
```
or if you have already run it then just:
```bash
docker start mysql
```
Then create a Table named tada for our app, this might fail if you have already start the server component and that has completed the migration:
```sql
docker exec -it mysql bash
mysql -u root -p
CREATE DATABASE tada;
```
Terminal 2 run the server: 
```bash
cd server
go run ./server.go
```
Now navigate to https://localhost:4000 you can see graphiql playground and query the graphql server.

Terminal 3 start the frontend:
```bash
cd frontend
npm install
npm start
```
Now navigate to https://localhost:3000 you can see react frontend.
