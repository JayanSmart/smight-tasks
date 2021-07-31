# SmartKnight Task Manager

## Goal:

I want to write a go system to manage chores (repeating) and todo items (once -
off) for my household.

This system will run off a raspberryPi4 and have a simple GUI (though that can
be a separate front-end project).

A stretch goal/enhancement is to also keep track of a shared shopping list.



## Design:

This backend needs to manage 2 distinct items:

Tasks:
  - Tasks are once off items which need to be completed.
  - Tasks can optionally be assigned to someone.
  - Tasks can optionally have due dates
  - Tasks can be marked as Pending, In Progress, Completed, Not Done

Repeating Tasks (chores):
  - Repeating Tasks should keep track of who last completed them, and when.
  - When one of these tasks are completed, the task should be assigned to the
    next person on the list of "assignees"
  - Repeating Tasks can have an optional cadence (how regularly they need to 
    be completed)
  - Repeating Tasks Can be Paused (marked inactive and hidden), and then they
    can be resumed and the tool should remember who should be assigned the 
    task. 
  - Tasks can be completed by anyone, but only if the assigned person 
    completes the task, should the assignee change to the next person.

### Tech Stack

To achieve this I want to define a simple api for managing Tasks in OAPI3.

I will then generate a golang base project based on that specification.
I will use mongoDB as a backend DB for storing the Tasks.

The DB may not be running on the RaspberryPi.



