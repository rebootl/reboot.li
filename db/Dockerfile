FROM mongo

#
COPY config.sh /docker-entrypoint-initdb.d/

# import script
COPY mongoimport.sh /docker-entrypoint-initdb.d/

# data to import
COPY entries.json /docker-entrypoint-initdb.d/

# data to import
COPY users.json /docker-entrypoint-initdb.d/
