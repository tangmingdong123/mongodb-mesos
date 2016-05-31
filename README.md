# 1 mongodb-mesos
a mesos scheduler for mesos . 
The scheduler provice a REST API to manage( create/delete/query ) your standalone mongodbs and replicaSets.

For standalone case, The scheduler monitor the mongodb's status, restart it when the task is killed or failed.

For replicaSet case, The scheduler will autoconfig the cluster ,and monitor the mongodb's status, restart it when the task is killed or failed.

You can get the mongodb instances detail infomation by REST API

# 2 start
./scheduler -master $mesos-master-ip:port -zk zk-ip:port -name schedulername -port httpport

for example : /scheduler -mesos 172.17.2.91:5050 -zk 172.17.2.91:2181 -name mongodb-mesos -port 37017

# 3 persistence
All standalone mongodbs and replicaSets' detail info are be saved in the zookeeper. The scheduler will reload these info when it restart. The zk' path is /${your scheduler name},and it is /mongodb-mesos by default. 

# 4 REST API
# create a standalone mongodb
curl -X DELETE --header 'Accept: application/json' --header 'Content-Type: application/x-www-form-urlencoded' -d 'cpu=1&mem=128' 'http://172.17.2.254:37017/standalone/1'

# delete a standalone mongodb
curl -X DELETE --header 'Accept: application/json' 'http://172.17.2.254:37017/standalone/1'

# list all standalone mongodbs
curl -X GET --header 'Accept: application/json' 'http://172.17.2.254:37017/standalone/list'

# get  a standalone mongodb's info
curl -X GET --header 'Accept: application/json' 'http://172.17.2.254:37017/standalone/1'

# create a replicaSet mongodb cluster
curl -X DELETE --header 'Accept: application/json' --header 'Content-Type: application/x-www-form-urlencoded' -d 'cpu=1&mem=128&instances=3' 'http://172.17.2.254:37017/replica/r1'

# delete a replicaSet mongodb cluster
curl -X DELETE --header 'Accept: application/json' 'http://172.17.2.254:37017/replica/r1'

# list all replicaSet mongodb clusters
curl -X GET --header 'Accept: application/json' 'http://172.17.2.254:37017/replica/list'

# get a replicaSet mongodb cluster's info
curl -X GET --header 'Accept: application/json' 'http://172.17.2.254:37017/replica/r1'
