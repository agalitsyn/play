Doc http://docs.datastax.com/en/cassandra/3.0/cassandra/cassandraAbout.html

Availability and Partitioning tolerance (AP) https://wiki.apache.org/cassandra/ArchitectureOverview

-	TTL на записи
-	Immutable записи, есть mutation
-	Быстрая запись, не очень быстрое чтение
-	Пишет локально в память, потом дампит на диск
-	На диск пишет последовательно
-	Replication factor настраивается
-	Есть разные стратегии репликации
-	Партиционирование автоматическое по хешу ключа (вычисляется на базе)
-	Compaction надо планировать
-	Место под диск надо планировать

Maintenance

-	http://www.datastax.com/dev/blog/common-mistakes-and-misconceptions
-	https://docs.datastax.com/en/cassandra/3.0/cassandra/tools/toolsRepair.html
-	https://docs.datastax.com/en/cassandra/3.0/cassandra/tools/toolsCompact.html
