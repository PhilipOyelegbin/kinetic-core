# Kinetic Core API

## 📝 Project Description

This project involves creating a backend system for a workout tracker application where users can sign up, log in, create workout plans, and track their progress. The system will feature JWT authentication, CRUD operations for workouts, and generate reports on past workouts.

## ✨ Requirements

You are required to develop an API for a workout tracker application that allows users to manage their workouts and track their progress. Your first task is to think about the database schema and the API endpoints that will be needed to support the application’s functionality. Here are some of the key features you should consider:

- **Exercise Data**
  You should write a data seeder to populate the database with a list of exercises. Each exercise should have a name, description, and category (e.g., cardio, strength, flexibility) or muscle group (e.g., chest, back, legs). Exercises will be used to create workout plans.

- **User Authentication and Authorization**
  Users will be able to sign up, log in, and log out of the application. You should use JWTs for authentication and authorization. Only authenticated users should be able to create, update, and delete workout plans. Needless to say, users should only be able to access their own workout plans.

  - Sign-Up: Allow users to create an account.
  - Login: Allow users to log in to their account.
  - JWT: Use JSON Web Tokens for authentication.

- **Workout Management**
  Users will be able to create their workout plans. Workout plans should consist of multiple exercises, each with a set number of repetitions, sets, and weights. Users should be able to update and delete their workout plans. Additionally, users should be able to schedule workouts for specific dates and times.

  - Create Workout: Allow users to create workouts composed of multiple exercises.
  - Update Workout: Allow users to update workouts and add comments.
  - Delete Workout: Allow users to delete workouts.
  - Schedule Workouts: Allow users to schedule workouts for specific dates and times.
  - List Workouts: List active or pending workouts sorted by date and time.
  - Generate Reports: Generate reports on past workouts and progress.

## Constraints

You are free to choose the programming language and database of your choice. Actual decisions for the database schema, API endpoints, and other implementation details are up to you. However, you should consider the following constraints:

- Database: Use a relational database to store user data, workout plans, and exercise data.
- API: Develop a RESTful API to interact with the database.
- Security: Implement JWT authentication to secure the API endpoints.
- Testing: Write unit tests to ensure the correctness of your code.
- Documentation: Learn about OpenAPI Specs. Document your API endpoints and provide examples of how to use them.

## 🛠️ Core Technologies Used

- Go (Golang): The primary programming language.

- Standard Library Packages: Some standard internal library log, net/http, strconv, time, encoding/json, etc.

- External Library Packages: Some external library Gin, GORM, swagger, swag cli, etc.

## 🚀 Installation

To get expense tracker up and running on your local machine, follow these steps:

1. Ensure Go is Installed:
   Make sure you have Go installed (version 1.18 or higher is recommended). You can download it from go.dev/dl/.
   Verify your installation:

```
go version
```

2. Clone the Repository (or create project manually):
   If you're starting from scratch as part of a learning exercise, you'd create the project structure manually as described in the task instructions. If this were a real repository:

```
git clone https://github.com/philipoyelegbin/expense-tracker
cd expense-tracker
```

3. Initialize Go Module (if not already done):

```
go mod init github.com/philipoyelegbin/expense-tracker.git      # Only if you created the project manually
```

4. Build the Executable:
   This command compiles your Go source code into a single executable binary.

```
go build -o expense-tracker
```

This will create an executable file named expense-tracker in your project's root directory.

## 💡 Usage

Once built, you can run the CLI commands from your terminal.

**General Usage**

```
./expense-tracker     # Prompt you interactively to select an action to perform
```

## 📂 Project Structure

```
expense-tracker/
  ├── main.go # Main entry point and CLI command handling
  ├── Makefile # App script runner file
  └── config/ # Directory for app configuration
    ├── dbConfig.go # Entails the database configuration
  └── docs/ # Directory for swagger generated docs
  └── utils/ # Directory for app utilities
    ├── utils.go # Entails some helper functions.
  └── controller/ # Directory for defined logic
    ├── user-controller.go # Defines the user logic for all user routes
    ├── auth-controller.go # Defines the registration and login logic
    ├── expense-controller.go # Defines the expense logic for all expense routes
  └── model/ # Directory for defined types
    ├── types.go # Defines the data model and instantiate database
  └── routes/ # Directory for routes
    └── user-routes.go # Contain the routes for all user actions
    └── auth-routes.go # Contain the routes for register and login action
    └── expense-routes.go # Contains the routes for all expense actions
```

## 💾 Data Persistence

All data are stored on a mysql database.

## 🤝 Contributing

Contributions are welcome! If you'd like to contribute, please:

1. Fork the repository.

2. Create a new branch (git checkout -b feature/your-feature-name).

3. Make your changes.

4. Commit your changes (git commit -m 'feat: Add new feature').

5. Push to the branch (git push origin feature/your-feature-name).

6. Open a Pull Request.
