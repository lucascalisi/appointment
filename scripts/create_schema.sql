USE appointment;

CREATE TABLE `users` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`email` varchar(150) NOT NULL,
	`password` varchar(32) NOT NULL,
	`status` varchar(500) NOT NULL,
	`whenCreated` datetime NOT NULL,
	`lastPasswordChange` datetime NULL,
	 PRIMARY KEY (`id`)
);

CREATE TABLE `roles`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(100) NOT NULL,
	`description` varchar(300) NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `users_roles`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`idUser` int(11) NOT NULL,
	`idRole` int(11) NOT NULL,
	PRIMARY KEY (`id`),
	FOREIGN KEY (idUser) REFERENCES users(id),
	FOREIGN KEY (idRole) REFERENCES roles(id)
);

CREATE TABLE `patients` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`dni` varchar(20)  NOT NULL UNIQUE,
	`name` varchar(500) NOT NULL,
	`idUser` int(11) NOT NULL,
	`sex` varchar(20) NOT NULL,
	`birthDay` datetime NULL,
	 PRIMARY KEY (`id`),
	 FOREIGN KEY (idUser) REFERENCES users(id)
);


CREATE TABLE `professionals` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`dni` varchar(20)  NOT NULL UNIQUE,
	`doctorNumber` varchar(20)  NOT NULL UNIQUE,
	`name` varchar(500) NOT NULL,
	`idUser` int(11) NOT NULL,
	`sex` varchar(50) NOT NULL,
	`birthDay` datetime NULL,
	 PRIMARY KEY (`id`),
	 FOREIGN KEY (idUser) REFERENCES users(id)
);

CREATE TABLE `specialties` (
	`id` int(11)  NOT NULL AUTO_INCREMENT,
	`name` varchar(100) NOT NULL,
	`description` varchar(300) NULL,
	 PRIMARY KEY (`id`)
);

CREATE TABLE `specialityDetails` (
	`id` int(11)  NOT NULL AUTO_INCREMENT,
	`idSpeciality` int(11) NOT NULL,
	`name` varchar(100) NOT NULL,
	`description` varchar(500) NULL,
	 PRIMARY KEY (`id`),
	 FOREIGN KEY (idSpeciality) REFERENCES specialties(id)
);

CREATE TABLE `specialityDetailsByProfessional` (
	`id` int(11)  NOT NULL AUTO_INCREMENT,
	`idSpecialityDetail` int(11) NOT NULL,
	`idProfessional` int(11) NOT NULL,
	 PRIMARY KEY (`id`),
	 FOREIGN KEY (idSpecialityDetail) REFERENCES specialityDetails(id),
	 FOREIGN KEY (idProfessional) REFERENCES professionals(id)
);

CREATE TABLE `professionalsSchedule` (
	`id` int(11)  NOT NULL AUTO_INCREMENT,
	`idProfessional` int(11)  NOT NULL,
	`year` int(11) NOT NULL,
	`month` int(11) NOT NULL,
	`dayOfWeek` int(11) NOT NULL,
	`startTime` time NOT NULL,
	`finishTime` time NOT NULL,
	 PRIMARY KEY (`id`),
	 FOREIGN KEY (idProfessional) REFERENCES professionals(id)
);

CREATE TABLE `appointments` (
	`id` int(11)  NOT NULL AUTO_INCREMENT,
	`idProfessional` int(11)  NOT NULL,
	`status` varchar(50) NOT NULL,
	`date` datetime NOT NULL,
	`idPatient` int(11)  NULL,
	 PRIMARY KEY (`id`),
	 FOREIGN KEY (idProfessional) REFERENCES professionals(id),
	 FOREIGN KEY (idPatient) REFERENCES patients(id)
);
