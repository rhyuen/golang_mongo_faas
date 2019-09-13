# ABOUT

This may or may not be a REST API repo for Golang using MongoDB and Zeit Functions.

It uses the @now/go@canary builder so you can not have one function per folder that the @now/go builder is beholden to.

# TODO

- `DONE` Middleware folder called `mw` and package named `mw` for db connection and disconnect.

# Summary of Changes and Results

SEPT 8 2019: 
> Made a package called 'types' in a folder called 'types' in this repo and successfully referenced it in `/functions/*.go`.
> Removed the below changes and used the above instead as it worked out.

SEPT 7 2019: 
> Made a package called 'types' with a struct.  Gave it its own repo and successfully referenced it in `/functions/*.go`.
> It worked as intended but bothered me as I didn't want shared types in two different repos..