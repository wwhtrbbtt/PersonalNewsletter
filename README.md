# PersonalNewsletter

PersonalNewsletter - Daily customized newsletter.
The idea is simple: scrape content from different websites, and get them every morning delivered per an email. Every user can customize what content they get at what time.


Every email consists of multiple "Modules". Each module has a name and can have a bunch of settings. These settings are just Key:Value pairs. This allows for a lot of customization.
These "Modules" get first converted into a "Feed". This feed doesn't contain any scraped data yet. The data gets scraped together by the code in the "aggregator" folder. It then can be converted into HTML with an template ("template/email1.html") and send via email ("sender/").


## Warning:

Please, don't try to set it up yourself. I only have to host it on GitHub because managing the dependencies is easier that way. I won't update it that often, and I won't make a guide.
I want to make this usable though eventually.

## Modules:

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

You have to a) 
    Add it to modules.go
    Add it to the FetchData function

You can sync the modules in `modules.go` with firebase by executing the program with the "sync" argument.