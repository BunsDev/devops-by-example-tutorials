Mention

- hosts, cpu/memory


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

CREATE TABLE authors (
    author_id  serial PRIMARY KEY,
    first_name varchar(40) NOT NULL,
    last_name  varchar(40) NOT NULL
);

GRANT ALL PRIVILEGES ON authors TO myapp;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO myapp;

INSERT INTO authors(first_name, last_name)
VALUES ('Olga', 'Savelieva');

SELECT * FROM authors;

Joints
https://learnsql.com/blog/sql-join-examples-with-explanations/

Go Driver - https://github.com/jackc/pgx
MysQL driver - https://go.dev/doc/database/