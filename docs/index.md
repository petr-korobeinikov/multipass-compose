# О проекте

`multipass-compose` разработан для оркестрации виртуальных машин на ОС Ubuntu с
помощью `multipass`[^multipass] и замены связки `Vagrant`[^vagrant]
и `VirtualBox`[^virtualbox] на компьютерах `Mac` с процессорами семейства `M1`.

`multipass-compose` также совместим с процессорами Intel, позволяя использовать
всей команде разработки один и тот же инструмент.

Сам `multipass` поддерживает различные виды виртуализации и чаще всего сильно
выигрывает по скорости запуска виртуальной машины перед `VirtualBox`.

[^multipass]: [multipass.run](https://multipass.run/)

[^vagrant]: [vagrantup.com](https://www.vagrantup.com/)

[^virtualbox]: [virtualbox.org](https://www.virtualbox.org/)
