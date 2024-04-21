
# Simple User Management System

This project implements a User Management System, providing CRUD functionalities to manage user information using either MySQL or MongoDB databases based on environment variables. The system includes a Vue.js frontend application, a Golang backend API, and utilizes Docker for containerization. Additionally, monitoring is can be enabled through Prometheus and Grafana.

## Services

1. MySQL: MySQL database container for storing user data.
2. MongoDB: MongoDB database container for storing user data.
3. Backend: Golang API service providing endpoints to interact with the database.
4. Prometheus: Monitoring and alerting tool for tracking system metrics.
5. Grafana: Visualization and monitoring platform for Prometheus metrics.
6. Frontend: Vue.js application providing a user interface for managing users.
## Environment Variables

- DB_PORT: Port for the database (MySQL or MongoDB).
- MYSQL_ROOT_PASSWORD: Root password for MySQL.
- DB_USER: Username for accessing the database.
- DB_PASSWORD: Password for accessing the database.
- DB_NAME: Name of the database.
- MONGO_PORT: Port for MongoDB.
- SERVER_PORT: Port for the backend API.
- PROMETHEUS_PORT: Port for Prometheus.
- GRAFANA_PORT: Port for Grafana.
- GRAFANA_USER: Admin username for Grafana.
- GRAFANA_PASSWORD: Admin password for Grafana.
- FRONTEND_PORT: Port for the Vue.js frontend.
## Getting Started
- Clone the repository.
- Set environment variables in the .env file or use them with their default values.
- Run "make docker-run" to start all services.
- Access the frontend application through the provided port.

## Views
- User List:
![image](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/ffd372e9-fdf9-4138-9756-506e31076fc5)

- Create User:
![image](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/51a924b8-4860-4a97-aa11-079e4082a3c9)

- Edit User:
![image](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/e6e95315-5fcd-4c9d-b23d-2b673803f82c)

- Show User:
![image](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/e37b1f4e-0ab9-4498-a9e1-3381d448711b)

- Alerts and warnings:

![image](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/d99ec461-e420-4c63-8a7c-c04baa97e4f2)

![Ekran görüntüsü 2024-04-21 193354](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/5a5c3550-b020-4994-8a5f-26b695d2fe97)

![Ekran görüntüsü 2024-04-21 193517](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/8b156a83-3a4b-47f2-a008-a66039508220)

![image](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/94cbfdc2-b8ec-4071-b7e9-d6df8a7a6b53)

![image](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/b0f5fae2-dcb1-4af5-bfd6-a00d3ef6782a)

## Swagger API Doc

You can reach out the Swagger API Documentation via `/swagger` URL.

![image](https://github.com/velicanercan/simple-user-mgmt/assets/22958833/52f69bf8-a54a-4455-a404-8b7772fc4c7c)

Feel free to reach out if you have any questions or need further assistance!