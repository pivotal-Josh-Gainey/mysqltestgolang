# mysqltestgolang

Simple app that tests mysql connectivity as an app deployed on cloud foundry. This app opens a connection to the specified mysql database and closes it immediately after, recording the time it took. 

Just supply the connection string minus the driver. For example:

My app is pushed to cloud foundry and has the URL: `mysqlconnectiontest-busy-hartebeest.cfapps-48.slot-59.pez.vmware.com`

To test a connection simply curl with POST to the `/test` endpoint and supply the details in the data as key named **dataSourceName**

The full curl format looks like the following:

`curl -d 'dataSourceName=db_user:password@tcp(localhost:3306)/my_db' -X POST 'URL/test'`

Note you may have to specify the protocol in the connection path `......@tcp(IP:PORT)....`

An example:

`curl -d 'dataSourceName=CuKhwGRMbIewzsliZlEa:M5I6oE1BTXBoTPFGkf4LeUfN1X7bcI@tcp(mysql.service.cf.internal:3306)/autoscale' -X POST 'mysqlconnectiontest-busy-hartebeest.cfapps-48.slot-59.pez.vmware.com/test'`


As you curl this, open another terminal to stream the logs. Successful connection will appear as follows:

```
   2020-10-06T14:07:03.40-0400 [APP/PROC/WEB/0] OUT Successfully opened connection to the database in 0.000008 seconds.
   2020-10-06T14:07:03.40-0400 [APP/PROC/WEB/0] OUT DB functionality appears in tact. Now closing connection...
```
