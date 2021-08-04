# watch-folder
A simple reference program written in golang to monitor file changes in a folder.  It does not depend on the iNotify for Linux or OS FileSystem event notifications. This allows it to recognise changes to a folder even if the change was done by another computer, for eg:- on a mounted share.

Code list the folder contents and compare with previous listing.  It schedules the folder listing check every 30 seconds.  If previous listing and new listing didn't match for what ever reason, it prints out the new listing.  It does not do any try to identify what changed.

Dependencies:

github.com/whiteShtef/clockwork is used as a scheduler



Instructions

git clone https://github.com/abyzthomas/watch-folder.git

cd watch-folder

go get github.com/whiteShtef/clockwork

go build .



