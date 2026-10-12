### 1. HDFS (Hadoop Distributed File System)

- **Description:** A distributed, highly scalable file system designed to run on commodity hardware, storing large files across multiple machines using block replication.
    
- **Common Use Cases:** Massively scaled data lakes, archival storage, and batch processing data repositories.
    
- **Bot & Threat Detection Example:** Storing petabytes of raw historical firewall logs, web server access logs, and proxy logs for long-term auditing and offline security forensic analysis.
    
- **Direct Integration:** Serves as the primary storage layer for **YARN**, **Apache Hive**, and **Apache Spark**.
    

### 2. YARN (Yet Another Resource Negotiator)

- **Description:** Hadoop’s cluster resource management and job scheduling framework that decouples resource management from the programming model.
    
- **Common Use Cases:** Managing and scheduling distributed computing workloads (like MapReduce, Spark, and Hive jobs) across a shared cluster.
    
- **Bot & Threat Detection Example:** Allocating CPU and memory resources to run a heavy nightly batch job that scans weeks of HDFS logs to uncover latent botnet signatures.
    
- **Direct Integration:** Manages compute resources for **HDFS**, **Apache Hive**, and **Apache Spark** running within a Hadoop cluster.
    

### 3. Apache Hive

- **Description:** A distributed data warehouse software built on top of Hadoop that facilitates data summarization, ad-hoc querying, and analysis of large datasets using a SQL-like interface (HiveQL).
    
- **Common Use Cases:** Data warehousing, SQL-on-Hadoop analytics, and business intelligence reporting.
    
- **Bot & Threat Detection Example:** Querying historical security logs via SQL to identify sudden volumetric spikes in failed login attempts originating from specific geographic regions or ASN blocks.
    
- **Direct Integration:** Runs on top of **HDFS** (for storage) and **YARN** (for compute), and is frequently federated or queried by **Presto**.
    

### 4. Apache Kafka

- **Description:** A distributed event streaming platform capable of handling trillions of events a day, designed around a publish-subscribe messaging model with ordered commit logs.
    
- **Common Use Cases:** Real-time data pipelines, event-driven architectures, and stream ingestion.
    
- **Bot & Threat Detection Example:** Ingesting live clickstream telemetry, API gateway headers, and user session activity in real-time to feed into threat-monitoring pipelines.
    
- **Direct Integration:** Acts as the primary data ingestion pipeline feeding **Apache Flink**, **Apache Spark Streaming**, and **Apache Pinot**.
    

### 5. Apache ZooKeeper

- **Description:** A centralized service for maintaining configuration information, naming, providing distributed synchronization, and maintaining group services for large distributed systems.
    
- **Common Use Cases:** Leader election, distributed locking, cluster configuration, and membership coordination.
    
- **Bot & Threat Detection Example:** Coordinating state and handling leader election across a cluster of distributed real-time threat-monitoring nodes to ensure high availability.
    
- **Direct Integration:** Used internally for coordination and consensus by **Apache Kafka**, **HDFS**, and **YARN**.
    

### 6. Presto (Trino)

- **Description:** A fast, distributed SQL query engine designed for interactive analytical queries against disparate data sources of any size.
    
- **Common Use Cases:** Interactive data exploration, federated queries across multiple data stores, and low-latency reporting.
    
- **Bot & Threat Detection Example:** Interactively joining live threat metrics stored in Pinot with historical user behavior logs in Hive to investigate a complex credential-stuffing attack.
    
- **Direct Integration:** Directly queries data from **Apache Hive**, **Apache Pinot**, and **HDFS** without needing data migration.
    

### 7. Apache Spark

- **Description:** A unified, in-memory analytics engine for large-scale data processing, supporting batch processing, stream processing, SQL, and machine learning libraries.
    
- **Common Use Cases:** ETL pipelines, iterative machine learning model training, and large-scale graph analysis.
    
- **Bot & Threat Detection Example:** Training a machine learning classifier on historical bot traffic features (such as mouse movement entropy and request intervals) stored in HDFS.
    
- **Direct Integration:** Runs on **YARN**, reads data from **HDFS** and **Apache Kafka**, and outputs processed features to data stores.
    

### 8. Apache Flink

- **Description:** A framework and distributed processing engine for stateful computations over unbounded and bounded data streams with low latency and high throughput.
    
- **Common Use Cases:** Real-time fraud detection, complex event processing (CEP), and stream analytics.
    
- **Bot & Threat Detection Example:** Processing real-time clickstream events from Kafka to instantly trigger a CAPTCHA challenge when a velocity check for rapid failed logins is breached within a rolling 60-second window.
    
- **Direct Integration:** Consumes real-time event streams from **Apache Kafka** and can run on **YARN**.
    

### 9. Apache Pinot

- **Description:** A real-time distributed OLAP datastore designed to execute analytical queries with ultra-low latency on rapidly mutating streaming data.
    
- **Common Use Cases:** Real-time dashboards, user-facing analytics, and real-time metric monitoring.
    
- **Bot & Threat Detection Example:** Powering a Security Operations Center (SOC) dashboard that displays live blocked threats, active malicious IP trends, and bot attack vectors with sub-second query latency.
    
- **Direct Integration:** Ingests streaming data directly from **Apache Kafka** and can be queried concurrently via **Presto**.