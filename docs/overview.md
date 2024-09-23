The Chess.com source plugin for CloudQuery loads data from the Chess.com API to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Configuration

Example configuration:

```yaml
kind: source
spec:
  name: "chess-com"
  path: "hermanschaaf/chess-com"
  version: "v1.0.1"
  destinations:
    - "DESTINATION_NAME"
  tables:
    - "chess_com_games"
    - "chess_com_archives"
  spec:
    # plugin spec section
    usernames: # list of usernames to sync
     - magnuscarlsen
     - hikaru
     - gothamchess
```


You can also use `backend_options` to configure [incremental syncs](https://docs.cloudquery.io/docs/advanced-topics/managing-incremental-tables):

```yaml
kind: source
spec:
  name: "chess-com"
  path: "hermanschaaf/chess-com"
  version: "v1.0.1"
  backend_options:
    table_name: "chess_com_backend_options"
    connection: "@@plugins.DESTINATION_NAME.connection"
  destinations:
    - "DESTINATION_NAME"
  tables:
    - "chess_com_games"
    - "chess_com_archives"
  spec:
    # plugin spec section
    usernames: # list of usernames to sync
     - magnuscarlsen
     - hikaru
     - gothamchess
```

### Example Queries

Playing around with this data can be pretty fun. The PGNs and FENs for all games are included and can easily be loaded into a Chess engine. 

Here is one example query that shows a selection of 15 Magnus Carlsen games:

```sql
select 
  white->>'username' as white, 
  black->>'username' as black,
  white->>'result' as result_white,
  time_class
from chess_com_games 
where username = 'magnuscarlsen'
limit 15;
```

```text
+------------------+------------------+--------------+------------+
| white            | black            | result_white | time_class |
|------------------+------------------+--------------+------------|
| MagnusCarlsen    | stepanosinovsky  | resigned     | rapid      |
| MagnusCarlsen    | solskytz         | win          | rapid      |
| MagnusCarlsen    | penguingm1       | win          | rapid      |
| MagnusCarlsen    | Darkcloudy       | win          | rapid      |
| MagnusCarlsen    | jbj              | win          | rapid      |
| MagnusCarlsen    | Kacparov         | win          | rapid      |
| MagnusCarlsen    | cool64chess      | win          | rapid      |
| MagnusCarlsen    | CP6033           | win          | rapid      |
| MagnusCarlsen    | Tildenbeatsu     | win          | rapid      |
| MagnusCarlsen    | RainnWilson      | win          | rapid      |
| MagnusCarlsen    | mtmnfy           | win          | rapid      |
| MagnusCarlsen    | TigranLPetrosyan | win          | blitz      |
| MagnusCarlsen    | TigranLPetrosyan | win          | blitz      |
| TigranLPetrosyan | MagnusCarlsen    | resigned     | blitz      |
| TigranLPetrosyan | MagnusCarlsen    | resigned     | blitz      |
+------------------+------------------+--------------+------------+
```
