# Overview

* Idea

PersonalNewsletter - Daily customized newsletter.
The idea is simple: scrape content from different websites, and get them every morning delivered per an email. Every user can customize what content they get at what time.

* Frontend

The frontend (build with Vue) allows users to customize their newsletters. It works "serverless" with Firebase. 

* Structure

Every email consists of multiple `Modules`. Each module has a name and can have a bunch of settings. 
These Modules get first converted into a "Feed". This feed doesn't contain any scraped data yet. The data gets scraped together by the code in the `aggregator` folder. It then can be converted into HTML with a template ("template/email1.html") and send via email ("sender/").

* Modules

    -   RSS Feed

* Planned modules:

    - Recipe of the day
    - Weather
    - Stocks
    - Horoscope
    - Subreddits

* How to run

_Backend_:

Clone the repo
Put your firebase keys in the main directory, rename them to `keys.json` (like `keys.example.json`)
Fill out the data in the `example.env` file, rename it to `.env`.
Run the golang code using `go run *.go`, or build it first using `go build *.go`

_Frontend_:

Install all needed NPM packets (`npm install`). 
Put your own Firebase credentials in the main.js file (`frontend/src/main.js`). Make sure that you created a database, synced the modules with the misc/newModules document, and have a collection named `newLetters`.


* Module structure

Current modules:

    - RSS feed

Each module consists of:

    - internalName (used internally, duh)
    - showName (name showed to the user on the frontend)
    - description (what does the module do?)
    - settings (collection of settings)

Each setting consists of:

    - internalName
    - showName
    - description
    - type (str, int, dropdown) 
    - data (only used for dropdown (name, value))

So, how do I add a new module?

You have to

    - Add it to modules.go
    - Add it to the FetchData function

You can sync the modules in `modules.go` with firebase by executing the program with the "sync" argument.