cd ./frontend
yarn build
docker build -t shopme-frontend .
docker run -p 8080:80 -d --name frontend shopme-frontend
cd ..

cd ./frontend-admin
yarn build:prod
docker build -t shopme-frontend-admin .
docker run -p 8081:80 -d --name frontend-admin shopme-frontend-admin
cd ..

cd ./backend-go
docker build -t shopme-backend .

docker run --name shopme-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v /database/data:/var/lib/mysql -d mysql

docker run --link shopme-mysql:mysql -p 8000:8000 -d --name backend shopme-backend