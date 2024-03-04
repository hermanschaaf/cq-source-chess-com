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

```
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