#if !defined(_GNU_SOURCE)
	#define _GNU_SOURCE
#endif


#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/file.h>
#include <execinfo.h>
#include <unistd.h>
#include <errno.h>
#include <wait.h>

// лимит для установки максимально кол-во открытых дискрипторов
#define FD_LIMIT			1024*10

// константы для кодов завершения процесса
#define CHILD_NEED_WORK			1
#define CHILD_NEED_TERMINATE	2

#define PID_FILE "/var/run/my_daemon.pid"

// функция записи лога
void WriteLog(char* Msg, ...)
{
	// тут должен быть код который пишет данные в лог
}

// функция загрузки конфига
int LoadConfig(char* FileName)
{
	// тут должен быть код для загрузки конфига
	return 1;
}

// функция которая загрузит конфиг заново
// и внесет нужные поправки в работу
int ReloadConfig()
{
	// код функции
	return 1;
}

// функция для остановки потоков и освобождения ресурсов
void DestroyWorkThread()
{
	// тут должен быть код который остановит все потоки и
	// корректно освободит ресурсы
}

// функция которая инициализирует рабочие потоки
int InitWorkThread()
{
	// код функции
	return 1;
}

// функция обработки сигналов
static void signal_error(int sig, siginfo_t *si, void *ptr)
{
	void*  ErrorAddr;
	void*  Trace[16];
	int    x;
	int    TraceSize;
	char** Messages;

	// запишем в лог что за сигнал пришел
	WriteLog("[DAEMON] Signal: %s, Addr: 0x%0.16X\n", strsignal(sig), si->si_addr);


	#if __WORDSIZE == 64 // если дело имеем с 64 битной ОС
		// получим адрес инструкции которая вызвала ошибку
		ErrorAddr = (void*)((ucontext_t*)ptr)->uc_mcontext.gregs[REG_RIP];
	#else
		// получим адрес инструкции которая вызвала ошибку
		ErrorAddr = (void*)((ucontext_t*)ptr)->uc_mcontext.gregs[REG_EIP];
	#endif

	// произведем backtrace чтобы получить весь стек вызовов
	TraceSize = backtrace(Trace, 16);
	Trace[1] = ErrorAddr;

	// получим расшифровку трасировки
	Messages = backtrace_symbols(Trace, TraceSize);
	if (Messages)
	{
		WriteLog("== Backtrace ==\n");

		// запишем в лог
		for (x = 1; x < TraceSize; x++)
		{
			WriteLog("%s\n", Messages[x]);
		}

		WriteLog("== End Backtrace ==\n");
		free(Messages);
	}

	WriteLog("[DAEMON] Stopped\n");

	// остановим все рабочие потоки и корректно закроем всё что надо
	DestroyWorkThread();

	// завершим процесс с кодом требующим перезапуска
	exit(CHILD_NEED_WORK);
}


// функция установки максимального кол-во дескрипторов которое может быть открыто 
int SetFdLimit(int MaxFd)
{
	struct rlimit lim;
	int           status;

	// зададим текущий лимит на кол-во открытых дискриптеров
	lim.rlim_cur = MaxFd;
	// зададим максимальный лимит на кол-во открытых дискриптеров
	lim.rlim_max = MaxFd;

	// установим указанное кол-во
	status = setrlimit(RLIMIT_NOFILE, &lim);

	return status;
}


int WorkProc()
{
	struct sigaction sigact;
	sigset_t         sigset;
	int              signo;
	int              status;

	// сигналы об ошибках в программе будут обрататывать более тщательно
	// указываем что хотим получать расширенную информацию об ошибках
	sigact.sa_flags = SA_SIGINFO;
	// задаем функцию обработчик сигналов
	sigact.sa_sigaction = signal_error;

	sigemptyset(&sigact.sa_mask);

	// установим наш обработчик на сигналы

	sigaction(SIGFPE, &sigact, 0); // ошибка FPU
	sigaction(SIGILL, &sigact, 0); // ошибочная инструкция
	sigaction(SIGSEGV, &sigact, 0); // ошибка доступа к памяти
	sigaction(SIGBUS, &sigact, 0); // ошибка шины, при обращении к физической памяти

	sigemptyset(&sigset);

	// блокируем сигналы которые будем ожидать
	// сигнал остановки процесса пользователем
	sigaddset(&sigset, SIGQUIT);

	// сигнал для остановки процесса пользователем с терминала
	sigaddset(&sigset, SIGINT);

	// сигнал запроса завершения процесса
	sigaddset(&sigset, SIGTERM);

	// пользовательский сигнал который мы будем использовать для обновления конфига
	sigaddset(&sigset, SIGUSR1);
	sigprocmask(SIG_BLOCK, &sigset, NULL);

	// Установим максимальное кол-во дискрипторов которое можно открыть
	SetFdLimit(FD_LIMIT);

	// запишем в лог, что наш демон стартовал
	WriteLog("[DAEMON] Started\n");

	// запускаем все рабочие потоки
	status = InitWorkThread();
	if (!status)
	{
		// цикл ожижания сообщений
		for (;;)
		{
			// ждем указанных сообщений
			sigwait(&sigset, &signo);

			// если то сообщение обновления конфига
			if (signo == SIGUSR1)
			{
				// обновим конфиг
				status = ReloadConfig();
				if (status == 0)
				{
					WriteLog("[DAEMON] Reload config failed\n");
				}
				else
				{
					WriteLog("[DAEMON] Reload config OK\n");
				}
			}
			else // если какой-либо другой сигнал, то выйдим из цикла
			{
				break;
			}
		}

		// остановим все рабочеи потоки и корректно закроем всё что надо
		DestroyWorkThread();
	}
	else
	{
		WriteLog("[DAEMON] Create work thread failed\n");
	}

	WriteLog("[DAEMON] Stopped\n");

	// вернем код не требующим перезапуска
	return CHILD_NEED_TERMINATE;
}


void SetPidFile(char* Filename)
{
	FILE* f;

	f = fopen(Filename, "w+");
	if (f)
	{
		fprintf(f, "%u", getpid());
		fclose(f);
	}
}



int MonitorProc()
{
	int       pid;
	int       status;
	int       need_start = 1;
	sigset_t  sigset;
	siginfo_t siginfo;

	// настраиваем сигналы которые будем обрабатывать
	sigemptyset(&sigset);

	// сигнал остановки процесса пользователем
	sigaddset(&sigset, SIGQUIT);

	// сигнал для остановки процесса пользователем с терминала
	sigaddset(&sigset, SIGINT);

	// сигнал запроса завершения процесса
	sigaddset(&sigset, SIGTERM);

	// сигнал посылаемый при изменении статуса дочернего процесс
	sigaddset(&sigset, SIGCHLD);

	// сигнал посылаемый при изменении статуса дочернего процесс
	sigaddset(&sigset, SIGCHLD);

	// пользовательский сигнал который мы будем использовать для обновления конфига
	sigaddset(&sigset, SIGUSR1);
	sigprocmask(SIG_BLOCK, &sigset, NULL);

	// данная функция создат файл с нашим PID'ом
	SetPidFile(PID_FILE);

	// бесконечный цикл работы
	for (;;)
	{
		// если необходимо создать потомка
		if (need_start)
		{
			// создаём потомка
			pid = fork();
		}

		need_start = 1;

		if (pid == -1) // если произошла ошибка
		{
			// запишем в лог сообщение об этом
			WriteLog("[MONITOR] Fork failed (%s)\n", strerror(errno));
		}
		else if (!pid) // если мы потомок
		{
			// данный код выполняется в потомке

			// запустим функцию отвечающую за работу демона
			status = WorkProc();

			// завершим процесс
			exit(status);
		}
		else // если мы родитель
		{
			// данный код выполняется в родителе

			// ожидаем поступление сигнала
			sigwaitinfo(&sigset, &siginfo);

			// если пришел сигнал от потомка
			if (siginfo.si_signo == SIGCHLD)
			{
				// получаем статус завершение
				wait(&status);

				// преобразуем статус в нормальный вид
				status = WEXITSTATUS(status);

				 // если потомок завершил работу с кодом говорящем о том, что нет нужны дальше работать
				if (status == CHILD_NEED_TERMINATE)
				{
					// запишем в лог сообщени об этом
					WriteLog("[MONITOR] Childer stopped\n");

					// прервем цикл
					break;
				}
				else if (status == CHILD_NEED_WORK) // если требуется перезапустить потомка
				{
					// запишем в лог данное событие
					WriteLog("[MONITOR] Childer restart\n");
				}
			}
			else if (siginfo.si_signo == SIGUSR1) // если пришел сигнал что необходимо перезагрузить конфиг
			{
				kill(pid, SIGUSR1); // перешлем его потомку
				need_start = 0; // установим флаг что нам не надо запускать потомка заново
			}
			else // если пришел какой-либо другой ожидаемый сигнал
			{
				// запишем в лог информацию о пришедшем сигнале
				WriteLog("[MONITOR] Signal %s\n", strsignal(siginfo.si_signo));

				// убьем потомка
				kill(pid, SIGTERM);
				status = 0;
				break;
			}
		}
	}

	// запишем в лог, что мы остановились
	WriteLog("[MONITOR] Stopped\n");

	// удалим файл с PID'ом
	unlink(PID_FILE);

	return status;
}


int main(int argc, char** argv)
{
	int status;
	int pid;

	// если параметров командной строки меньше двух, то покажем как использовать демана
	if (argc != 2)
	{
		printf("Usage: ./my_daemon filename.cfg\n");
		return -1;
	}

	// загружаем файл конфигурации
	status = LoadConfig(argv[1]);

	if (!status) // если произошла ошибка загрузки конфига
	{
		printf("Error: Load config failed\n");
		return -1;
	}

	// создаем потомка
	pid = fork();

	if (pid == -1) // если не удалось запустить потомка
	{
		// выведем на экран ошибку и её описание
		printf("Start Daemon Error: %s\n", strerror(errno));

		return -1;
	}
	else if (!pid) // если это потомок
	{
		// данный код уже выполняется в процессе потомка
		// разрешаем выставлять все биты прав на создаваемые файлы,
		// иначе у нас могут быть проблемы с правами доступа
		umask(0);

		// создаём новый сеанс, чтобы не зависеть от родителя
		setsid();

		// переходим в корень диска, если мы этого не сделаем, то могут быть проблемы.
		// к примеру с размантированием дисков
		chdir("/");

		// закрываем дискрипторы ввода/вывода/ошибок, так как нам они больше не понадобятся
		close(STDIN_FILENO);
		close(STDOUT_FILENO);
		close(STDERR_FILENO);

		// Данная функция будет осуществлять слежение за процессом
		status = MonitorProc();

		return status;
	}
	else // если это родитель
	{
		// завершим процес, т.к. основную свою задачу (запуск демона) мы выполнили
		return 0;
	}
}

