## Deploy

### Hot to run:
```
git clone git@github.com:hesidoryn/jt.git
cd jt
go build
./jt -config=config.json
```
### Configuration
#### Basic configuration:
Port is mandatory.
```json
{
    "port": "3333"
}
```

#### Security
JT provides a tiny layer of authentication. A client can authenticate itself by sending the AUTH command followed by the password. You can specify password in `config.json` file:
```json
{
    "port": "3333",
    "password": "Password!"
}
```

#### Persistence
You can create point-in-time snapshots of the dataset by using SAVE command.
You need to specify path to file with snapshot in `config.json` file:
```json
{
    "port": "3333",
    "db": "dump.db"
}
```