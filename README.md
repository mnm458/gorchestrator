# gorchestrator

A container orchestrator written purely in go
![image](https://github.com/user-attachments/assets/f86afded-befd-4106-a68b-6367840a96f8)

## Components

### Scheduler

The scheduler encompasses the aspect of feasibility, scoring and picking. These are the generic phases of the scheduler:

1. Feasibility: This phase checks if its possible to schedule a task onto a worker. There could be a possibility of the task not being schedulable onto any worker. The other possibility is task being schedulable only to a subset of workers.
2. Scoring: This phase takes the workers who are candidates for a task (determined by the feasibility phase), and gives each one a score.
3. Picking: Simple pick the best scoring candidate for the task.

### Manager

Manager uses the scheduler. The API is the primary mechanism for interacting with gorchestrator. The API servers the following purposes:

- Users can submit jobs and request jobs to be stopped via the API.
- Users can query the API to get information about job and worker status.

The manager has a job storage which allows for keeping track of all jobs in the system in order to make good scheduling decisions, as well as to provide answers to user queries about job and worker statuses. Lastly, the manager keeps track of worker metrics, such as the number of jobs a worker is currently running, how much memory it has available, CPU load/usage etc.

### Worker

The worker also has an API, however, serves a different purpose. The primary user of the API is the manager, which allows the manager to send tasks to the worker, indicate when to stop tasks, retrieve metrics about worker state etc.
Each worker has a task runtime, in this case, Docker. Worker nodes also keep track of its own work, which is done in the Task storage layer. Finally, the worker provides metrics about its own state, which it makes available via its API.
