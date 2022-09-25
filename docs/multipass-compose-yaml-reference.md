# multipass-compose.yaml

## Структура multipass-compose.yaml

Ниже приведён пример полной конфигурации кластера из трёх виртуальных машин.

```yaml
services:
  web-server: # (1)
    image: focal # (2)
    cpus: 2 # (3)
    mem: 2G # (4)
    disk: 5G # (5)
    cloud-init: cloud-init/config.yaml # (6)
  database:
    image: focal
    cpus: 4
    mem: 8G
    disk: 25G
    cloud-init: cloud-init/config.yaml
  backend:
    image: focal
    cpus: 4
    mem: 4G
    disk: 10G
    cloud-init: cloud-init/config.yaml
```

1. Имя виртуальной машины, выбирается произвольно.
2. Имя образа, обязательный параметр. Полный список доступен по
   команде `multipass find`.
3. Количество CPU.
4. Количество RAM.
5. Размер диска.
6. Путь или URL-адрес `user-data`-конфигурации `cloud-init`.

Количество машин и их имена могут быть любыми и выбираются исходя из ваших
задач.

## Примеры конфигурации

<!-- @formatter:off -->
Больше примеров конфигурации можно найти в отдельном репозитории [multipass-compose-showcase](https://github.com/pkorobeinikov/multipass-compose-showcase).
<!-- @formatter:on -->
