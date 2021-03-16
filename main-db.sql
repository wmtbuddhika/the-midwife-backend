-- MySQL dump 10.13  Distrib 8.0.19, for macos10.15 (x86_64)
--
-- Host: midwife.cnlodl6gzvqn.us-east-1.rds.amazonaws.com    Database: MidWife
-- ------------------------------------------------------
-- Server version	5.7.22-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `BloodGroup`
--

DROP TABLE IF EXISTS `BloodGroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `BloodGroup` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Code` varchar(45) DEFAULT NULL,
  `Name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `BloodGroup`
--

LOCK TABLES `BloodGroup` WRITE;
/*!40000 ALTER TABLE `BloodGroup` DISABLE KEYS */;
/*!40000 ALTER TABLE `BloodGroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Children`
--

DROP TABLE IF EXISTS `Children`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Children` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date_of_registration` varchar(45) DEFAULT NULL,
  `delivery_method` varchar(45) DEFAULT NULL,
  `no_of_apgar` varchar(45) DEFAULT NULL,
  `weight` varchar(45) DEFAULT NULL,
  `head_round` varchar(45) DEFAULT NULL,
  `height` varchar(45) DEFAULT NULL,
  `vitamin_k` varchar(45) DEFAULT NULL,
  `feed_first` varchar(45) DEFAULT NULL,
  `feed_correct` varchar(45) DEFAULT NULL,
  `feed_position` varchar(45) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Children`
--

LOCK TABLES `Children` WRITE;
/*!40000 ALTER TABLE `Children` DISABLE KEYS */;
INSERT INTO `Children` VALUES (3,'','1','','','','','1','1','1','',103),(4,'','1','','','','','1','1','1','',105);
/*!40000 ALTER TABLE `Children` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ChildrenVaccine`
--

DROP TABLE IF EXISTS `ChildrenVaccine`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ChildrenVaccine` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `vaccine_id` int(11) DEFAULT NULL,
  `effective_date` date DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ChildrenVaccine`
--

LOCK TABLES `ChildrenVaccine` WRITE;
/*!40000 ALTER TABLE `ChildrenVaccine` DISABLE KEYS */;
/*!40000 ALTER TABLE `ChildrenVaccine` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Contact`
--

DROP TABLE IF EXISTS `Contact`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Contact` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `address` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `mobile_phone_no` varchar(45) DEFAULT NULL,
  `land_phone_no` varchar(45) DEFAULT NULL,
  `emergency_address` varchar(45) DEFAULT NULL,
  `emergency_mobile_phone_no` varchar(45) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `latitude` decimal(13,10) DEFAULT NULL,
  `longitude` decimal(13,10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Contact`
--

LOCK TABLES `Contact` WRITE;
/*!40000 ALTER TABLE `Contact` DISABLE KEYS */;
INSERT INTO `Contact` VALUES (4,'Colombo','wmtbmail@gmail.com','+94761206292','0112787899','Colombo','',102,6.8142790000,79.9436260000),(5,'','','','','','',104,6.8142790000,79.9436260000),(6,'Colombo','','0761206292','','','',106,NULL,NULL),(7,'Colombo','','0761206292','','','',107,NULL,NULL),(8,'','','0761206292','','','',108,NULL,NULL),(9,'Colombo','','0761206292','','','',109,NULL,NULL),(10,'Colombo','','0761206292','','','',110,NULL,NULL),(11,'','','0761206292','','','',111,NULL,NULL);
/*!40000 ALTER TABLE `Contact` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Face`
--

DROP TABLE IF EXISTS `Face`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Face` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(45) DEFAULT NULL,
  `file_name` varchar(45) DEFAULT NULL,
  `file_path` varchar(45) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  `nic_no` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Face`
--

LOCK TABLES `Face` WRITE;
/*!40000 ALTER TABLE `Face` DISABLE KEYS */;
/*!40000 ALTER TABLE `Face` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Family`
--

DROP TABLE IF EXISTS `Family`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Family` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mother_nic_no` varchar(45) DEFAULT NULL,
  `father_nic_no` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Family`
--

LOCK TABLES `Family` WRITE;
/*!40000 ALTER TABLE `Family` DISABLE KEYS */;
INSERT INTO `Family` VALUES (16,'901371849V','901371849B'),(17,'901371849V',''),(18,'901371859V',''),(19,'effsfdsf','901371849V'),(20,'effsfdsf','901371849V');
/*!40000 ALTER TABLE `Family` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `FamilyChildren`
--

DROP TABLE IF EXISTS `FamilyChildren`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `FamilyChildren` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `family_id` varchar(45) DEFAULT NULL,
  `user_id` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `FamilyChildren`
--

LOCK TABLES `FamilyChildren` WRITE;
/*!40000 ALTER TABLE `FamilyChildren` DISABLE KEYS */;
INSERT INTO `FamilyChildren` VALUES (5,'17','103'),(6,'18','105');
/*!40000 ALTER TABLE `FamilyChildren` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Login`
--

DROP TABLE IF EXISTS `Login`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Login` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(45) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Login`
--

LOCK TABLES `Login` WRITE;
/*!40000 ALTER TABLE `Login` DISABLE KEYS */;
INSERT INTO `Login` VALUES (2,'901371849V','901371849V',1,1),(4,'937340150V','937340150V',102,1);
/*!40000 ALTER TABLE `Login` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Mother`
--

DROP TABLE IF EXISTS `Mother`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Mother` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `no_of_pregnancy` varchar(45) DEFAULT NULL,
  `death_births` varchar(45) DEFAULT NULL,
  `last_period_date` varchar(45) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Mother`
--

LOCK TABLES `Mother` WRITE;
/*!40000 ALTER TABLE `Mother` DISABLE KEYS */;
INSERT INTO `Mother` VALUES (1,'','0','',100),(2,'2','0','',102),(3,'','0','',104),(4,'2','0','2020-05-01',106),(5,'2','0','2020-05-01',107),(6,'2','0','2020-05-01',109),(7,'2','0','2020-05-01',110);
/*!40000 ALTER TABLE `Mother` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `User` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `surname` varchar(45) DEFAULT NULL,
  `first_name` varchar(45) DEFAULT NULL,
  `middle_name` varchar(45) DEFAULT NULL,
  `last_name` varchar(45) DEFAULT NULL,
  `preferred_name` varchar(45) DEFAULT NULL,
  `nic_no` varchar(45) DEFAULT NULL,
  `date_of_birth` varchar(45) DEFAULT NULL,
  `country_of_birth` varchar(45) DEFAULT NULL,
  `city_of_birth` varchar(45) DEFAULT NULL,
  `marital_state` varchar(45) DEFAULT NULL,
  `religion` varchar(45) DEFAULT NULL,
  `blood_group` varchar(45) DEFAULT NULL,
  `alcoholic` varchar(45) DEFAULT NULL,
  `diseases` varchar(45) DEFAULT NULL,
  `education` varchar(45) DEFAULT NULL,
  `occupation` varchar(45) DEFAULT NULL,
  `employer` varchar(45) DEFAULT NULL,
  `remarks` varchar(45) DEFAULT NULL,
  `user_type_id` varchar(45) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=112 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (1,NULL,'admin',NULL,'admin',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,'1','1'),(102,'WM','Thushara','De','Buddhika','Thushara','901371849V','2020-05-19','Sri Lanka','Colombo','1','1','1','0','No','Diploma','Farmer','Self','NO','2','1'),(103,'WM','Nethuki','','Shenaya','',NULL,'2020-03-18','','',NULL,'1','1',NULL,NULL,NULL,NULL,NULL,'','4','1'),(104,'D','Kamala','','Peiris','','901371859V','','','','1','1','1','0','','','','','','2','1'),(105,'','Saman','','Peiris','',NULL,'','','',NULL,'1','1',NULL,NULL,NULL,NULL,NULL,'','4','1');
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `UserType`
--

DROP TABLE IF EXISTS `UserType`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `UserType` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(45) DEFAULT NULL,
  `name` varchar(45) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `UserType`
--

LOCK TABLES `UserType` WRITE;
/*!40000 ALTER TABLE `UserType` DISABLE KEYS */;
/*!40000 ALTER TABLE `UserType` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Vaccine`
--

DROP TABLE IF EXISTS `Vaccine`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Vaccine` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `code` varchar(45) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  `month` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Vaccine`
--

LOCK TABLES `Vaccine` WRITE;
/*!40000 ALTER TABLE `Vaccine` DISABLE KEYS */;
/*!40000 ALTER TABLE `Vaccine` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Weight`
--

DROP TABLE IF EXISTS `Weight`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Weight` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` varchar(45) DEFAULT NULL,
  `weight` decimal(5,3) DEFAULT NULL,
  `month` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Weight`
--

LOCK TABLES `Weight` WRITE;
/*!40000 ALTER TABLE `Weight` DISABLE KEYS */;
INSERT INTO `Weight` VALUES (14,NULL,3.700,0,103),(15,NULL,4.300,1,103),(16,NULL,5.100,3,103);
/*!40000 ALTER TABLE `Weight` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'MidWife'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-05-14 22:39:09
