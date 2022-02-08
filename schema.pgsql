-- DROP TABLE Users_Addresses;
-- DROP TABLE Users;
-- DROP TABLE Addresses;

CREATE TABLE Users (
    u_id SERIAL UNIQUE PRIMARY KEY,
    first_name text NOT NULL,
    middle_name text NOT NULL,
    last_name text NOT NULL
);

CREATE TABLE Addresses (
    a_id SERIAL UNIQUE PRIMARY KEY,
    street text NOT NULL,
    area text NOT NULL,
    pincode int NOT NULL,
    city text NOT NULL
);

CREATE TABLE Users_Addresses (
    id SERIAL UNIQUE PRIMARY KEY,
    u_id int,
    a_id int,
    FOREIGN KEY (u_id) REFERENCES Users(u_id) ON DELETE CASCADE,
    FOREIGN KEY (a_id) REFERENCES Addresses(a_id) ON DELETE CASCADE
);
"SELECT * FROM Users WHERE $1 = $1"