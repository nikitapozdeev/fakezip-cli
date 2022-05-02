# fakezip-cli
A small CLI utility for cloning the structure of zip archives.

It copies the structure of the archive but leaves the contents of the files empty.

### Why
I don’t think anyone will need this, due to a very situational use case, but who knows.

I needed this utility at work, when there are huge archives with a lot of files on a remote server 
and I need to write a script for processing them based on the archive structure without downloading a tons of gigabytes to the local machine.

### Usage
```console
$ fz source.zip target.zip
```
or just 
```console
$ fz source.zip
```
It will take the name of the source file and create the target file with the "-fake" suffix.
