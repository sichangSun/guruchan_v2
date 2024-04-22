-- create user table
CREATE TABLE User (
userId INT(6) PRIMARY KEY AUTO_INCREMENT,
userName VARCHAR(30) NOT NULL,
userPass VARCHAR(30) NOT NULL
);

-- create tag table
CREATE TABLE Tag (
tagId INT(20) PRIMARY KEY,
userId INT(6) NOT NUll,
foodId INT(20) NOT NULL,
tagTittle VARCHAR(40) NOT NULL,
typecode INT(2) DEFAULT '00',
t_addTime TIMESTAMP NOT NULL,
t_updataTime TIMESTAMP,
isDel INT(1) DEFAULT '0'
);

-- create food table
CREATE TABLE Food (
foodId INT(20) PRIMARY KEY,
userId INT(6) NOT NULL,
restaurantName VARCHAR(40) NOT NULL,
foodName VARCHAR(40),
address VARCHAR(16),
fullAddress VARCHAR(60),
isLikeflag INT(1) DEFAULT '0',
typecode INT(2),--種類
f_addTime TIMESTAMP NOT NULL,
f_updataTime TIMESTAMP,
firstTime TIMESTAMP,--初店行く時間
testedFlag INT(1),--行く予定？行った？
foodImg VARCHAR(100),
idDel INT(1) DEFAULT '0'
);
