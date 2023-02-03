CREATE DATABASE slutasnusa;
USE slutasnusa;
CREATE TABLE slutasnusa (
    id INT NOT NULL AUTO_INCREMENT, 
    username CHAR(50), 
    quitdate TIMESTAMP,
    prillorperday INT NOT NULL,
    priceperdosa INT NOT NULL,
    prillorperdosa INT NOT NULL,
    created_at timestamp default current_timestamp,
    PRIMARY KEY (id)
    );

INSERT INTO slutasnusa (username, quitdate, prillorperday, priceperdosa, prillorperdosa) 
    VALUES 
        ('kristoffer', '2022-11-10T15:00:00', 11, 1999, 20), 
        ('tom', '2023-01-01T07:00:00', 12, 2999, 22 );

