# Wadjet

Wadjet is a project I'm designing with two intents:

- Learn how make TUI software
- Produce a daemon that will manage all of my projects

# Wadjet Client

This is repository for the client TUI, it is built using the following libraries

- [BubbleTea](https://github.com/charmbracelet/bubbletea)
  - Application logic
- [LipGloss](https://github.com/charmbracelet/lipgloss)
  - Layouts and styling

# Planned flow of usage

Underneath the hood, the Wadjet TUI communicates with the Wadjet daemon, the daemon will watch all applications and periodically check their status, all data gathered by the daemon will be presented by the client.

Wadjet runs all of it's applications inside a TMux session, this makes it so applications will be running on the background and you can access their logs

## Launching Wadjet
You're on your terminal inside the folder of an application, you launch Wadjet, the TUI client will open up in a full terminal view, from here you'll see the wadjet presentation screen and the side bar containing all the menus

## Registry
You navigate to the registry on the sidebar and enter it, a list of your registered projects is shown, if you've opened Wadjet in a folder recognized by the Wadjet daemon, the project will be the first of the list and will be highlighted, this list shows the following information:

- name: project name
- path: project path
- status:
  - green: running
  - orange: partially running
  - red: not running

Pressing the "right arrow" key will show the project's registered applications below the selected project

## Registering a project 
If you're not in a registered path, you can press the "add new application" key, this will move you to the registration form, the current folder will be used as the path for the project but this can be changed, you'll be able to give a name for your project, after registration you'll go back to the registry list

## Registering a new application
With a project selected, pressing "e" will enter the project management screen, here you'll see a list of applications registered under this project with the following data:

- name: name of your application
- command: name of the command that will show on tools such as top (e.g.: go, node, your-application-name)
  - this is used to check if what's listening on the registered port is in fact your application
- port: port number it was assigned
- runner: tag that shows what runner the application has, such as: go, node, docker...
- status: the status of your application
  - green: running
  - orange: something program is using this application's port (command and name will be shown)
  - red: application set to be running but isn't or crashed
  - pink: application not set to be running
  - gray: application has no runner script
- last test results:
  - green: all passed
  - red: not all tests passed
  - pink: no test results
  - gray: application has no test script

## Registering Scripts
With the cursor selecting the desired application, hit the "e" key, this will bring you to the application configuration screen, here you will see a list of registered scripts for this application, pressing the "add new script" key will move you to a form with the following fields:

- file: you can type the name of the script file here, the form will check if the file exists before registering
- type:
  - runner
  - test

## Running applications
While selecting a project in the register, hitting enter will open the runner modal, this will list all the applications inside this project, you can then press "r" to run it or "t" to test it, if the application has that script configured it will send a command for the daemon to add it to the running jobs, which will run the application on the background