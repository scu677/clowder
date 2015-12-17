BEGIN TRANSACTION;
DROP TABLE IF EXISTS Machines;
CREATE TABLE Machines(uuid TEXT PRIMARY KEY, name TEXT, model TEXT, cpu TEXT, memory TEXT, hdd TEXT);
INSERT INTO Machines VALUES("4c4c4544-0036-5810-8043-b5c04f533532","blackmarsh1","Dell PowerEdge R320","Xeon E5-2407 (Sandy Bridge)","16GB","1TB SATA");

DROP TABLE IF EXISTS NetworkCards;
CREATE TABLE NetworkCards(mac TEXT PRIMARY KEY, uuid TEXT);
INSERT INTO NetworkCards VALUES("44:a8:42:34:6e:00","4c4c4544-0036-5810-8043-b5c04f533532");

DROP TABLE IF EXISTS Binding;
CREATE TABLE Binding(mac TEXT PRIMARY KEY, ip TEXT);
INSERT INTO Binding VALUES("44:a8:42:34:6e:00", "192.168.1.21");
INSERT INTO Binding VALUES("c8:00:84:66:d1:c0", "192.168.1.100");

DROP TABLE IF EXISTS Pxe;
CREATE TABLE Pxe(uuid TEXT PRIMARY KEY, rootpath TEXT, bootfile TEXT);
INSERT INTO Pxe VALUES("4c4c4544-0036-5810-8043-b5c04f533532", "192.168.1.1:/var/cluster/blackmarsh1", "pxeboot");


DROP TABLE IF EXISTS Users;
CREATE TABLE Users(username TEXT PRIMARY KEY, sshkey TEXT);

DROP TABLE IF EXISTS Reservations;
CREATE TABLE Reservations(machine TEXT PRIMARY KEY, username TEXT, fromday NUMERIC, untill NUMERIC, active NUMERIC);

COMMIT;
