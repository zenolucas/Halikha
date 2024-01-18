DROP TABLE IF EXISTS Artists;
DROP TABLE IF EXISTS Artworks;
DROP TABLE IF EXISTS Customers;
DROP TABLE IF EXISTS OrderItems;
DROP TABLE IF EXISTS Orders;
DROP TABLE IF EXISTS Users;

CREATE TABLE Users (
    UserID INT AUTO_INCREMENT PRIMARY KEY,
    Username VARCHAR(50),
    PasswordHash VARCHAR(100),
    Email VARCHAR(100),
    UserType ENUM('Artist', 'Customer') NOT NULL
);

CREATE TABLE Artists (
    ArtistID INT PRIMARY KEY,
    UserID INT UNIQUE,
    ArtistName VARCHAR(50),
    Bio TEXT,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE Customers (
    CustomerID INT PRIMARY KEY,
    UserID INT UNIQUE,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE Artworks (
    ArtworkID INT PRIMARY KEY,
    ArtistID INT,
    Title VARCHAR(100),
    ArtDescription TEXT,
    ImageFilePath VARCHAR(255),
    UploadDate DATETIME,
    FOREIGN KEY (ArtistID) REFERENCES Artists(ArtistID)
);

CREATE TABLE Orders (
    OrderID INT PRIMARY KEY,
    UserID INT, -- This can reference either Artists or Customers
    OrderDate DATETIME,
    TotalAmount DECIMAL(10, 2),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE OrderItems (
    OrderItemID INT PRIMARY KEY,
    OrderID INT,
    ArtworkID INT,
    Quantity INT,
    PricePerItem DECIMAL(10, 2),
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ArtworkID) REFERENCES Artworks(ArtworkID)
);