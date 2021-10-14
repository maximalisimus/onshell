# onshell

Run all scriptes the shebang on SHELL environment

Запуск shell скриптов в текущем shebang.

<img src="https://github.com/maximalisimus/onshell/raw/main/image/scr.png"  height="400">

## Оглавление

1. [Информация](#Информация)
2. [Information](#Information)
3. [Сборка](#Сборка)
4. [Build](#Build)
5. [Release](#Release)
6. [About-RUS](#About-RUS)
7. [About-EN](#About-EN)

[:arrow_up:Информация](#Информация)

Данное приложение разработано для запуска **SHELL** скриптов непосредственном в текущем **SHELL** из которого запущена программа.
Дело в том, что в разных ОС может быть не установлен **BASH** и по умолчанию может использоваться другой **SHELL**.

В качестве примера можно привести дистрибутив **Archlinux**, **ISO** образ. В нём по умолчанию отсутствует *bash* как таковой и используется *zsh*.

[:arrow_up:Information](#Information)

This application is designed to run **SHELL** scripts directly in the current **SHELL** from which the program is launched.
The fact is that **BASH** may not be installed in different operating systems and a different **SHELL** may be used by default.

An example is the distribution **ArchLinux**, **ISO** image. It does not have *bash* as such by default and uses *zsh*.

[:arrow_up:Сборка](#Сборка)

Вы можете запустить данную программу следующей командой:

```
	$ cd src
	$ go run onshell.go test.sh
```

Файл **test.sh** является тестовым скриптом, который вам необходимо запустить в текущем *SHELL*.

Для сборки в исполняемое приложение воспользуйтесь воспользуйтесь следующей командой:

```
	$ cd src
	$ GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' onshell.go
	# или
	$ GOOS=linux GOARCH=386 go build -ldflags '-w -s' onshell.go
```

При это исполняемый файл будет именоваться следующим образом: **onshell**.
Вы также можете сразу непосредственно создать исполняемые файлы для рызных ОС, воспользовавшись **Makefile**.

```
	$ cd src
	$ make all
	# и далее запуск
	$ onshell-amd64 test.sh
```

[:arrow_up:Build](#Build)

You can run this program with the following command:

```
	$ cd src
	$ go run onshell.go test.sh
```

File **test.sh** is a test script that you need to run in the current *SHELL*.

To build into an executable application, use the following command:

```
	$ cd src
	$ GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' onshell.go
	# или
	$ GOOS=linux GOARCH=386 go build -ldflags '-w -s' onshell.go
```

In this case, the executable file will be named as follows: **onshell**.
You can also directly create executables for different operating systems using **makefile**.

```
	$ cd src
	$ make all
	# и далее запуск
	$ onshell-amd64 test.sh
```

[:arrow_up:Release](#Release)

### Russian

Данный проект содержит актуальные **Release-версии** по следующей ссылке: [Release-versions](https://github.com/maximalisimus/onshell/releases).

### English

This project contains the latest **Release-versions** at the following link: [Release-versions](https://github.com/maximalisimus/onshell/releases).

[:arrow_up:About-RUS](#About-RUS)

Автор данной разработки **Shadow**: [maximalisimus](https://github.com/maximalisimus).

Имя автора: **maximalisimus**: [E-Mail](mailto:maximalis171091@yandex.ru).

Дата создания: **13.10.2021**

[:arrow_up:About-EN](#About-EN)

The author of this development **Shadow**: [maximalisimus](https://github.com/maximalisimus).

Author's name: **maximalisimus**: [E-Mail](mailto:maximalis171091@yandex.ru).

Date of creation: **13.10.2021**
