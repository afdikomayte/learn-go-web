# learn-go-web
learn basic web in golang

# structur project
```bash
project_folder
 |
 |--module_folder
 |  |
 |  |--application
 |  |   |
 |  |   |--query
 |  |   |   |
 |  |   |   |--interface.go
 |  |   
 |  |--domain
 |  |   |
 |  |   |--entity.go
 |  |
 |  |--infrastructure
 |  |   |
 |  |   |--mysql
 |  |   |   |
 |  |   |   |--interface_Imp.go
 |  |   |   |
 |  |--web
 |  |   |
 |  |   |--entity_name
 |  |   |   |
 |  |   |   |--entity_name.html
 |  |   |   |
 |
```
📦learn-go-web
 ┣ 📂config
 ┃ ┣ 📜database.go
 ┃ ┗ 📜database_test.go
 ┣ 📂customer
 ┃ ┣ 📂application
 ┃ ┃ ┗ 📂query
 ┃ ┣ 📂domain
 ┃ ┗ 📂infrastructure
 ┃ ┃ ┗ 📂mysql
 ┣ 📂handler
 ┃ ┗ 📜handler_customer.go
 ┣ 📜README.md
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┗ 📜main.go
