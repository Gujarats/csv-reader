# CSV Reader
This will read all the csv files from directory. Finding the same column on each files csv.
And shows the same values on given column which exist on all the csv files.

# Install
You need to set your PATH for go environtment

```shell
$ git clone https://github.com/Gujarats/csv-reader
$ go install
$ . ~/.bashrc  // refresh your bash
$ csv-reader 
```

Or 

```shell
$ git clone https://github.com/Gujarats/csv-reader
$ go build 
$ ./csv-reader 
```

## Usage
Reading files csv on directory `/tmp/files/` for a specific column name

```shell
$ csv-reader --file /tmp/files/ column YourColumnName
```

Or reading files csv on directory `/tmp/files/` for given columns

```shell
$ csv-reader --file /tmp/files/ column YourColumnName1,YourColumnName2,YourColumnName3
```

Output sample : 

```shell
$ csv-reader --file /tmp/files/ column Column 
$ Finding values on column = Column
final result = [Kota Administrasi Jakarta Utara Bekasi Lhokseumawe Subulussalam Bengkulu Kota Administrasi Jakarta Pusat Serang Kota Administrasi Jakarta Barat Kota Administrasi Jakarta Timur Jambi Bogor Cirebon Meulaboh Cilegon Depok Bandung Sungai Penuh Cimahi Denpasar Kota Administrasi Jakarta Selatan Tangerang Gorontalo Langsa Sabang Pangkalpinang Tangerang Selatan]
final result size = 26
time consume = 0.004201499
```


