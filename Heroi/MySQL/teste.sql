CREATE DATABASE heroes_db;

USE heroes_db;

CREATE TABLE heroes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome_real VARCHAR(100) NOT NULL,
    nome_heroi VARCHAR(100) NOT NULL,
    sexo ENUM('M', 'F', 'O') NOT NULL,
    altura DECIMAL(5, 2) NOT NULL,
    peso DECIMAL(5, 2) NOT NULL,
    data_nascimento DATE NOT NULL,
    local_nascimento VARCHAR(150) NOT NULL
);