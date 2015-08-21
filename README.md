Clowder: a tool for herding cats
--------------------------------
This Branch hs been listed a s a pratice branch, it includes many code files so this wil be a summarry for some of the most important ones.

over all "DBandUI" folder should have all of the important files in it, it was used to take what was working from the "Database-mock" folder and the "ui-mock" folder and to start combining them. 

design/DBandUI/DBConnect/DBConnect.go  this contains the code for connecting to the database
design/DBandUI/DBGet/DBGet.go  this contains the code retriving information form the database
design/DBandUI/wiki3/wiki3.go  this contains the code for commuaction between the files for data retreaval and the
webinterfae (design/ui-mocks//wiki3/wiki3.go may have a mor compleat verson of this)
design/DBandUI/DBInsert/DBInsert.go this is much like the DBGet file but is used to enter data into the database, this sill needs to be tied in with the webinterface.
            
--------------------------------
to run the web interface run simply but in the file name you want to run in the comand line. ~/goWorkspace/src/github.com/musec/clowder/design/ui-mocks/wiki3/wiki3 
then open that file in the broswer http://localhost:8080/view/test should open this one
