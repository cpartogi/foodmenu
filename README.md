## How To Setup 

### Prerequisites : 
1. Golang version 1.14 or higher
2. MySQL for database

### Steps :
1. Create database foodmenu
2. Go to folder /scripts/migrations
3. Copy content in file create_table_menu.sql
4. Paste point 3 to mysql client and run query
5. For first time installation use command : make install
6. To run unit test use command : make test
7. To run in local use command : make local
8. To run using docker container use : make compose-up
9. To stop docker container use command : make compose-down
10. From browser open this address : http://localhost:7100/swagger/index.html