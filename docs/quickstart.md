# Быстрый старт

Создайте файл `multipass-compose.yaml`:

```shell
touch multipass-compose.yaml
```

Приведённая ниже конфигурация описывает три виртуальные машины с
именами `web-server`, `database` и `backend` без дополнительных параметров
конфигурации:

```yaml title="multipass-compose.yaml"
services:
  web-server:
    image: focal
  database:
    image: focal
  backend:
    image: focal
```

Скопируйте её в файл `multipass-compose.yaml`.

Для запуска виртуальных машин введите команду:

```shell
multipass-compose up
```

Дождитесь её завершения и проверьте состояние только что запущенных виртуальных
машин:

```shell
multipass-compose status
```

Для остановки виртуальных машин выполните команду:

```shell
multipass-compose down
```

После выполнения этой команды виртуальные машины будут полностью удалены.

## Примеры конфигурации

<!-- @formatter:off -->
Больше примеров конфигурации можно найти в отдельном репозитории [multipass-compose-showcase](https://github.com/pkorobeinikov/multipass-compose-showcase).
<!-- @formatter:on -->
