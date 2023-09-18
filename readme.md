# Bookings and Reservations

This is the repository for my bookings and reservation project.

 - In order to start the project make sure that you have installed [VSCode](https://code.visualstudio.com/) and last version of [Golang](https://go.dev/)

 - To run the project clone the Repo to your device and download all used extentions then to run a web application open the terminal and write ( go run ./cmd/web)
    To halt the web application use ( Ctrl + C ).


Existing pages are 
- About (/about)
- Home (/home) 
- Rooms (/generals-quarters,/majors-suite)
- Contact (/contact)
- Reservations (/reservation-summary")



- Uses Go version 1.21
- Uses [Chi router](https://github.com/go-chi/chi)
- Uses [Alex Edwards SCS](https://github.com/alexedwards/scs/v2) session managament
- Uses [nosurf](https://github.com/justinas/nosurf )