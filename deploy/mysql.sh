#Install MySQL on Docker, set password and join to ops_net network
docker run --name mysql -it -p 3306:3306 -e MYSQL_ROOT_PASSWORD=friendly_almond_database -d mysql:5.7 --network ops_net
#Access to MySQL container
docker exec -it mysql bash
#Login mysql
mysql -uroot -p
friendly_almond_database
#Setting authentication for developer can use GUI to access the database
use mysql;
update user set authentication_string = password('friendly_almond_database') where user = 'root';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'friendly_almond_database' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'root'@'127.0.0.1' IDENTIFIED BY 'friendly_almond_database' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'root'@'47.74.20.100' IDENTIFIED BY 'friendly_almond_database' WITH GRANT OPTION;