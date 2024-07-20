# Task Manager CLI

## Description

This command-line tool is designed for managing tasks and subtasks. It allows users to create tasks, manage their completion status, and handle subtasks efficiently. Tasks are saved to and loaded from a JSON file to ensure data persistence.

## Basic Operations

### 1. Add a New Task
- **Command**: `add <description>`
- **Description**: Adds a new task with the provided description. Each task is assigned a unique ID.

### 2. List All Tasks
- **Command**: `list`
- **Description**: Lists all tasks along with their subtasks. Displays the task ID, completion status, and description.

### 3. Mark a Task as Completed
- **Command**: `complete <taskID>`
- **Description**: Marks a specified task as completed if all its subtasks are also completed.

### 4. Delete a Task
- **Command**: `delete <taskID>`
- **Description**: Deletes the task with the given ID. Note that this will remove the task and all its subtasks permanently.

## Subtask Operations

### 1. Add a Subtask
- **Command**: `addsubtask <parentID> <description>`
- **Description**: Adds a subtask to the task with the specified parent ID. The subtask is assigned a unique ID.

### 2. List Subtasks for a Task
- **Command**: `listsubtasks <taskID>`
- **Description**: Lists all subtasks for the specified task. Displays each subtask's ID, completion status, and description.

### 3. Mark a Subtask as Completed
- **Command**: `completesubtask <taskID> <subtaskID>`
- **Description**: Marks a specific subtask as completed. Updates the parent task's completion status if all subtasks are completed.

### 4. Delete a Subtask
- **Command**: `deletesubtask <taskID> <subtaskID>`
- **Description**: Deletes a specific subtask from the task with the given ID. This action is permanent.

## Advanced Task Management

### 1. List Completed Tasks with Subtasks
- **Command**: `listcompleted`
- **Description**: Lists all tasks that are marked as completed, along with their subtasks.

### 2. List Incomplete Tasks with Subtasks
- **Command**: `listincomplete`
- **Description**: Lists all tasks that are not completed, along with their subtasks.

## Data Persistence

- **Tasks File**: Tasks are stored in a file named `tasks.json` located in the current working directory. Ensure this file is writable for saving tasks.
- **Loading and Saving**: The program automatically loads tasks from `tasks.json` at startup and saves any changes when exiting.

## Requirements

- Go programming language installed (version 1.x or later).
- Basic knowledge of command-line operations.

## Additional Notes

- The `getStatusSymbol` function is used to display a checkmark (✓) for completed tasks and an X (x) for incomplete tasks.
- Ensure that `tasks.json` is located in the working directory for proper data handling.
