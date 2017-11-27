## Deploy

### Hot to run:
```
git clone git@github.com:hesidoryn/jt.git
cd jt
go build
./jt -port=4567 -password=asdf -dump=dump.db
```
### Configuration
#### Basic configuration:
Default port is 3333. You can specify port using `port` flag, i.e.:
```
./jt -port=4567
```

#### Security
JT provides a tiny layer of authentication. A client can authenticate itself by sending the AUTH command followed by the password. You can specify password using `password` flag, i.e.:
```
./jt -password=asdf
```

#### Persistence
You can create point-in-time snapshots of the dataset by using SAVE command.
You need to specify path to file with snapshot using `dump` flag, i.e.:
```
./jt -dump=dump.db
```
