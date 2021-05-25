# PersonalNewsletter

PersonalNewsletter - Daily customized newsletter.
The idea is simple: scrape content from different websites, and get them every morning delivered per an email. Every user can customize what content they get at what time.


Every email consists of multiple "Modules". Each module has a name and can have a bunch of settings. These settings are just Key:Value pairs. This allows for a lot of customization.
These "Modules" get first converted into a "Feed". This feed doesn't contain any scraped data yet. The data gets scraped together by the code in the "aggregator" folder. It then can be converted into HTML with an template ("template/email1.html") and send via email ("sender/").