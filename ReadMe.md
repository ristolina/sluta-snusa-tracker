# Individual project - Cloud Native Computing - Kristoffer Petterson
This is a webapp called Sluta Snusa Tracker. Is is used to track the users progress when quitting snus.

# Structure
The app is using a Django frontend with user management. The users data (that is tracked) is saved in a separated DB accessed through a separate service written in Golang.
## Frontend
Django webapp using MDL(Material Design Lite) for CSS. User handling and authentication is using built in Django Authentication.
## Backend
Go Gin Gonic HTTP server with a database connection to a SQL databse. Accepting GET and POST requests and return JSON.
## Database
MariaDB SQL database. 
Table format:
TABLE slutasnusa (
    id INT NOT NULL, 
    username CHAR(50), 
    quitdate TIMESTAMP,
    prillorperday INT NOT NULL,
    priceperdosa INT NOT NULL,
    prillorperdosa INT NOT NULL,
    created_at timestamp default current_timestamp,
    PRIMARY KEY (id)
    );


# How to use it?
Log in to your previously created account. Your tracking data is shown on the Home page. To update the data go to the Settings page.

# How to try it for yourself?
1. Clone the repo or download the files.
2. Create files with the different secret keys as referenced in the docker-compose.yaml
3. Run docker compose build
4. Run docker compose up
5. Connect to docker mariaDB instance and create the slutasnusa table and insert test data
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
        ('test', '2023-01-01T07:00:00', 12, 2999, 22 );

6. Webapp is available on http://localhost:8000