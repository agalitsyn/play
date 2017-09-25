# cpp python thrift

Пример клиент-серверного приложения с использованием Thrift

## Настройка окружения

Ставим пакеты из [статьи](https://thrift.apache.org/docs/install/debian)
```sh
sudo apt-get install libboost-dev libboost-test-dev libboost-program-options-dev libboost-system-dev libboost-filesystem-dev libevent-dev automake libtool flex bison pkg-config g++ libssl-dev ant
```

Дополнительно понадобятся:
```sh
sudo apt-get thrift-compiler python-dev python-thrift
```

### Сборка Thrift

Можно скачать исходники с [github](https://github.com/apache/thrift/tree/0.9.3).
В README есть описание, как установить. Версия `0.9.3`.

Важно, что на выходе вы должны обнаружить в своей системе файлы thrift в директориях:
* `/usr/local/lib`
* `/usr/local/include/thrift`

Возможно вы натолкнетесь на проблему в сборке tutorial файлов для C++
```
CppServer.cpp: In member function ‘virtual int32_t CalculatorHandler::calculate(int32_t, const tutorial::Work&)’:
CppServer.cpp:61:43: error: no match for ‘operator<<’ (operand types are ‘std::basic_ostream<char>’ and ‘const tutorial::Work’)
     cout << "calculate(" << logid << ", " << work << ")" << endl;
```

Можно просто закоментировать строки
```
diff --git a/tutorial/cpp/CppClient.cpp b/tutorial/cpp/CppClient.cpp
index 2763fee..33ce353 100644
--- a/tutorial/cpp/CppClient.cpp
+++ b/tutorial/cpp/CppClient.cpp
@@ -71,7 +71,7 @@ int main() {
     // costly copy construction
     SharedStruct ss;
     client.getStruct(ss, 1);
-    cout << "Received log: " << ss << endl;
+//    cout << "Received log: " << ss << endl;

     transport->close();
   } catch (TException& tx) {
diff --git a/tutorial/cpp/CppServer.cpp b/tutorial/cpp/CppServer.cpp
index eafffa9..bc755eb 100644
--- a/tutorial/cpp/CppServer.cpp
+++ b/tutorial/cpp/CppServer.cpp
@@ -58,7 +58,7 @@ public:
   }

   int32_t calculate(const int32_t logid, const Work& work) {
-    cout << "calculate(" << logid << ", " << work << ")" << endl;
+//    cout << "calculate(" << logid << ", " << work << ")" << endl;
     int32_t val;

     switch (work.op) {
```

## Генерация кода

На этом этапе сгенерится код по `.thrift` файлу

### Python

```sh
thrift -r --gen py --out python helloworld.thrift
```

### C++

```sh
thrift -r --gen cpp --out cpp helloworld.thrift
```

#### Сборка

```sh
cd cpp/
make
```

#### Запуск

```sh
cd cpp/
export LD_LIBRARY_PATH=/usr/local/lib/:${LD_LIBRARY_PATH}
./server
```
