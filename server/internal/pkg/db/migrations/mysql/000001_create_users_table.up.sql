CREATE TABLE
    IF NOT EXISTS Users(
        ID INT NOT NULL UNIQUE AUTO_INCREMENT,
        Username VARCHAR (127) NOT NULL UNIQUE,
        Password VARCHAR (127) NOT NULL,
        Email VARCHAR (127) NOT NULL UNIQUE,
        Title VARCHAR (127),
        Location VARCHAR (127),
        Gender ENUM ('MALE', 'FEMALE', 'OTHER'),
        Picture VARCHAR (127),
        Company VARCHAR (127),
        FaveLangauge VARCHAR (127),
        FaveFramework VARCHAR (127),
        FaveTool VARCHAR (127),
        PRIMARY KEY (ID)
    )