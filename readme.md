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

# Planned features

The planned flow of using Wadjet is thus:

## 
You're on your terminal inside the folder of an application, you launch Wadjet and navigate to the application registry, if the folder you're in is already known to wadjet it will show the application highlighted at the top of the list, 

## Registering a new application
If you're not in a registered path, you can press the "add new application" key, this will move you to the registration form, where it will have the current folder you're on as it's path, you'll be able to give a name for your application and assign it a port

Wadjet wiill have a list of registered applications, through Wadjet you'll be able to check if they are running, launch them, register new applications, check if anything other then your ran application is being run on the port you've registered your application

- Register an application
- Register scripts for the application (run, build, test...)
- Run registered applications
- Kill running applications
