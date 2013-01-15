GolemWiki
=========

GolemWiki is a single use, single purpose wiki intended to be usable for hackdays 
(such as the [NHSHackDay Wiki](http://wiki.nhshackday.com/wiki/NHSHackDayWiki)) 
that have recurring events and want to group the content for those events.

The data is currently stored in [PostgreSQL](http://www.postgresql.org/) and the 
app written in [Go](http://golang.org/) and is still under development (read 
"Not yet finished/runnable").

Some of the features are:

* OAUTH based login in the hope that tying users to a public account
might make them stop posting spam.
* Single level grouping of pages
* Single binary for installation
