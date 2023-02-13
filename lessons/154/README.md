Mention

- hosts, cpu/memory
- latest versions


If you google you find this, let's test it:

Ultimately, it comes down to how you use them. MySQL is generally known to be faster with read-only commands at the cost of concurrency, while PostgreSQL works better with read-write operations, massive datasets, and complicated queries.

PostgreSQL is a feature-rich Database that can handle complex queries, while MySQL is a far simpler Database that is relatively simpler to set up, and manage and is fast, reliable, and easy to understand.



ssh -i ~/.ssh/id_ed25519 aputra@192.168.50.222

ansible-playbook -i inventory.ini --private-key ~/.ssh/id_ed25519 infra.yaml


aputra  ALL=(root) ALL

sudo vim /etc/sudoers.d/aputra
aputra ALL=(ALL:root) NOPASSWD: ALL







sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get -y install postgresql

sudo vim /etc/postgresql/15/main/postgresql.conf
listen_addresses = '*'


sudo vim /etc/postgresql/15/main/pg_hba.conf
host  all  all 0.0.0.0/0 scram-sha-256

sudo systemctl restart postgresql



psql -h 192.168.50.222 -p 5432 -U postgres


CREATE DATABASE benchmarks;
CREATE ROLE myapp WITH LOGIN PASSWORD 'devops123';

psql -h 192.168.50.222 -p 5432 -U myappv2 -d benchmarks

\c benchmarks



INSERT INTO authors(first_name, last_name)
VALUES ('Olga', 'Savelieva');

SELECT * FROM authors;

Joints
https://learnsql.com/blog/sql-join-examples-with-explanations/

Go Driver - https://github.com/jackc/pgx
MysQL driver - https://go.dev/doc/database/
MySQL driver - https://github.com/go-sql-driver/mysql

DROP TABLE authors;
CREATE TABLE authors (
    author_id  serial PRIMARY KEY,
    first_name varchar(40) NOT NULL,
    last_name  varchar(40) NOT NULL
);
SELECT count(*) FROM authors;






### Install MySQL

MySQL examples - https://github.com/go-sql-driver/mysql/wiki/Examples

sudo apt install mysql-server
sudo systemctl status mysql

sudo vim /etc/mysql/mysql.conf.d/mysqld.cnf
bind-address            = 0.0.0.0
mysqlx-bind-address     = 0.0.0.0

sudo systemctl restart mysql

sudo mysql -u root

CREATE DATABASE benchmarks;
CREATE USER 'myapp'@'0.0.0.0' IDENTIFIED BY 'devops123';
USE benchmarks;
CREATE TABLE authors (
    author_id  serial PRIMARY KEY,
    first_name varchar(40) NOT NULL,
    last_name  varchar(40) NOT NULL
);




CREATE USER 'myappv2'@'localhost' IDENTIFIED BY 'devops123';

GRANT ALL PRIVILEGES ON *.* TO 'myappv2'@'localhost' WITH GRANT OPTION;



CREATE USER 'myappv2'@'%' IDENTIFIED BY 'devops123';

GRANT ALL PRIVILEGES ON *.* TO 'myappv2'@'%' WITH GRANT OPTION;

FLUSH PRIVILEGES;






100 * (1 - ((avg_over_time(node_memory_MemFree_bytes[10m]) + avg_over_time(node_memory_Cached_bytes[10m]) + avg_over_time(node_memory_Buffers_bytes[10m])) / avg_over_time(node_memory_MemTotal_bytes[10m])))

node_memory_MemTotal_bytes - node_memory_MemFree_bytes

100 - ((node_memory_MemAvailable_bytes{instance="192.168.50.222:9100"} * 100) / node_memory_MemTotal_bytes{instance="192.168.50.222:9100"})

100 - ((node_memory_MemAvailable_bytes{instance="192.168.50.87:9100"} * 100) / node_memory_MemTotal_bytes{instance="192.168.50.87:9100"})

SELECT first_name,last_name FROM authors WHERE author_id = 1;

PG pool config
https://github.com/jackc/pgx/blob/4fc4f9a60337af3bd7c6abdf6c71460712d112fc/pgxpool/doc.go
https://github.com/jackc/pgx/blob/master/examples/url_shortener/main.go