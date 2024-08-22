
# Symptoms Tracker (SymptomsTrack) API

This repository contains the backend API for the SymptomsTrack app, designed to track and manage daily health symptoms securely. The API is built using Go and utilizes Gorm for ORM and PostgreSQL for database management.

## Getting Started

To start this application, follow these steps:

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-repo/symptoms-track.git

2. **Navigate to the Project Directory:**
   ```bash
   cd symptoms-track

3. **Run the Server:**
   ```bash
    go run main.go

## Features and Implementation

### Completed Features

- **User Onboarding:**
  - Secure account registration with email and password.
  - Ability to review the privacy statement and terms of service.

- **Authentication:**
  - Implemented using authentication tokens for secure access.

- **Symptom Management:**
  - Users can log in and add daily symptoms.
  - Users can view their past week's and month's symptoms.

### Pending Features

- **Symptom Creation:**
  - The symptom creation feature is under development and is approximately 80% complete.

- **Enhancements:**
  - Further refactoring of symptom management, including fixing issues with creating and updating symptoms, and adding the ability to get all symptoms.

## Feedback and Contributions

- **Reviewer:** Prabode
  - Comment: "The implementation demonstrates a basic knowledge of Go and Gorm, with the overall project being about 80% complete. Key areas such as symptom creation and management are nearly finalized."

## Project Structure

- **Controllers:** Manage API endpoints and process incoming requests.
- **Models:** Define the database schema and map data using Gorm.
- **Routes:** Define the API routes for accessing different features.
- **Database:** PostgreSQL is used for data persistence, managed through Gorm.

## Contribution Guidelines

1. **Fork the repository.**
2. **Create a new branch** with your feature or bug fix.
3. **Commit your changes** and push them to your branch.
4. **Create a pull request** for review.

## Testing

- Ensure to write unit tests for the controller methods to validate the API functionality. Integration tests are also recommended for verifying end-to-end functionality.

