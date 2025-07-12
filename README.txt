Project Setup

1. Plan the technologies to be used, make roadmap and folder structure

2. Install dependencies

go get -u github.com/gin-gonic/gin : Gin Web framework

go get -u gorm.io/gorm : GORM as ORM for DB

go get -u gorm.io/driver/postgres : Postgres Driver for GORM

go get -u github.com/go-redis/redis/v8 : Redis/v8 for cache

then git init creates .git directory to track changes

git remote add origin https://github.com/KrishKJ/targeting-engine.git : Add Your Remote GitHub Repo

docker compose up -d :  Start Containers from Compose File after installing docker. 

docker compose down -v : To nuke down the containers completely






REDIS CACHE INTEGRATION

1. Define a struct that represents cached campaign with rules

2. Load all ACTIVE campaigns and rules from PostgreSQL at app boot

3. Marshal that data to JSON and store it in Redis

4. Update ProcessDelivery() to use Redis (not DB)

5. Add a utility to refresh cache anytime (bonus: for future updates)