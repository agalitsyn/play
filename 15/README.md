# linux containers notes

Cостоят из трех частей:

1.	Изоляция - Linux namespaces Сейчас можно изолировать сеть, юзеров, хостнейм, маунты и прочее http://man7.org/linux/man-pages/man7/namespaces.7.html

2.	Ограничение действий - Linux capabilities, apparmor/selinux, seccomp. каждый CAP отвечает за некоторый набор действий http://man7.org/linux/man-pages/man7/capabilities.7.html Apparmor и SELinux это системы безопасности от Canonical и RedHat. Для них можно писать более сложные правила. http://man7.org/linux/man-pages/man2/seccomp.2.html - позволяет фильтровать syscalls и аргументы к ним.

3.	Ограничение ресурсов - Linux cgroups, Память, CPU и прочее. http://man7.org/linux/man-pages/man7/cgroups.7.html

http://lk4d4.darth.io/categories/namespaces/
