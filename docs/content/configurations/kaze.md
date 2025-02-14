# Kaze Configuration

Kaze is the master of the entire Gasper ecosystem which performs the following tasks

* Equal distribution of applications and databases among worker nodes
* User Authentication based on JWT (JSON Web Token)
* User API for performing operations on any application/database in any node (Identity Access Management is handled with JWT)
* Admin API for fetching and managing information of all nodes, applications, databases and users
* Removal of inactive nodes from the cloud ecosystem
* Re-scheduling of applications in case of node failure

Kaze API docs are available [here](/api)

The following section deals with the configuration of Kaze

```toml
##########################
#   Kaze Configuration   #
##########################

[services.kaze]
# Time Interval (in seconds) in which `Kaze` sends health-check probes
# to all worker nodes and removes inactive nodes from the central registry-server.
cleanup_interval = 600
deploy = true   # Deploy Kaze?
port = 3000
```

!!!tip
    You can reduce the value of **cleanup_interval** parameter in the above configuration if you need changes in your ecosystem to propagate faster but this will in turn increase the load on the Redis central registry server so *choose wisely*
