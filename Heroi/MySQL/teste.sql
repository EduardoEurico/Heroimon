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

ALTER TABLE heroes
    ADD poderes JSON NOT NULL,                          -- Adicionar coluna poderes como JSON
    ADD nivel_forca INT NOT NULL CHECK(nivel_forca BETWEEN 0 AND 100),  -- Adicionar nível de força com restrição de 0 a 100
    ADD popularidade INT NOT NULL CHECK(popularidade BETWEEN 0 AND 100), -- Adicionar popularidade com restrição de 0 a 100
    ADD status ENUM('Ativo', 'Inativo', 'Banido') NOT NULL,  -- Adicionar status do herói
    ADD vitorias INT NOT NULL,                           -- Adicionar coluna para vitórias
    ADD derrotas INT NOT NULL;                           -- Adicionar coluna para derrotas
