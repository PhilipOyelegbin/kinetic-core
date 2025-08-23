# Kinetic Core API

## 📝 Project Description

This project involves creating a backend system for a workout tracker application where users can sign up, log in, create workout plans, and track their progress. The system will feature JWT authentication, CRUD operations for workouts, generate reports on past workouts and rate limiter.

- Click **[here](https://kinetic-core.onrender.com)** to preview the API

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

  - Create Workout: Allow users to create workout plan composed of an exercise.
  - Update Workout: Allow users to update workout.
  - Delete Workout: Allow users to delete workout.
  - Schedule Workouts: Allow users to schedule workout for specific dates.
  - Filter Workouts: Filter workout based on status.
  - Generate Reports: Generate report on user workout.

## Constraints

You are free to choose the programming language and database of your choice. Actual decisions for the database schema, API endpoints, and other implementation details are up to you. However, you should consider the following constraints:

- Database: Use a relational database to store user data, workout plans, and exercise data.
- API: Develop a RESTful API to interact with the database.
- Security: Implement JWT authentication to secure the API endpoints.
- Documentation: Document your API endpoints and provide examples of how to use them.

## 🛠️ Core Technologies Used

- Go (Golang): The primary programming language.

- Standard Library Packages: Some standard internal library log, net/http, strconv, time, etc.

- External Library Packages: Some external library Gin, GORM, swaggo, swag cli, etc.

## 🚀 Installation

To get kinetic core up and running on your local machine, follow these steps:

1. Ensure Go is Installed:
   Make sure you have Go installed (version 1.18 or higher is recommended). You can download it from go.dev/dl/.
   Verify your installation:

   ```
   go version
   ```

2. Clone the Repository (or create project manually):
   If you're starting from scratch as part of a learning exercise, you'd create the project structure manually as described in the task instructions. If this were a real repository:

   ```
   git clone https://github.com/philipoyelegbin/kinetic-core
   cd kinetic-core
   ```

3. Initialize Go Module (if not already done):

   ```
   go mod init github.com/philipoyelegbin/kinetic-core.git      # Only if you created the project manually
   ```

4. Set up .env file in the root folder with the following varaibles

   ```
   PORT=8081
   DATABASE_URL=username:password@tcp(host:3306)/database_name?charset=utf8&parseTime=True&loc=Local
   JWT_SECRET=your_jwt_secret
   JWT_EXPIRATION_TIME=1
   SMTP_HOST=your_mail_host
   SMTP_PORT=your_mail_port
   SMTP_USER=your_mail_address
   SMTP_PASSWORD=your_mail_password
   APP_URL=http://localhost:8081/api/v1
   RATE=5
   CAPACITY=2
   ```

5. Build the Executable:
   This command compiles your Go source code into a single executable binary.

   ```
   go build -o kinetic-core
   ```

   This will create an executable file named kinetic-core in your project's root directory.

## 💡 Usage

Once built, you can run the CLI commands from your terminal.

**General Usage**

    ```
    ./kinetic-core     # Prompt you interactively to select an action to perform
    ```

## 📂 Project Structure

```
kinetic-core/
  └── api/    # Directory for route definition.
    └── routes.go   # Contain the routes for the application
  └── cmd/    # Directory for application-specific executables.
    └── main.go   # Main app entry point
  └── docs/   # Directory for swagger generated docs
  └── internal/   # Directory for private application and library code that is not intended for public.
    └── config/   # App configuration directory
    └── controllers/    # App function controller directory
    └── model/   # Database schema directory
  └── pkg/    # Directory for library code that is safe for external applications to import.
    └── middleware/   # App middleware directory
    └── seeders/    # Data seeder directory
    └── utils/    # App untility function directory
  └── go.mod    # Contain app installed packages
  └── go.sum    # Contain summary of installed packages
  └── Makefile    # App script runner
  └── README.md   # Documentation for the app
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
