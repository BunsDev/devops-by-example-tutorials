If you google you find this, let's test it:

Ultimately, it comes down to how you use them. MySQL is generally known to be faster with read-only commands at the cost of concurrency, while PostgreSQL works better with read-write operations, massive datasets, and complicated queries.

PostgreSQL is a feature-rich Database that can handle complex queries, while MySQL is a far simpler Database that is relatively simpler to set up, and manage and is fast, reliable, and easy to understand.



ssh -i ~/.ssh/id_ed25519 aputra@192.168.50.222

ansible-playbook -i inventory.ini --private-key ~/.ssh/id_ed25519 infra.yaml


aputra  ALL=(root) ALL

sudo vim /etc/sudoers.d/aputra
aputra ALL=(ALL:root) NOPASSWD: ALL