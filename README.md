# multipass-compose

> Compose multipass™ virtual machines like a ... charm!

## TL;DR

1. Define your `multipass-compose.yaml`:

   ```yaml
   services:
     web-server:
       image: focal
     database:
       image: focal
     backend:
       image: focal
   ```

2. Start machines up:

   ```shell
   multipass-compose up
   ```

3. Check the status:

   ```shell
   multipass-compose status
   ```

4. Do the job :hammer_and_wrench: :carpentry_saw: :hammer: :wrench:
5. Turn machines off:

   ```shell
   multipass-compose down
   ```



This project is licensed under the terms of the MIT license.
