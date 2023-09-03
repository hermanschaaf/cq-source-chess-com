# Table: chess_com_archives

This table shows data for Chess Com Archives.

The composite primary key for this table is (**username**, **year**, **month**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|username (PK)|`utf8`|
|year (PK)|`int64`|
|month (PK)|`int64`|