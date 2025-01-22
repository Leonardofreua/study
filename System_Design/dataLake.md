# Data Lake

- Throw data into text files (csv, json, logs, pdf, etc.) into a big distributed storage system like Amazon S3 or Azure Blob Storage
- Common approach for 'Big data' and unstructured data
- There are tools for try create a schema for that data, such as: AWS Glue and Azure Data Factory
  - These tools are ETL (extract, transform, and load)
- Cloud-based features let you query the data:
  - AWS Athena (serverless) | Azure Synapse Analytics
  - AWS Redshift | Azure Synapse Analytics
  - Through these tools you can use SQL to analyze structure and semistructure as a traditional database
- You still need to think about how to partition the raw data for the best performance
  - Think like folders on a file system
  - Create folders/partitions by labels. For example: categories of a product, date (day, week, month, year), user name
  - Think about the most appropriate way to organize files according to the application context