# Команды (CLI)

В разделе описаны команды утилиты `multipass-compose`.

## multipass-compose init

Инициализация конфигурационного файла `multipass-compose.yaml`:

```shell
multipass-compose init
```

## multipass-compose up

Запуск виртуальных машин, описанных в файле `multipass-compose.yaml`:

```shell
multipass-compose up
```

## multipass-compose status

Вывод статуса виртуальных машин, описанных в файле `multipass-compose.yaml`:

```shell
multipass-compose status
```

## multipass-compose ip <machine>

Вывод `IPv4`-адреса виртуальной машины, имя которой передано в качестве
аргумента `<machine>`:

```shell
multipass-compose ip <machine>
```

Эта команда может быть полезна, например, при
генерации `Dynamic Inventory`[^dynamic-inventory] для `ansible`.

## multipass-compose down

Выключение и удаление виртуальных машин, описанных в
файле `multipass-compose.yaml`:

```shell
multipass-compose down
```

[^dynamic-inventory]: [Working with dynamic inventory](https://docs.ansible.com/ansible/latest/user_guide/intro_dynamic_inventory.html)
