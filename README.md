# как запускать сервис
чтобы запустить с хранилищем на PostgreSQL
```
make -j 2 build-docker apply-migrations STORAGE=postgres
```

чтобы запустить с хранилищем на inmemory
```
make -j 2 build-docker apply-migrations STORAGE=inmemory
```

делать запросы к серверу лучше где-то через 30 секунд (потому что нужно время для применения миграций), но если вдруг docker-образы будут билдится дольше 30 секунд, то можно запустить сначала
```
make build-docker
```
Дождаться, пока сервис стартанет, а потом в отдельном терминале применить миграции
```
make apply-migrations
```

Если вдруг нужно сбилдить только образ приложения, тогда
```
STORAGE=inmemory|postgres docker-compose build my_app
```