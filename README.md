# watch-folder
A simple reference program written in golang to monitor file changes in a folder.  It does not depend on the iNotify for Linux or other OS FileSystem event notifications. This allows it to recognise changes to a folder even if the change was done by another computer, for eg:- on a mounted share.  This could work on Windows or MAC, but only tested on Linux.  Included a compiled version for Linux 64-bit.

Code list the folder contents and compare with previous listing.  It schedules the folder listing check every 30 seconds.  If previous listing and new listing didn't match for what ever reason, it prints out the new listing.  It does not try to identify what changed.

No guarantees or warrantees for the functionality of the code.  Use as you please.  Customize as you please.  This will just give you a head start.

WARNING!: This process may not be efficient if the folder structure is very large and may increase resouce usage or even crash the server.  I am not responsible for any issues caused by using this code.

Dependencies:

github.com/whiteShtef/clockwork is used as a scheduler



Instructions
--- bash
git clone https://github.com/abyzthomas/watch-folder.git

cd watch-folder

go get github.com/whiteShtef/clockwork

go build .



